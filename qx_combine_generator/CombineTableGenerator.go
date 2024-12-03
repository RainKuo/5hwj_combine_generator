package qx_combine_generator

import (
	"encoding/csv"
	"os"
	"strconv"
)

func CombineTableGenerate() {
	distFile, _ := os.Create("QXCombines.csv")

	writer := csv.NewWriter(distFile)
	combineDatas := QXCombineGenerator()
	tableId := 0
	for controlFlag := 0; controlFlag < 3; controlFlag++ {
		tableId = 100000*(controlFlag+1) + 1
		for _, combine := range combineDatas {
			var row []string
			// row = append(row, strconv.Itoa(tableId), strconv.Itoa(controlFlag), strconv.Itoa(combine))
			row = append(row, strconv.Itoa(tableId), strconv.Itoa(controlFlag), strconv.Itoa(combine.JokerCount),
				strconv.Itoa(combine.MoIntensity), strconv.Itoa(combine.SynthesisIntensity), strconv.Itoa(combine.EryuanClone))
			_ = writer.Write(row)
			tableId++
		}
		println(tableId)
	}
	writer.Flush()
}
