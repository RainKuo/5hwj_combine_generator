package combine_generator

import "strconv"

type HandCardStat struct {
	ConfigID1st   int // 玩家1配置id
	ConfigID2nd   int // 玩家2配置id
	ConfigID3rd   int // 玩家3配置id
	BombCount     int
	KingBombCount int
	TripleCount   int
	PairsCount    int
	SingleCount   int
	RemainCards   []int // 剩余牌
}

type HandCard struct {
	BombCount     int
	KingBombCount int
	TripleCount   int
	PairsCount    int
	SingleCount   int
}

// 基本牌型枚举
const CombineTypeNone = 0
const CombineTypeTriple = 1   // 三张
const CombineTypePair = 2     // 对子
const CombineTypeSingle = 3   // 单牌
const CombineTypeBomb = 4     // 炸弹
const CombineTypeKingBomb = 5 // 王炸

const IndexJoker = 13

// CombineList 牌型列表(按优先级排序)
var CombineList []int

// 牌数列表,index从3到王
var CarCountList []int

func init() {
	CombineList = []int{CombineTypeTriple, CombineTypePair, CombineTypeSingle, CombineTypeBomb, CombineTypeKingBomb}
	CarCountList = []int{
		4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 2,
	}
}

func (hc *HandCard) Reset() {
	hc.TripleCount = 0
	hc.PairsCount = 0
	hc.SingleCount = 0
	hc.BombCount = 0
	hc.KingBombCount = 0
}

func (hc *HandCard) CopyFrom(h *HandCard) {
	hc.TripleCount = h.TripleCount
	hc.PairsCount = h.PairsCount
	hc.SingleCount = h.SingleCount
	hc.BombCount = h.BombCount
	hc.KingBombCount = h.KingBombCount
}

func (hc *HandCard) ToString() string {
	var str string
	str = str + strconv.Itoa(hc.BombCount) + ","
	str = str + strconv.Itoa(hc.KingBombCount) + ","
	str = str + strconv.Itoa(hc.TripleCount) + ","
	str = str + strconv.Itoa(hc.PairsCount) + ","
	str = str + strconv.Itoa(hc.SingleCount)
	return str
}
