package babao_combine_generator

import (
	"fmt"
	"strings"
)

type OneBabaoCombine struct {
	Combines    []int
	M           int
	JokerCount  int
	HuapaiCount int
}

func (c *OneBabaoCombine) CombineToStr() string {
	str := fmt.Sprintf("%v", c.Combines)
	str = strings.ReplaceAll(str, " ", ",")
	str = strings.Trim(str, "\"")
	return str
}

func BabaoCombineGenerator() []OneBabaoCombine {
	babaoCombineDatas := make([]OneBabaoCombine, 0)
	finalCombine := make([]int, 0)
	combine_types := [4]int{0, 1, 2, 3}
	mVal := [3]int{5, 6, 7}
	jokers := [4]int{2, 3, 4, 5}
	huapai := [4]int{2, 3, 4, 5}
	for _, combine1 := range combine_types {
		data1 := combine1
		for _, combine2 := range combine_types {
			data2 := (data1 << 4) | combine2
			for _, combine3 := range combine_types {
				data3 := (data2 << 4) | combine3
				for _, combine4 := range combine_types {
					data4 := (data3 << 4) | combine4
					for _, m := range mVal {
						data5 := (data4 << 4) | m
						for _, jokerCount := range jokers {
							data6 := (data5 << 4) | jokerCount
							for _, hua := range huapai {
								data7 := (data6 << 4) | hua
								finalCombine = append(finalCombine, data7)
								obc := OneBabaoCombine{
									Combines:    make([]int, 0),
									M:           m,
									JokerCount:  jokerCount,
									HuapaiCount: hua,
								}
								obc.Combines = append(obc.Combines, combine1, combine2, combine3, combine4)
								babaoCombineDatas = append(babaoCombineDatas, obc)
							}
						}
					}
				}
			}
		}
	}
	return babaoCombineDatas
}
