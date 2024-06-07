package babao_combine_generator

import (
	"encoding/csv"
	"os"
	"strconv"
)

func CombineTableGenerate() {
	distFile, _ := os.Create("babaoCombines.csv")

	writer := csv.NewWriter(distFile)
	combineDatas := BabaoCombineGenerator()
	tableId := 0
	for controlFlag := 0; controlFlag < 3; controlFlag++ {
		tableId = 100000*(controlFlag+1) + (controlFlag + 1)
		for _, combine := range combineDatas {
			var row []string
			// row = append(row, strconv.Itoa(tableId), strconv.Itoa(controlFlag), strconv.Itoa(combine))
			row = append(row, strconv.Itoa(tableId), strconv.Itoa(controlFlag), combine.CombineToStr(), strconv.Itoa(combine.M), strconv.Itoa(combine.JokerCount), strconv.Itoa(combine.HuapaiCount))
			_ = writer.Write(row)
			tableId++
		}

	}

}
