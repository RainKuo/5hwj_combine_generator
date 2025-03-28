package combine_generator

import (
	"CombineGenerator/utils"
	"strconv"
)

type Generator struct {
	Cards        []int
	CardCombines []int
	Configs      *CombineConfigs
}

type Container struct {
	Cards        []int
	CardCombines []int
	CardStat     *HandCard
	ConfigID1st  uint32
	ConfigID2nd  uint32
	ConfigID3rd  uint32
	RemainCards  *HandCard
	EventID1st   string
	EventID2nd   string
	EventID3rd   string
}

func NewContainer() *Container {
	ctn := &Container{}
	ctn.Reset()
	return ctn
}

func (c *Container) Reset() {
	c.Cards = make([]int, 14)
	c.CardCombines = make([]int, 14)
	c.CardStat = &HandCard{}
	c.CardStat.Reset()
	c.RemainCards = &HandCard{}
	c.RemainCards.Reset()
	c.ConfigID1st = 0
	c.ConfigID2nd = 0
	c.ConfigID3rd = 0
}

func (c *Container) CopyFrom(ctn *Container) {
	copy(c.Cards, ctn.Cards)
	copy(c.CardCombines, ctn.CardCombines)
	c.CardStat.CopyFrom(ctn.CardStat)
}

func (c *Container) ValToStrList() []string {
	var str []string
	str = append(str, strconv.Itoa(int(c.ConfigID1st)))
	str = append(str, strconv.Itoa(int(c.ConfigID2nd)))
	str = append(str, strconv.Itoa(int(c.ConfigID3rd)))
	str = append(str, c.RemainCards.ToString())
	str = append(str, strconv.Itoa(c.CardStat.BombCount))
	str = append(str, strconv.Itoa(c.CardStat.KingBombCount))
	str = append(str, strconv.Itoa(c.CardStat.TripleCount))
	str = append(str, strconv.Itoa(c.CardStat.PairsCount))
	str = append(str, strconv.Itoa(c.CardStat.SingleCount))
	// TODO: 差事件ID
	return str
}

func (c *Container) GetTotalCount() int {
	sum := c.CardStat.GetTotal() + c.RemainCards.GetTotal()
	return sum
}

func (c *Container) CalcRemainCT() {
	for _, ct := range c.CardCombines {
		switch ct {
		case CombineTypeTriple:
			c.RemainCards.TripleCount++
		case CombineTypePair:
			c.RemainCards.PairsCount++
		case CombineTypeSingle:
			c.RemainCards.SingleCount++
		case CombineTypeKingBomb:
			c.RemainCards.KingBombCount++
		case CombineTypeBomb:
			c.RemainCards.BombCount++
		}
	}
}
func (c *Container) SetEventID(cc *CombineConfigs) {
	c.EventID1st = cc.GetConfigByID(c.ConfigID1st).EventID
	c.EventID2nd = cc.GetConfigByID(c.ConfigID2nd).EventID
	c.EventID3rd = cc.GetConfigByID(c.ConfigID3rd).EventID
}
func NewGenerator(path string) *Generator {
	generator := &Generator{
		Cards:   make([]int, 14),
		Configs: NewCombineConfigs(),
	}
	generator.Configs.Init(path)
	generator.Reset()
	return generator
}

func (g *Generator) Reset() {
	g.CardCombines = []int{
		CombineTypeBomb, CombineTypeBomb, CombineTypeBomb, CombineTypeBomb,
		CombineTypeBomb, CombineTypeBomb, CombineTypeBomb, CombineTypeBomb,
		CombineTypeBomb, CombineTypeBomb, CombineTypeBomb, CombineTypeBomb,
		CombineTypeBomb, CombineTypeKingBomb,
	}
	copy(g.Cards, CarCountList)
}

func (g *Generator) CopyFrom(ctn *Container) {
	copy(g.Cards, ctn.Cards)
	copy(g.CardCombines, ctn.CardCombines)
}

