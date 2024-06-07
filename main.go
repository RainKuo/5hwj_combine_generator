package main

import (
	"CombineGenerator/classical_combine_generator"
	"CombineGenerator/combine_generator"
	"CombineGenerator/db"
	"CombineGenerator/proto/out/classical_combine"
	"CombineGenerator/utils"
	"encoding/csv"
	"fmt"
	"github.com/golang/protobuf/proto"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	SEQBOMB_BOMB_2_3  = 1 // 3个ID都只有2-3炸
	SEQBOMB_BOMB_1    = 2 // 至少有1个ID只有1炸
	SEQBOMB_BOMB_4    = 3 // 至少有1个ID有4炸
	SEQBOMB_BOMB_NONE = 4 // 没有炸弹
)

const (
	BUXIPAI_BOMB_1 = 1 // 3个ID只有1-2炸
	BUXIPAI_BOMB_0 = 2 // 有ID有0炸
	BUXIPAI_BOMB_3 = 3 // 有ID有3炸
	BUXIPAI_OTHER  = 4 // 其他类型
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

func OnSave(ctn *combine_generator.Container, writer *csv.Writer, configID int, bombType int, controlFlag int) {
	str := []string{strconv.Itoa(int(ctn.ConfigID1st)), strconv.Itoa(int(ctn.ConfigID2nd)),
		strconv.Itoa(int(ctn.ConfigID3rd))}
	exKey := strings.Join(str, "")
	if _, ok := existed[exKey]; !ok {
		var row []string
		row = append(row, strconv.Itoa(configID), strconv.Itoa(controlFlag), strconv.Itoa(bombType), strconv.Itoa(int(ctn.ConfigID1st)), strconv.Itoa(int(ctn.ConfigID2nd)),
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

func GenerateSeqbombCombines() {
	utils.ExcelToJson("excel/连炸基础牌型.xlsx")

	rd := &db.RedisDriver{}
	rd.ConnectRedis()
	generator := combine_generator.NewGenerator("./res/config/config_seqbomb_base_combines.json")
	distFile, _ := os.Create("seqbombCombines.csv")

	writer := csv.NewWriter(distFile)
	defer writer.Flush()
	existed = make(map[string]int)
	for controlFlag := 0; controlFlag < 3; controlFlag++ {
		idBegin := 10000 * (controlFlag + 1)
		for i := 0; i < 1000; i++ {
			succ, ctn := generator.DoGenerate()
			//if ctn.GetTotalCount() != 54 {
			//	succ = false
			//}
			if ctn.CardStat.GetTotal() != 51 || ctn.RemainCards.GetTotal() != 3 {
				succ = false
			}
			if succ {
				ids := []uint32{ctn.ConfigID1st, ctn.ConfigID2nd, ctn.ConfigID3rd}
				bomb1 := false
				bomb23 := true
				bomb4 := false
				for _, id := range ids {
					conf := generator.Configs.ConfigIDMap[id]
					if conf.Bomb == 1 {
						bomb1 = true
						bomb23 = false
					} else if conf.Bomb >= 2 && conf.Bomb <= 3 {
						// bomb23 = true
					} else if conf.Bomb == 4 {
						bomb4 = true
						bomb23 = false
					}
				}

				if bomb1 {
					OnSave(ctn, writer, idBegin+i, SEQBOMB_BOMB_1, controlFlag)
				}
				if bomb23 {
					OnSave(ctn, writer, idBegin+i, SEQBOMB_BOMB_2_3, controlFlag)
				}
				if bomb4 {
					OnSave(ctn, writer, idBegin+i, SEQBOMB_BOMB_4, controlFlag)
				}
				if !bomb1 && !bomb23 && !bomb4 {
					OnSave(ctn, writer, idBegin+i, SEQBOMB_BOMB_NONE, controlFlag)
				}
			}
		}
	}

	// generator.GenerateTest()
	writer.Flush()
}

func GenerateBuxipaiCombines() {
	utils.ExcelToJson("excel/不洗牌基础牌型.xlsx")

	rd := &db.RedisDriver{}
	rd.ConnectRedis()
	generator := combine_generator.NewGenerator("./res/config/config_buxipai_base_combines.json")
	distFile, _ := os.Create("buxipaiCombines.csv")

	writer := csv.NewWriter(distFile)
	defer writer.Flush()
	existed = make(map[string]int)
	for controlFlag := 0; controlFlag < 3; controlFlag++ {
		idBegin := 10000 * (controlFlag + 1)
		for i := 0; i < 1000; i++ {
			succ, ctn := generator.DoGenerate()
			//if ctn.GetTotalCount() != 54 {
			//	succ = false
			//}
			if ctn.CardStat.GetTotal() != 51 || ctn.RemainCards.GetTotal() != 3 {
				succ = false
			}
			if succ {
				ids := []uint32{ctn.ConfigID1st, ctn.ConfigID2nd, ctn.ConfigID3rd}
				bomb12 := true
				bomb0 := false
				bomb3 := false
				for _, id := range ids {
					conf := generator.Configs.ConfigIDMap[id]
					if conf.Bomb == 0 {
						bomb0 = true
						bomb12 = false
					} else if conf.Bomb >= 1 && conf.Bomb <= 2 {
						// bomb12 = true
					} else if conf.Bomb == 3 {
						bomb3 = true
						bomb12 = false
					} else {
						bomb12 = false
					}
				}

				if bomb12 {
					OnSave(ctn, writer, idBegin+i, BUXIPAI_BOMB_1, controlFlag)
				}
				if bomb0 {
					OnSave(ctn, writer, idBegin+i, BUXIPAI_BOMB_0, controlFlag)
				}
				if bomb3 {
					OnSave(ctn, writer, idBegin+i, BUXIPAI_BOMB_3, controlFlag)
				}
				if !bomb12 && !bomb0 && !bomb3 {
					OnSave(ctn, writer, idBegin+i, BUXIPAI_OTHER, controlFlag)
				}
			}
		}
	}

	// generator.GenerateTest()
	writer.Flush()
}

func OnCompressClassicalCombine(combine classical_combine_generator.BaseCombine) int64 {
	data := combine.Event
	data = (data << 5) | combine.BigCount
	data = (data << 5) | combine.Bomb
	data = (data << 5) | combine.Triple
	data = (data << 5) | combine.Pair
	data = (data << 5) | combine.Single
	return data
}

func OnSaveClassicalCombine(id int, controlFlag int, combines []classical_combine_generator.BaseCombine, writer *csv.Writer, cfg *classical_combine.Config) {
	var row []string
	remainBigCount := 6 - combines[0].BigCount - combines[1].BigCount
	compressed0 := OnCompressClassicalCombine(combines[0])
	compressed1 := OnCompressClassicalCombine(combines[1])
	compressed2 := OnCompressClassicalCombine(combines[2])
	rivalType := classical_combine_generator.GetRivalType(combines)
	row = append(row, strconv.Itoa(id), strconv.Itoa(controlFlag), strconv.Itoa(rivalType), strconv.FormatInt(compressed0, 10),
		strconv.FormatInt(compressed1, 10), strconv.FormatInt(compressed2, 10), strconv.FormatInt(remainBigCount, 10))
	writer.Write(row)
	singleConfig := classical_combine.ClassicalCombineConfig{}
	singleConfig.ID = uint32(id)
	singleConfig.ControlFlag = uint32(controlFlag)
	singleConfig.RivalType = uint32(rivalType)
	singleConfig.Combine0 = uint64(compressed0)
	singleConfig.Combine1 = uint64(compressed1)
	singleConfig.Combine2 = uint64(compressed2)
	singleConfig.RemainBigCount = uint32(remainBigCount)
	cfg.Configs = append(cfg.Configs, &singleConfig)

}

func GenerateClassicalCombines() {
	classicalConfigs := &classical_combine.Config{}
	distFile, _ := os.Create("classicalCombines.csv")
	writer := csv.NewWriter(distFile)
	defer writer.Flush()

	ccg := classical_combine_generator.ClassicalCombineGenerator{}
	ccg.OnGenerate()
	// TODO: 计算每一条剩余大牌数量,存在牌库里
	fmt.Println("ccg.ValidCombines size:", len(ccg.ValidCombines))
	var baseId int = 100000
	var controlFlag int = 0
	for ; controlFlag < 3; controlFlag++ {
		id := baseId + baseId*controlFlag
		for idx, combine := range ccg.ValidCombines {
			OnSaveClassicalCombine(id+idx, controlFlag, combine, writer, classicalConfigs)
		}
	}
	binaryFile, _ := os.Create("classical_combines.bin")
	marshal, _ := proto.Marshal(classicalConfigs)

	_, _ = binaryFile.Write(marshal)
}

func toAlphaString(i int) string {
	if i <= 0 {
		return ""
	}
	return string((i-1)%26 + 'A')
}

func main() {
	//GenerateSeqbombCombines()
	//GenerateBuxipaiCombines()
	// GenerateClassicalCombines()
	// babao_combine_generator.CombineTableGenerate()
	GenerateClassicalCombines()
}
