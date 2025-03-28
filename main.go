package main

import (
	"CombineGenerator/classical_combine_generator"
	"CombineGenerator/combine_generator"
	"CombineGenerator/db"
	"CombineGenerator/landlord_4_combine_generator"
	"CombineGenerator/proto/out/classical_combine"
<<<<<<< HEAD
=======
	"CombineGenerator/proto/out/landlord_4p_combines"
>>>>>>> 8245a9e27f6b100ede62a37303f1b1e77ac77d9e
	"CombineGenerator/utils"
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/golang/protobuf/proto"
)

const (
	SEQBOMB_BOMB_2_3  = 1 // 3个ID都只有2-3炸
	SEQBOMB_BOMB_1    = 2 // 至少有1个ID只有1炸
	SEQBOMB_BOMB_4    = 3 // 至少有1个ID有4炸
	SEQBOMB_BOMB_NONE = 4 // 没有炸弹
)

const (
	BUXIPAI_BOMB_0 = 0 // 所有玩家只有0炸/1炸的牌
	BUXIPAI_BOMB_2 = 2 // 单个玩家最多只有2炸的牌
	BUXIPAI_OTHER  = 3 // 其他类型
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

func OnSaveSeqbombCombines(ctn *combine_generator.Container, writer *csv.Writer, configID int, bombType int, controlFlag int, group_intensity *combine_generator.GroupIntensity) {
	str := []string{strconv.Itoa(int(ctn.ConfigID1st)), strconv.Itoa(int(ctn.ConfigID2nd)),
		strconv.Itoa(int(ctn.ConfigID3rd))}
	exKey := strings.Join(str, "")
	if _, ok := existed[exKey]; !ok {
		var row []string
		row = append(row, strconv.Itoa(configID), strconv.Itoa(controlFlag), strconv.Itoa(bombType), strconv.Itoa(int(ctn.ConfigID1st)), strconv.Itoa(int(ctn.ConfigID2nd)),
			strconv.Itoa(int(ctn.ConfigID3rd)), ctn.RemainCards.ToString(), strconv.Itoa(ctn.CardStat.BombCount),
			strconv.Itoa(ctn.CardStat.KingBombCount), strconv.Itoa(ctn.CardStat.TripleCount),
			strconv.Itoa(ctn.CardStat.PairsCount), strconv.Itoa(ctn.CardStat.SingleCount),
			ctn.EventID1st, ctn.EventID2nd, ctn.EventID3rd,
			group_intensity.ToString())
		err := writer.Write(row)
		existed[exKey] = 1
		if err != nil {
			fmt.Println(err)
		}
	}
}

func GenerateSeqbombCombines() {
	utils.ExcelToJson("excel/连炸基础牌型.xlsx")

	generator := combine_generator.NewGenerator("./res/config/config_seqbomb_base_combines.json")
	distFile, _ := os.Create("seqbombCombines_tmp.csv")

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
				// 整合牌组的强度
				group_intensity := combine_generator.GroupIntensity{
					generator.Configs.ConfigIDMap[ctn.ConfigID1st].Intensity,
					generator.Configs.ConfigIDMap[ctn.ConfigID2nd].Intensity,
					generator.Configs.ConfigIDMap[ctn.ConfigID3rd].Intensity,
				}
				group_intensity.Sort()

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
					OnSaveSeqbombCombines(ctn, writer, idBegin+i, SEQBOMB_BOMB_1, controlFlag, &group_intensity)
				}
				if bomb23 {
					OnSaveSeqbombCombines(ctn, writer, idBegin+i, SEQBOMB_BOMB_2_3, controlFlag, &group_intensity)
				}
				if bomb4 {
					OnSaveSeqbombCombines(ctn, writer, idBegin+i, SEQBOMB_BOMB_4, controlFlag, &group_intensity)
				}
				if !bomb1 && !bomb23 && !bomb4 {
					OnSaveSeqbombCombines(ctn, writer, idBegin+i, SEQBOMB_BOMB_NONE, controlFlag, &group_intensity)
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
				bomb2 := false
				bomb01 := true
				other := false
				// bomb3 := false
				var bombs uint32 = 0

				for _, id := range ids {
					conf := generator.Configs.ConfigIDMap[id]
					if conf.Bomb == 0 || conf.Bomb == 1 {
						// bomb01 = true
					} else if conf.Bomb <= 2 && !other {
						bomb2 = true
						bomb01 = false
					} else {
						other = true
						bomb01 = false
						bomb2 = false
					}
					bombs += conf.Bomb
				}

				if bomb01 {
					OnSave(ctn, writer, idBegin+i, BUXIPAI_BOMB_0, controlFlag)
				} else if bomb2 {
					OnSave(ctn, writer, idBegin+i, BUXIPAI_BOMB_2, controlFlag)
				} else {
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

func IntsToString(ints []uint32) string {
	strArr := make([]string, len(ints))
	for i := 0; i < len(ints); i++ {
		strArr[i] = strconv.Itoa(int(ints[i]))
	}
	return strings.Join(strArr, ",")
}

func OnSaveLandlord4Combines(id int, writer *csv.Writer, cfg *landlord_4p_combines.Landlord_4PCombineConfig) {
	var row []string
	row = append(row, strconv.Itoa(id), IntsToString(cfg.Player1HandCards), IntsToString(cfg.Player2HandCards), IntsToString(cfg.Player3HandCards), IntsToString(cfg.Player4HandCards),
		IntsToString(cfg.BombMultis), IntsToString(cfg.Turns), IntsToString(cfg.Scores))
	_ = writer.Write(row)
}

func GenerateLandlord4pCombines() {
	landlord4Configs := &landlord_4p_combines.Config{}
	landlord_4_combine_generator.Init()
	landlord_4_combine_generator.Landlord4CombineGenerate()
	distFile, _ := os.Create("landlord4Combines.csv")
	writer := csv.NewWriter(distFile)
	defer writer.Flush()

	existed = make(map[string]int)
	id := 1000001
	for id < 1100000 {
		cg4 := landlord_4_combine_generator.Landlord4CombineGenerate()
		if cg4 == nil {
			continue
		}
		if cg4.IsValid() {
			cfg := &landlord_4p_combines.Landlord_4PCombineConfig{}
			cg4.FillProto(cfg)
			marshaled, _ := proto.Marshal(cfg)
			hash := sha256.New()
			hash.Write(marshaled)
			hashBytes := hash.Sum(nil)
			hashString := hex.EncodeToString(hashBytes)
			if _, ok := existed[hashString]; !ok {
				existed[hashString] = 1
				landlord4Configs.Configs = append(landlord4Configs.Configs, cfg)
				OnSaveLandlord4Combines(id, writer, cfg)
				id++
			}

		}
	}
	binaryFile, _ := os.Create("landlord4_combines.bin")
	marshal, _ := proto.Marshal(landlord4Configs)

	_, _ = binaryFile.Write(marshal)
	fmt.Println("Valid config count: ", len(landlord4Configs.Configs))
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
	GenerateSeqbombCombines()
	// GenerateBuxipaiCombines()
	// GenerateClassicalCombines()
	// babao_combine_generator.CombineTableGenerate()
	// GenerateClassicalCombines()
	// qx_combine_generator.CombineTableGenerate()
	//tower_combine_generator.CombineTableGenerate()
	// test_settings_generator.TestSettingsGenerate()
	// tower_combine_generator.CombineTableGenerate()
	// GenerateLandlord4pCombines()
}
