package qx_combine_generator

type OneQXCombine struct {
	JokerCount         int // 赢家王牌数量
	MoIntensity        int // 摸牌强度
	SynthesisIntensity int // 连续合成强度
	EryuanClone        int // 是否有二元克隆(0无 1有)
}

func QXCombineGenerator() []OneQXCombine {
	qxCombineDatas := make([]OneQXCombine, 0)
	winnerJokerCount := [4]int{2, 3, 4, 5}
	winnerMoIntens := [3]int{3, 4, 5}  // 赢家摸牌强度
	winnerSynIntens := [3]int{3, 4, 5} // 赢家连续合成强度
	otherMoIntens := [2]int{1, 2}
	otherSynIntens := [2]int{1, 2}
	eryuanClones := [2]int{0, 1}
	for _, jokerCount := range winnerJokerCount {
		for _, wmi := range winnerMoIntens {
			for _, wsi := range winnerSynIntens {
				for _, omi1 := range otherMoIntens {
					for _, omi2 := range otherMoIntens {
						for _, omi3 := range otherMoIntens {
							for _, osi1 := range otherSynIntens {
								for _, osi2 := range otherSynIntens {
									for _, osi3 := range otherSynIntens {
										for _, eryuanClone := range eryuanClones {
											combine := &OneQXCombine{
												JokerCount:         0,
												MoIntensity:        0,
												SynthesisIntensity: 0,
												EryuanClone:        0,
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
											combine.EryuanClone = eryuanClone
											qxCombineDatas = append(qxCombineDatas, *combine)
										}

									}
								}
							}
						}
					}
				}
			}
		}
	}
	return qxCombineDatas
}
