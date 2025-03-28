package landlord_4_combine_generator

import "CombineGenerator/utils"

type CardsAnalyzer struct {
	counts [17]int
	values [12]int
}

func (analyzer *CardsAnalyzer) Init(cards [MAX_CARD_COUNT_LENGTH]int) {
	for i := 5; i < len(cards); i++ {
		analyzer.counts[i] = cards[i]
		for count := 1; count <= cards[i]; count++ {
			analyzer.values[count] = analyzer.values[count] | 1<<i
		}
	}
}

func (analyzer *CardsAnalyzer) Indices(count int) int {
	return analyzer.values[count]
}

func (analyzer *CardsAnalyzer) FindRepeatSingle(repeatCount int) int {
	indices := analyzer.Indices(repeatCount)
	validIndex := make([]int, 0)
	if indices == 0 {
		return -1
	}
	curIndex := 0
	for indices != 0 {
		if indices&1 == 1 {
			validIndex = append(validIndex, curIndex)
		}
		curIndex++
		indices >>= 1
	}
	if len(validIndex) != 0 {
		randIndex := utils.RandInt(0, len(validIndex))
		return validIndex[randIndex]
	}
	return -1
}

func (analyzer *CardsAnalyzer) FindSequence(starCount int, seqCount int) int {
	if seqCount == 0 {
		if analyzer.counts[16] >= starCount {
			return 0
		}
		return -1
	}
	indices := analyzer.Indices(starCount)
	validBeginIndex := make([]int, 0)
	mask := 1<<seqCount - 1
	if indices == 0 {
		return -1
	}
	curIndex := 0
	for indices != 0 && curIndex < 14-seqCount+1 {
		if indices&mask == mask {
			validBeginIndex = append(validBeginIndex, curIndex)
		}
		curIndex++
		indices >>= 1
	}
	if len(validBeginIndex) != 0 {
		randIndex := utils.RandInt(0, len(validBeginIndex))
		return validBeginIndex[randIndex]
	}
	return -1
}
