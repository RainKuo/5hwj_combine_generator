package test_settings_generator

import (
	"CombineGenerator/utils"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 牌值映射
var pokerMap = map[string]uint32{
	"3":      3,
	"4":      4,
	"5":      5,
	"6":      6,
	"7":      7,
	"8":      8,
	"9":      9,
	"10":     10,
	"J":      11,
	"Q":      12,
	"K":      13,
	"A":      14,
	"2":      15,
	"Joker":  16,
	"Joker2": 17,
}

type RawSetting struct {
	ID                uint32
	GameID            uint32
	HandCardsOwn      string
	HandCardsLast     string
	HandCardsNext     string
	HandCardsOpposite string
	StartCall         uint32
	Comment           string
	SeatNum           uint32
}
type TargetSetting struct {
	ID                uint32 // 配置id
	GameID            uint32
	HandCardsOwn      []uint32 // 本家手牌
	HandCardsLast     []uint32 // 上家手牌
	HandCardsNext     []uint32 // 下家手牌
	HandCardsOpposite []uint32 // 对家手牌
	RemainCards       []uint32 // 底牌
	StartCall         uint32   //  初始起叫
	Comment           string   // 备注信息
	SeatNum           uint32
}

func TestSettingsGenerate() {
	utils.ExcelToJson("excel/连炸牌型测试配置表.xlsx")

	file, err := os.Open("res/config/config_test_settings.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	var raw []RawSetting
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&raw)
	if err != nil {
		fmt.Println(err.Error())
	}

	var settings []TargetSetting // target
	for _, setting := range raw {
		dealer := NewCardsDealer()
		item := TargetSetting{}
		item.ID = setting.ID
		item.GameID = setting.GameID
		item.StartCall = setting.StartCall
		item.Comment = setting.Comment
		item.SeatNum = setting.SeatNum
		item.HandCardsOwn = ExtractNumbers(setting.HandCardsOwn)
		item.HandCardsLast = ExtractNumbers(setting.HandCardsLast)
		item.HandCardsNext = ExtractNumbers(setting.HandCardsNext)
		if setting.SeatNum == 4 {
			item.HandCardsOpposite = ExtractNumbers(setting.HandCardsOpposite)
		}
		dealer.CheckAndFillHandCards(&item)
		dealer.GenerateRemainCards(&item)
		if dealer.CountCheck(&item) {
			settings = append(settings, item)
		}
	}

	//Save(settings)
	Append(settings)
}

func ExtractNumbers(str string) []uint32 {
	if str == "" {
		return nil
	}
	var numbers []uint32
	parts := strings.Split(str, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		poker, ok := pokerMap[part]
		if !ok {
			panic(fmt.Errorf("无效的牌值: %s", part))
		}
		numbers = append(numbers, poker)

		//if strings.HasPrefix(part, "0x") {
		//	// 16进制
		//	num, err := strconv.ParseInt(part[2:], 16, 32)
		//	if err != nil {
		//		return nil, fmt.Errorf("解析16进制数字失败: %w", err)
		//	}
		//	numbers = append(numbers, uint8(num))
		//} else {
		//	// 10进制
		//	num, err := strconv.ParseInt(part, 10, 32)
		//	if err != nil {
		//		return nil, fmt.Errorf("解析10进制数字失败: %w", err)
		//	}
		//	numbers = append(numbers, uint8(num))
		//}
	}
	return numbers
}

// 覆盖
func Save(settings []TargetSetting) {
	existed := make(map[uint32]any, len(settings))
	targetFile, _ := os.Create("config_target_test_settings.csv")
	defer targetFile.Close()
	writer := csv.NewWriter(targetFile)
	defer writer.Flush()

	for _, item := range settings {
		if _, ok := existed[item.ID]; ok {
			continue
		}
		var row []string
		row = append(row, strconv.Itoa(int(item.ID)), strconv.Itoa(int(item.GameID)), ToString(item.HandCardsOwn), ToString(item.HandCardsLast), ToString(item.HandCardsNext), ToString(item.HandCardsOpposite),
			ToString(item.RemainCards), strconv.Itoa(int(item.StartCall)), item.Comment)
		writer.Write(row)
	}
}

// 追加
func Append(settings []TargetSetting) {
	existed := make(map[uint32]bool)
	filePath := "config_target_test_settings.csv"

	// 检查文件是否存在
	if _, err := os.Stat(filePath); err == nil {
		file, err := os.Open(filePath)
		if err != nil {
			return
		}
		defer file.Close()

		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			return
		}

		// 将现有记录的ID存储到映射中
		for _, record := range records {
			if len(record) > 0 {
				id, err := strconv.Atoi(record[0])
				if err != nil {
					return
				}
				existed[uint32(id)] = true
			}
		}
	}

	// 以追加模式打开文件
	targetFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer targetFile.Close()

	writer := csv.NewWriter(targetFile)
	defer writer.Flush()

	// 遍历settings，只写入不存在的记录
	for _, item := range settings {
		if _, ok := existed[item.ID]; ok {
			continue
		}
		row := []string{
			strconv.Itoa(int(item.ID)),
			strconv.Itoa(int(item.GameID)),
			ToString(item.HandCardsOwn),
			ToString(item.HandCardsLast),
			ToString(item.HandCardsNext),
			ToString(item.HandCardsOpposite),
			ToString(item.RemainCards),
			strconv.Itoa(int(item.StartCall)),
			item.Comment,
		}
		if err := writer.Write(row); err != nil {
			return
		}
		existed[item.ID] = true // 标记已存在
	}
}

func ToString(arr []uint32) string {
	if len(arr) == 0 {
		return ""
	}
	jsonBytes, err := json.Marshal(arr)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}
