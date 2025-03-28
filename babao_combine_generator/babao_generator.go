package babao_combine_generator

type OneBabaoCombine struct {
	JokerCount         int // 赢家王牌数量
	MoIntensity        int // 摸牌强度
	SynthesisIntensity int // 连续合成强度
}

func BabaoCombineGenerator() []OneBabaoCombine {
	babaoCombineDatas := make([]OneBabaoCombine, 0)
	winnerJokerCount := [4]int{2, 3, 4, 5}
	winnerMoIntens := [3]int{3, 4, 5}  // 赢家摸牌强度
	winnerSynIntens := [3]int{3, 4, 5} // 赢家连续合成强度
	otherMoIntens := [2]int{1, 2}
	otherSynIntens := [2]int{1, 2}
	for _, jokerCount := range winnerJokerCount {
		for _, wmi := range winnerMoIntens {
			for _, wsi := range winnerSynIntens {
				for _, omi1 := range otherMoIntens {
					for _, omi2 := range otherMoIntens {
						for _, omi3 := range otherMoIntens {
							for _, osi1 := range otherSynIntens {
								for _, osi2 := range otherSynIntens {
									for _, osi3 := range otherSynIntens {
										combine := &OneBabaoCombine{
											JokerCount:         0,
											MoIntensity:        0,
											SynthesisIntensity: 0,
										}
										tmpMoIntensity := wmi
										tmpMoIntensity = tmpMoIntensity<<3 | omi1
										tmpMoIntensity = tmpMoIntensity<<3 | omi2
										tmpMoIntensity = tmpMoIntensity<<3 | omi3
										combine.MoIntensity = tmpMoIntensity

										tmpSynIntensity := wsi
										tmpSynIntensity = tmpSynIntensity<<3 | osi1
										tmpSynIntensity = tmpSynIntensity<<3 | osi2
										tmpSynIntensity = tmpSynIntensity<<3 | osi3
										combine.SynthesisIntensity = tmpSynIntensity

										combine.JokerCount = jokerCount
										babaoCombineDatas = append(babaoCombineDatas, *combine)
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return babaoCombineDatas
}