func (g *Generator) GeneratePlayer1(ctn *Container, configID uint32) bool {
	conf := g.Configs.GetConfigByID(configID)
	//fmt.Printf("ConfigID[%d], Bomb:[%d], KingBomb:[%d], Triple:[%d], Pair:[%d], Single:[%d] \n",
	//	conf.ID, conf.Bomb, conf.KingBomb, conf.Triple, conf.Pair, conf.Single)
	// 第一个玩家的牌肯定能凑齐,不用判断没有
	// 炸弹
	for i := 0; i < int(conf.Bomb); i++ {
		exist, idx := ctn.HasCombineType(CombineTypeBomb)
		if exist {
			ctn.Cards[idx] = 0
			ctn.SetCombineType(idx)
			ctn.CardStat.BombCount++
		}
	}
	// 王炸
	if conf.KingBomb > 0 {
		ctn.CardCombines[IndexJoker] = CombineTypeNone
		ctn.Cards[IndexJoker] = 0
		ctn.CardStat.KingBombCount++
	}
	// 三张
	for i := 0; i < int(conf.Triple); i++ {
		exist, idx := ctn.HasCombineType(CombineTypeTriple)
		if !exist {
			exist, idx = ctn.HasCombineType(CombineTypeBomb)
		}
		if exist {
			ctn.Cards[idx] -= 3
			ctn.SetCombineType(idx)
			ctn.CardStat.TripleCount++
		}
	}
	// 对子
	pairsMap := make(map[int]int)
	for i := 0; i < int(conf.Pair); i++ {
		exist, idx := ctn.HasCombineType(CombineTypePair)
		if _, find := pairsMap[idx]; find {
			exist = false
		}
		if !exist {
			exist, idx = ctn.HasCombineType(CombineTypeBomb)
		}
		if !exist {
			exist, idx = ctn.HasCombineType(CombineTypeTriple)
		}
		if exist {
			ctn.Cards[idx] -= 2
			ctn.SetCombineType(idx)
			ctn.CardStat.PairsCount++
		} else {
			return false
		}
	}

	singleMap := make(map[int]int)
	for i := 0; i < int(conf.Single); i++ {
		exist, idx := ctn.HasCombineType(CombineTypeSingle)
		if _, find := singleMap[idx]; find {
			exist = false
		}
		if !exist {
			exist, idx = ctn.HasCombineType(CombineTypeBomb)
		}
		if !exist {
			exist, idx = ctn.HasCombineType(CombineTypeTriple)
		}
		if _, find := singleMap[idx]; find {
			exist = false
		}
		if !exist {
			exist, idx = ctn.HasCombineType(CombineTypePair)
		}
		if exist {
			ctn.Cards[idx] -= 1
			ctn.SetCombineType(idx)
			ctn.CardStat.SingleCount++
		}
	}
	return true
}

func (g *Generator) GenerateOne(ctn *Container, configID uint32) bool {
	conf := g.Configs.GetConfigByID(configID)
	//fmt.Printf("ConfigID[%d], Bomb:[%d], KingBomb:[%d], Triple:[%d], Pair:[%d], Single:[%d] \n",
	//	conf.ID, conf.Bomb, conf.KingBomb, conf.Triple, conf.Pair, conf.Single)
	// 炸弹
	for i := 0; i < int(conf.Bomb); i++ {
		exist, idx := ctn.HasCombineType(CombineTypeBomb)
		if exist {
			ctn.Cards[idx] = 0
			ctn.SetCombineType(idx)
			ctn.CardStat.BombCount++
		} else {
			return false
		}
	}
	// 王炸
	if conf.KingBomb > 0 {
		ctn.CardCombines[IndexJoker] = CombineTypeNone
		ctn.Cards[IndexJoker] = 0
		ctn.CardStat.KingBombCount++
	}
	// 三张
	for i := 0; i < int(conf.Triple); i++ {
		exist, idx := ctn.HasCombineType(CombineTypeTriple)
		if !exist {
			exist, idx = ctn.HasCombineType(CombineTypeBomb)
		}
		if exist {
			ctn.Cards[idx] -= 3
			ctn.SetCombineType(idx)
			ctn.CardStat.TripleCount++
		} else {
			return false
		}
	}
	// 对子
	pairsMap := make(map[int]int)
	for i := 0; i < int(conf.Pair); i++ {
		exist, idx := ctn.HasCombineType(CombineTypePair)
		if _, find := pairsMap[idx]; find {
			exist = false
		}
		if !exist {
			exist, idx = ctn.HasCombineType(CombineTypeBomb)
		}
		if !exist {
			exist, idx = ctn.HasCombineType(CombineTypeTriple)
		}
		if exist {
			ctn.Cards[idx] -= 2
			ctn.SetCombineType(idx)
			ctn.CardStat.PairsCount++
		} else {
			return false
		}
	}
	// 单张
	singleMap := make(map[int]int)
	for i := 0; i < int(conf.Single); i++ {
		exist, idx := ctn.HasCombineType(CombineTypeSingle)
		if _, find := singleMap[idx]; find {
			exist = false
		}
		if !exist {
			exist, idx = ctn.HasCombineType(CombineTypeBomb)
		}
		if !exist {
			exist, idx = ctn.HasCombineType(CombineTypeTriple)
		}
		if _, find := singleMap[idx]; find {
			exist = false
		}
		if !exist {
			exist, idx = ctn.HasCombineType(CombineTypePair)
		}
		if exist {
			ctn.Cards[idx] -= 1
			ctn.SetCombineType(idx)
			ctn.CardStat.SingleCount++
		} else {
			return false
		}
	}
	return true
}

