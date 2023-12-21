package main

import (
	"CombineGenerator/combine_generator"
	"CombineGenerator/db"
	"CombineGenerator/utils"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func SaveDB(md *db.MysqlDriver, rd *db.RedisDriver, ctn *combine_generator.Container, wg *sync.WaitGroup, mtx *sync.Mutex, writer *csv.Writer, configID *int) {
	defer wg.Done()
	ret := md.Query("select * from landlord_combines where config_id_1st = ? "+
		"and config_id_2nd = ? and config_id_3rd = ?", ctn.ConfigID1st, ctn.ConfigID2nd, ctn.ConfigID3rd)
	if len(ret) == 0 {
		md.Insert("landlord_combines", []string{"config_id_1st", "config_id_2nd", "config_id_3rd", "remain_cards",
			"bomb", "king_bomb", "triple", "pair", "single"}, ctn.ValToStrList())
		if len(ret) > 0 {
			mtx.Lock()
			defer mtx.Unlock()
			var row []string
			*configID += 1
			row = append(row, strconv.Itoa(*configID), strconv.Itoa(int(ctn.ConfigID1st)), strconv.Itoa(int(ctn.ConfigID2nd)),
				strconv.Itoa(int(ctn.ConfigID3rd)), ctn.RemainCards.ToString(), strconv.Itoa(ctn.CardStat.BombCount),
				strconv.Itoa(ctn.CardStat.KingBombCount), strconv.Itoa(ctn.CardStat.TripleCount),
				strconv.Itoa(ctn.CardStat.PairsCount), strconv.Itoa(ctn.CardStat.SingleCount))
			err := writer.Write(row)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

}

var existed map[string]int

func OnSave(ctn *combine_generator.Container, writer *csv.Writer, configID int) {
	str := []string{strconv.Itoa(int(ctn.ConfigID1st)), strconv.Itoa(int(ctn.ConfigID2nd)),
		strconv.Itoa(int(ctn.ConfigID3rd))}
	exKey := strings.Join(str, "")
	if _, ok := existed[exKey]; !ok {
		var row []string
		row = append(row, strconv.Itoa(configID), strconv.Itoa(int(ctn.ConfigID1st)), strconv.Itoa(int(ctn.ConfigID2nd)),
			strconv.Itoa(int(ctn.ConfigID3rd)), ctn.RemainCards.ToString(), strconv.Itoa(ctn.CardStat.BombCount),
			strconv.Itoa(ctn.CardStat.KingBombCount), strconv.Itoa(ctn.CardStat.TripleCount),
			strconv.Itoa(ctn.CardStat.PairsCount), strconv.Itoa(ctn.CardStat.SingleCount),
			ctn.EventID1st, ctn.EventID2nd, ctn.EventID3rd)
		err := writer.Write(row)
		existed[exKey] = 1
		if err != nil {
			fmt.Println(err)
		}
	}
}

func GenerateCombines() {
	utils.ExcelToJson("excel/Combines.xlsx")

	rd := &db.RedisDriver{}
	rd.ConnectRedis()
	generator := combine_generator.NewGenerator()
	distFile, _ := os.Create("data.csv")

	writer := csv.NewWriter(distFile)
	defer writer.Flush()
	existed = make(map[string]int)
	idBegin := 10000
	for i := 0; i < 30000; i++ {
		succ, ctn := generator.DoGenerate()
		if ctn.GetTotalCount() != 54 {
			succ = false
		}
		if succ {
			OnSave(ctn, writer, idBegin+i)
		}
	}
	// generator.GenerateTest()
	writer.Flush()
}

func toAlphaString(i int) string {
	if i <= 0 {
		return ""
	}
	return string((i-1)%26 + 'A')
}

func main() {
	GenerateCombines()
}
