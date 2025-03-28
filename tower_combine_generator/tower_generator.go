package tower_combine_generator

type OneQXCombine struct {
	JokerCount   int // 赢家王牌数量
	GoldCount    int // 赢家金牌数量
	MoIntensity  int // 摸牌强度
	BaoIntensity int // 连续爆强度
}

func QXCombineGenerator() []OneQXCombine {
	qxCombineDatas := make([]OneQXCombine, 0)
	winnerJokerCount := [2]int{3, 4}   // 赢家王牌总数
	winnerGoldCount := [3]int{3, 4, 5} // 赢家金牌总数
	winnerMoIntens := [3]int{3, 4, 5}  // 赢家摸牌强度
	winnerBaoIntens := [3]int{3, 4, 5} // 赢家连续爆强度
	otherMoIntens := [2]int{1, 2}
	otherBaoIntens := [2]int{1, 2}
	for _, jokerCount := range winnerJokerCount {
		for _, wmi := range winnerMoIntens {
			for _, wsi := range winnerBaoIntens {
				for _, omi1 := range otherMoIntens {
					for _, omi2 := range otherMoIntens {
						for _, omi3 := range otherMoIntens {
							for _, osi1 := range otherBaoIntens {
								for _, osi2 := range otherBaoIntens {
									for _, osi3 := range otherBaoIntens {
										for _, wgc := range winnerGoldCount {
											combine := &OneQXCombine{
												JokerCount:   0,
												MoIntensity:  0,
												BaoIntensity: 0,
												GoldCount:    0,
											}
											tmpMoIntensity := wmi
											tmpMoIntensity = tmpMoIntensity<<3 | omi1
											tmpMoIntensity = tmpMoIntensity<<3 | omi2
											tmpMoIntensity = tmpMoIntensity<<3 | omi3
											combine.MoIntensity = tmpMoIntensity

											tmpBaoIntensity := wsi
											tmpBaoIntensity = tmpBaoIntensity<<3 | osi1
											tmpBaoIntensity = tmpBaoIntensity<<3 | osi2
											tmpBaoIntensity = tmpBaoIntensity<<3 | osi3
											combine.BaoIntensity = tmpBaoIntensity

											combine.GoldCount = wgc

											combine.JokerCount = jokerCount
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
