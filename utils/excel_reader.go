package utils

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"strconv"
)

func ExcelToJson(path string) {
	xlxs, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}

	sheets := xlxs.GetSheetList()
	for _, sheetName := range sheets {
		allData := make([]map[string]interface{}, 0)
		rows, _ := xlxs.GetRows(sheetName)
		fields := rows[0]
		fieldTypes := rows[2]
		dataList := rows[7:]
		for _, cols := range dataList {

			data := make(map[string]interface{})
			for i := 0; i < len(fields); i++ {
				key := fields[i]
				val := cols[i]
				fieldType := fieldTypes[i]
				if key == "string" && val == "0" {
					val = ""
				}
				switch fieldType {
				case "uint32":
					tmp, _ := strconv.Atoi(val)
					data[key] = tmp
				case "string":
					data[key] = val
				default:
					data[key] = val
				}

			}
			allData = append(allData, data)
		}

		if str, err := json.Marshal(allData); err == nil {
			confPath := fmt.Sprintf("res/config/config_%s.json", sheetName)
			file, err := os.Create(confPath)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			file.Write(str)
			file.Close()
		}

	}
}
