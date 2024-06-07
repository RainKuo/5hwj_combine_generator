package classical_combine_generator

const (
	EVENT_NONE                = 0    // 没有事件
	EVENT_SINGLE_SEQUENCE_5   = 1001 // 5连顺
	EVENT_TWO_TRIPLE_SEQUENCE = 1002 // 2连飞机
	EVENT_THREE_PAIR_SEQUENCE = 1005 // 3连对
)

const (
	NORMAL_RIVAL    = 0 //    1. 常见有对抗：2家或3家有连对、顺子，3家有三张
	NORMAL_NO_RIVAL = 1 //    2. 常见无对抗：i、iii、iv 之外的发牌
	RARE_RIVAL      = 2 //    3. 稀有有对抗：2家或3家有炸弹，飞机；满足条件i且（第一家有大于等于5张大牌）
	RATE_NO_RIVAL   = 3 //    4. 稀有无对抗：（第一家有大于等于5张大牌）但不满足i、iii条件。
)

type BaseCombine struct {
	BigCount  int64
	Bomb      int64
	Triple    int64
	Pair      int64
	Single    int64
	Event     int64
	RivalType int64
}

func (c *BaseCombine) CalcSingle() {
	c.Single = 17 - c.BigCount - c.Bomb*4 - c.Triple*3 - c.Pair*2
}

func (c *BaseCombine) IsValid() bool {
	return c.BigCount+c.Bomb*4+c.Triple*3+c.Pair*2 <= 17
}
