package test_settings_generator

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

const (
	handCardsNum  = 17 // 手牌数量
	tableCardsNum = 54 // 牌堆数量
)

var (
	// index是牌值 1->14, 2->15
	sCardsCountData = [18]uint32{0, 0, 0, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 1, 1}
	sCardsColors    = map[uint32][]uint32{
		0: {}, 1: {}, 2: {},
		3: {0, 1, 2, 3}, 4: {0, 1, 2, 3}, 5: {0, 1, 2, 3}, 6: {0, 1, 2, 3}, 7: {0, 1, 2, 3}, 8: {0, 1, 2, 3}, 9: {0, 1, 2, 3}, 10: {0, 1, 2, 3},
		11: {0, 1, 2, 3}, 12: {0, 1, 2, 3}, 13: {0, 1, 2, 3}, 14: {0, 1, 2, 3}, 15: {0, 1, 2, 3},
		16: {0xE},
		17: {0xF},
	}
)

func GetCardColor(card uint32) uint32 {
	if card == 0xE {
		return 0xE
	} else if card == 0xF {
		return 0xF
	}
	return card >> 4
}
func GetCardValue(card uint32) uint32 {
	if card == 0xE {
		return 16
	} else if card == 0xF {
		return 17
	}
	return card & 0xF
}

type CardsDealer struct {
	valsBucket   []uint32
	colorsBucket map[uint32][]uint32
}

func NewCardsDealer() *CardsDealer {
	cd := &CardsDealer{
		valsBucket:   make([]uint32, 18),
		colorsBucket: make(map[uint32][]uint32, 18),
	}
	copy(cd.valsBucket[:], sCardsCountData[:])
	for k, v := range sCardsColors {
		cd.colorsBucket[k] = make([]uint32, len(v))
		copy(cd.colorsBucket[k], v[:])
	}

	return cd
}

// 判断手牌是否还在牌堆中（避免重复）
// 判断手牌收否够14张，不够就补充

func (d *CardsDealer) CheckAndFillHandCards(item *TargetSetting) {
	// 本家手牌
	{
		tmp := make([]uint32, len(item.HandCardsOwn))
		copy(tmp, item.HandCardsOwn)
		item.HandCardsOwn = d.onCheckAndFill(tmp)
	}
	// 下家手牌
	{
		tmp := make([]uint32, len(item.HandCardsNext))
		copy(tmp, item.HandCardsNext)
		item.HandCardsNext = d.onCheckAndFill(tmp)
	}
	// 上家手牌
	{
		tmp := make([]uint32, len(item.HandCardsLast))
		copy(tmp, item.HandCardsLast)
		item.HandCardsLast = d.onCheckAndFill(tmp)
	}
	// 对家手牌
	if item.SeatNum == 4 {
		tmp := make([]uint32, len(item.HandCardsOpposite))
		copy(tmp, item.HandCardsOpposite)
		item.HandCardsOpposite = d.onCheckAndFill(tmp)
	}
}
func (d *CardsDealer) GenerateRemainCards(item *TargetSetting) {
	for {
		card := d.generateOneCard()
		if card == 0 {
			break
		}
		item.RemainCards = append(item.RemainCards, card)
	}
}

func (d *CardsDealer) onCheckAndFill(cards []uint32) []uint32 {
	var res []uint32
	for _, idx := range cards {
		if idx == 1 {
			idx = 14
		} else if idx == 2 {
			idx = 15
		}
		if d.valsBucket[idx] <= 0 {
			continue // val没有了，忽略
		}
		eachColors := d.colorsBucket[idx]
		if eachColors == nil || len(eachColors) == 0 {
			continue // color没有了，忽略
		}
		colorIdx := 0
		color := eachColors[0]

		var poker uint32
		if idx == 16 || idx == 17 {
			poker = color | 0x40
		} else {
			if idx == 14 {
				poker = (color << 4) | 1
			} else if idx == 15 {
				poker = (color << 4) | 2
			} else {
				poker = (color << 4) | idx
			}
		}
		res = append(res, poker)
		d.valsBucket[idx]--
		d.colorsBucket[idx] = slices.Delete(eachColors, colorIdx, colorIdx+1)
	}

	if len(res) > handCardsNum {
		panic(fmt.Errorf("手牌数量错误: %d, %d\n", len(res), handCardsNum))
	} else if len(res) < handCardsNum {
		// 补充
		diff := handCardsNum - len(res)
		for i := 0; i < diff; i++ {
			poker := d.generateOneCard()
			if poker == 0 {
				panic(fmt.Errorf("牌堆没有牌了"))
			}
			res = append(res, poker)
		}
	}

	return res
}
func (d *CardsDealer) generateOneCard() uint32 {
	var poker uint32 = 0
	var idx, color uint32

	var tmp []uint32 // 牌堆中剩余牌的索引
	for i, v := range d.valsBucket {
		if v > 0 {
			tmp = append(tmp, uint32(i))
		}
	}
	if len(tmp) <= 0 {
		return 0
	}
	// 随机选择一张牌
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	idx = tmp[r.Intn(len(tmp))]

	eachColors := d.colorsBucket[idx]
	if eachColors == nil || len(eachColors) == 0 {
		return 0
	}
	color = eachColors[0]
	if idx == 16 || idx == 17 {
		poker = color | 0x40
	} else {
		if idx == 14 {
			poker = (color << 4) | 1
		} else if idx == 15 {
			poker = (color << 4) | 2
		} else {
			poker = (color << 4) | idx
		}
	}
	d.valsBucket[idx]--
	d.colorsBucket[idx] = slices.Delete(eachColors, 0, 1)

	return poker
}

func (d *CardsDealer) CountCheck(item *TargetSetting) bool {
	length := len(item.HandCardsOwn) + len(item.HandCardsNext) + len(item.HandCardsLast) + len(item.HandCardsOpposite) + len(item.RemainCards)
	if length != tableCardsNum {
		fmt.Printf("手牌数量错误: ID[%d] length[%d], tableCardsNum[%d]\n", item.ID, length, tableCardsNum)
		return false
	}
	// poker重复性检查
	count := make(map[uint32]int)
	for _, card := range item.HandCardsOwn {
		if count[card] != 0 {
			fmt.Printf("手牌重复:  ID[%d] card[%d]\n", item.ID, card)
			return false
		}
		count[card]++
	}
	for _, card := range item.HandCardsNext {
		if count[card] != 0 {
			fmt.Printf("手牌重复:  ID[%d] card[%d]\n", item.ID, card)
			return false
		}
		count[card]++
	}
	for _, card := range item.HandCardsLast {
		if count[card] != 0 {
			fmt.Printf("手牌重复:  ID[%d] card[%d]\n", item.ID, card)
			return false
		}
	}
	for _, card := range item.HandCardsOpposite {
		if count[card] != 0 {
			fmt.Printf("手牌重复:  ID[%d] card[%d]\n", item.ID, card)
			return false
		}
	}
	for _, card := range item.RemainCards {
		if count[card] != 0 {
			fmt.Printf("手牌重复:  ID[%d] card[%d]\n", item.ID, card)
			return false
		}
	}
	return true
}