func (g *Generator) GenerateOther(ctn *Container, configID *uint32) bool {
	var configList []uint32
	for _, ct := range CombineList {
		_, exist := ctn.HasCombineType(ct)
		if exist > -1 {
			configList = g.Configs.GetConfigsByCombineType(uint32(ct))
			break
		}
	}
	if len(configList) > 0 {
		begin := false
		loopTimes := 0
		for !begin || loopTimes < 5 {
			loopTimes++
			begin = true
			randConfigID := RandomOneID(configList)
			tmpCtn := NewContainer()
			tmpCtn.CopyFrom(ctn)
			genSucc := g.GenerateOne(tmpCtn, randConfigID)
			if genSucc {
				ctn.CopyFrom(tmpCtn)
				*configID = randConfigID
				return true
			}
			utils.RemoveSliceItem(configList, randConfigID)
			if len(configList) == 0 {
				break
			}
		}
	}
	return false
}

func (g *Generator) GenerateTest() {
	ctn := NewContainer()
	copy(ctn.Cards, g.Cards)
	copy(ctn.CardCombines, g.CardCombines)
	g.GeneratePlayer1(ctn, 1014)
	g.GenerateOne(ctn, 1001)
	g.GenerateOne(ctn, 1005)
}

func (g *Generator) DoGenerate() (bool, *Container) {
	tmpAllIDList := make([]uint32, len(g.Configs.AllIDList))
	copy(tmpAllIDList, g.Configs.AllIDList)
	for {
		g.Reset()
		// 生成第一个玩家的手牌
		confID1st := RandomOneID(tmpAllIDList)
		ctn := NewContainer()
		copy(ctn.Cards, g.Cards)
		copy(ctn.CardCombines, g.CardCombines)
		if g.GeneratePlayer1(ctn, confID1st) {
			ctn.ConfigID1st = confID1st
			g.CopyFrom(ctn)
			// 生成第二个玩家手牌
			loop2nd := 0
			for loop2nd < 5 {
				loop2nd++
				copy(ctn.Cards, g.Cards)
				copy(ctn.CardCombines, g.CardCombines)
				if g.GenerateOther(ctn, &ctn.ConfigID2nd) {
					g.CopyFrom(ctn)
					// 生成第三个玩家手牌
					for loop3rd := 0; loop3rd < 5; loop3rd++ {
						copy(ctn.Cards, g.Cards)
						copy(ctn.CardCombines, g.CardCombines)
						if g.GenerateOther(ctn, &ctn.ConfigID3rd) {
							ctn.CalcRemainCT()
							ctn.SetEventID(g.Configs)
							return true, ctn
						}
					}
				}
			}
			utils.RemoveSliceItem(tmpAllIDList, confID1st)
		} else {
			utils.RemoveSliceItem(tmpAllIDList, confID1st)
		}
		if len(tmpAllIDList) == 0 {
			break
		}
	}
	return false, nil
}

func (g *Generator) HasCombineType(ct int) (bool, int) {
	for i := 0; i < g.CardCombines[i]; i++ {
		if g.CardCombines[i] == ct {
			return true, i
		}
	}
	return false, 0
}

func (g *Generator) SubCard(count uint32) {

}

func (g *Generator) SetCombineType(idx int) {
	ct := CombineTypeNone
	switch g.Cards[idx] {
	case 0:
		break
	case 1:
		ct = CombineTypeSingle
	case 2:
		if idx == IndexJoker {
			ct = CombineTypeKingBomb
		} else {
			ct = CombineTypePair
		}
	case 3:
		ct = CombineTypeTriple
	case 4:
		ct = CombineTypeBomb
	}
	g.CardCombines[idx] = ct
}

func (c *Container) SetCombineType(idx int) {
	ct := CombineTypeNone
	switch c.Cards[idx] {
	case 0:
		break
	case 1:
		ct = CombineTypeSingle
	case 2:
		if idx == IndexJoker {
			ct = CombineTypeKingBomb
		} else {
			ct = CombineTypePair
		}
	case 3:
		ct = CombineTypeTriple
	case 4:
		ct = CombineTypeBomb
	}
	c.CardCombines[idx] = ct
}

func (c *Container) HasCombineType(ct int) (bool, int) {
	for i := 0; i < len(c.CardCombines); i++ {
		if c.CardCombines[i] == ct {
			return true, i
		}
	}
	return false, -1
}
