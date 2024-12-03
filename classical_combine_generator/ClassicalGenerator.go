package classical_combine_generator

import "math/rand"

type ValidCombine []BaseCombine

type ClassicalCombineGenerator struct {
	ValidCombines []ValidCombine
}

func (ccg *ClassicalCombineGenerator) OnGenerate() {
	ccg.GenerateFirst()
}

func (ccg *ClassicalCombineGenerator) GenerateFirst() {
	baseCombines := make([]BaseCombine, 3)
	bigsCount := []int64{2, 3, 4, 5, 6}
	bombsCount := []int64{0, 1}
	triplesCount := []int64{0, 1, 2}
	pairsCount := []int64{0, 3}
	for _, bigCount := range bigsCount {
		baseCombines[0].BigCount = bigCount
		for _, pairCount := range pairsCount {
			baseCombines[0].Pair = pairCount
			for _, tripleCount := range triplesCount {
				baseCombines[0].Triple = tripleCount
				for _, bombCount := range bombsCount {
					baseCombines[0].Bomb = bombCount
					if !baseCombines[0].IsValid() {
						continue
					}
					baseCombines[0].CalcSingle()
					if baseCombines[0].Single < 4 {
						continue
					}
					ccg.GenerateSecond(baseCombines)
				}
			}
		}
	}
}

func (ccg *ClassicalCombineGenerator) GenerateSecond(_baseCombines []BaseCombine) {
	baseCombines := make([]BaseCombine, len(_baseCombines))
	copy(baseCombines, _baseCombines)
	bigsCount := []int64{0, 1, 2}
	bombsCount := []int64{0, 1}
	triplesCount := []int64{0, 1, 2}
	pairsCount := []int64{0, 3}
	for _, bigCount := range bigsCount {
		baseCombines[1].BigCount = bigCount
		if baseCombines[0].BigCount+baseCombines[1].BigCount > 6 {
			continue
		}
		for _, pairCount := range pairsCount {
			baseCombines[1].Pair = pairCount
			for _, tripleCount := range triplesCount {
				baseCombines[1].Triple = tripleCount
				for _, bombCount := range bombsCount {
					baseCombines[1].Bomb = bombCount
					if !baseCombines[1].IsValid() {
						continue
					}
					baseCombines[1].CalcSingle()
					if baseCombines[1].Single < 4 {
						continue
					}
					ccg.GenerateThird(baseCombines)
				}
			}
		}
	}
}

func (ccg *ClassicalCombineGenerator) GenerateThird(_baseCombines []BaseCombine) {
	baseCombines := make([]BaseCombine, len(_baseCombines))
	copy(baseCombines, _baseCombines)
	bombsCount := []int64{0, 1}
	triplesCount := []int64{0, 1, 2}
	pairsCount := []int64{0, 3}

	for _, pairCount := range pairsCount {
		baseCombines[2].Pair = pairCount
		for _, tripleCount := range triplesCount {
			baseCombines[2].Triple = tripleCount
			for _, bombCount := range bombsCount {
				baseCombines[2].Bomb = bombCount
				if !baseCombines[2].IsValid() {
					continue
				}
				baseCombines[2].CalcSingle()
				if baseCombines[2].Single < 4 {
					continue
				}
				// 事件验证
				ccg.CheckEvent(baseCombines)
			}
		}
	}

}

func CheckEventSingleSequence(baseCombine BaseCombine) bool {
	return baseCombine.Single >= 5
}

func CheckEventThreePairSequence(baseCombine BaseCombine) bool {
	return baseCombine.Pair >= 3
}

func CheckEventTwoTripleSequence(baseCombine BaseCombine) bool {
	return baseCombine.Triple >= 2
}

type CheckFunc func(baseCombines BaseCombine) bool

var EventCheckers = map[int64]CheckFunc{
	EVENT_NONE:                nil,
	EVENT_SINGLE_SEQUENCE_5:   CheckEventSingleSequence,
	EVENT_THREE_PAIR_SEQUENCE: CheckEventThreePairSequence,
	EVENT_TWO_TRIPLE_SEQUENCE: CheckEventTwoTripleSequence,
}

var EventSort = []int64{EVENT_NONE, EVENT_SINGLE_SEQUENCE_5, EVENT_THREE_PAIR_SEQUENCE, EVENT_TWO_TRIPLE_SEQUENCE}

func (ccg *ClassicalCombineGenerator) CheckEvent(_baseCombines []BaseCombine) {
	baseCombines := make([]BaseCombine, len(_baseCombines))
	copy(baseCombines, _baseCombines)
	for _, event0 := range EventSort {
		checker0 := EventCheckers[event0]
		if checker0 == nil {
			baseCombines[0].Event = EVENT_NONE
		} else {
			if !checker0(baseCombines[0]) {
				continue
			}
			baseCombines[0].Event = event0
		}

		for _, event1 := range EventSort {
			checker1 := EventCheckers[event1]
			if checker1 == nil {
				baseCombines[1].Event = EVENT_NONE
			} else {
				if !checker1(baseCombines[1]) {
					continue
				}
				baseCombines[1].Event = event1
			}
			for _, event2 := range EventSort {
				checker2 := EventCheckers[event2]
				if checker2 == nil {
					baseCombines[2].Event = EVENT_NONE
				} else {
					if !checker2(baseCombines[2]) {
						continue
					}
					baseCombines[2].Event = event2
				}
				if rand.Intn(100) <= 25 {
					tmpCombine := make([]BaseCombine, len(baseCombines))
					copy(tmpCombine, baseCombines)
					ccg.ValidCombines = append(ccg.ValidCombines, tmpCombine)
				}
			}
		}
	}
}

func GetRivalType(baseCombines []BaseCombine) int {
	baseRivalType := NORMAL_NO_RIVAL
	pairSequenceCount := 0
	bombCount := 0
	singleSequenceCount := 0
	twoTripleCount := 0
	tripleCount := 0
	for _, combine := range baseCombines {
		if combine.Event == EVENT_THREE_PAIR_SEQUENCE {
			pairSequenceCount++
		}
		if combine.Event == EVENT_TWO_TRIPLE_SEQUENCE {
			twoTripleCount++
		}
		if combine.Event == EVENT_SINGLE_SEQUENCE_5 {
			singleSequenceCount++
		}
		if combine.Bomb > 0 {
			bombCount++
		}
		if combine.Triple > 0 {
			tripleCount++
		}
	}
	if (pairSequenceCount >= 2 || singleSequenceCount >= 2 || tripleCount >= 3) && baseCombines[0].BigCount >= 5 {
		baseRivalType = RARE_RIVAL
	} else if bombCount >= 2 || twoTripleCount >= 2 {
		baseRivalType = RARE_RIVAL
	} else if pairSequenceCount >= 2 || singleSequenceCount >= 2 || tripleCount >= 3 {
		baseRivalType = NORMAL_RIVAL
	} else if baseCombines[0].BigCount >= 5 {
		baseRivalType = RATE_NO_RIVAL
	} else {
		baseRivalType = NORMAL_NO_RIVAL
	}
	return baseRivalType
}
