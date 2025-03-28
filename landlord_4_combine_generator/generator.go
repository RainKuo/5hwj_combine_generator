package landlord_4_combine_generator

import (
	"CombineGenerator/proto/out/landlord_4p_combines"
	"CombineGenerator/utils"
	"fmt"
	"sort"
)

const (
	SI_XING_ZHA = iota
	ER_WANG_ZHA
	WU_XING_ZHA
	LIU_XING_ZHA
	SAN_WANG_ZHA
	QI_XING_ZHA
	SI_XING_ER_LIAN_ZHA
	BA_XING_ZHA
	SI_WANG_ZHA
	WU_XING_ER_LIAN_ZHA
	WU_WANG_ZHA
	SI_XING_SAN_LIAN_ZHA
	LIU_XING_ER_LIAN_ZHA
	LIU_WANG_ZHA
	QI_XING_ER_LIAN_ZHA
	QI_WANG_ZHA
	WU_XING_SAN_LIAN_ZHA
	SI_XING_SI_LIAN_ZHA
	BA_XING_ER_LIAN_ZHA
	BA_WANG_ZHA
	LIU_XING_SAN_LIAN_ZHA
	JIU_WANG_ZHA
	SI_XING_WU_LIAN_ZHA
	WU_XING_SI_LIAN_ZHA
	SHI_WANG_ZHA
	QI_XING_SAN_LIAN_ZHA
	SHI_YI_WANG_ZHA
	SI_XING_LIU_LIAN_ZHA
	LIU_XING_SI_LIAN_ZHA
	BA_XING_SAN_LIAN_ZHA
	SHI_ER_WANG_ZHA
	WU_XING_WU_LIAN_ZHA
	SI_XING_QI_LIAN_ZHA
	QI_XING_SI_LIAN_ZHA
	BOMB_COMBINE_MAX
)

// 普通牌型枚举
const (
	NORMAL_BA_LIAN_FEI_JI = BOMB_COMBINE_MAX + iota
	NORMAL_QI_LIAN_FEI_JI
	NORMAL_LIU_LIAN_FEI_JI
	NORMAL_WU_LIAN_FEI_JI
	NORMAL_SI_LIAN_FEI_JI
	NORMAL_SAN_LIAN_FEI_JI
	NORMAL_ER_LIAN_FEI_JI
	NORMAL_SHI_LIAN_DUI
	NORMAL_JIU_LIAN_DUI
	NORMAL_BA_LIAN_DUI
	NORMAL_QI_LIAN_DUI
	NORMAL_LIU_LIAN_DUI
	NORMAL_WU_LIAN_DUI
	NORMAL_SI_LIAN_DUI
	NORMAL_SAN_LIAN_DUI
	NORMAL_SHI_LIAN_SHUN
	NORMAL_JIU_LIAN_SHUN
	NORMAL_BA_LIAN_SHUN
	NORMAL_QI_LIAN_SHUN
	NORMAL_LIU_LIAN_SHUN
	NORMAL_WU_LIAN_SHUN
	NORMAL_TRIPLE
	NORMAL_PAIR
	NORMAL_SINGLE
)

const MAX_CARD_COUNT_LENGTH = 17 // 最大牌值列表长度
const MAX_HAND_CARD_COUNT = 25   // 最大手牌数量

var StarTable map[int][2]int               // 连炸对应的星数表[几连炸][0](最小星数)[1](最大星数)
var CombineMutiples [BOMB_COMBINE_MAX]int  // 牌型倍数
var PassibleBombCardCount []int            // 可能的炸弹牌型张数
var NearBombInfoMap map[int][]NearBombInfo // 张数与炸弹类型数据的映射
var PassibleNormalCombineCount []int       // 可能的普通牌型张数
var NearNormalInfoMap map[int][]NearNormalInfo

type PlayerHandCardInfo struct {
	cards     []int // 手牌
	turnCount int   // 手数(几手能出完)
	multiple  int   // 倍数
	score     int   // 得分
	id        int
}

func (info *PlayerHandCardInfo) CalcScore() {
	info.score = 0
	if info.turnCount == 1 {
		info.turnCount = 2
	}
	info.score = info.multiple / (info.turnCount - 1)
}

// NearBombInfo 最接近的炸弹数据
type NearBombInfo struct {
	starCount   int // 星数
	seqCount    int // 连数
	combineType int // 炸弹类型
	cardCount   int // 牌型张数
}

type NearNormalInfo struct {
	starCount   int // 星数
	seqCount    int // 连数
	combineType int // 普通类型
	cardCount   int // 牌型张数
}

type CombineGenerator4 struct {
	PlayerHandCardsInfo []PlayerHandCardInfo
	firstHandChoices    [10]int                    // 第一手牌
	cardCounts          [MAX_CARD_COUNT_LENGTH]int //牌数 [val]count
}

func Init() {
	StarTable = make(map[int][2]int)
	StarTable[0] = [2]int{6, 12}
	StarTable[1] = [2]int{4, 8}
	StarTable[2] = [2]int{6, 8}
	StarTable[3] = [2]int{5, 8}
	StarTable[4] = [2]int{4, 6}
	StarTable[5] = [2]int{4, 4}
	StarTable[6] = [2]int{4, 4}
	CombineMutiples = [BOMB_COMBINE_MAX]int{2, 2, 2, 4, 4, 4, 4, 4, 4, 8, 8, 8, 8, 8, 16, 16, 16, 16,
		16, 16, 32, 32, 32, 32, 32, 64, 64, 64, 64, 64, 64, 128, 128, 128}
	PassibleBombCardCount = []int{25, 24, 21, 20, 18, 16, 15, 14, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	NearBombInfoMap = make(map[int][]NearBombInfo)
	NearBombInfoMap[25] = []NearBombInfo{{5, 5, WU_XING_WU_LIAN_ZHA, 25}}
	NearBombInfoMap[24] = []NearBombInfo{{8, 3, BA_XING_SAN_LIAN_ZHA, 24}, {4, 6, SI_XING_LIU_LIAN_ZHA, 24}, {6, 4, LIU_XING_SI_LIAN_ZHA, 24}}
	NearBombInfoMap[21] = []NearBombInfo{{7, 3, QI_XING_SAN_LIAN_ZHA, 21}}
	NearBombInfoMap[20] = []NearBombInfo{{5, 4, WU_XING_SI_LIAN_ZHA, 20}, {4, 5, SI_XING_WU_LIAN_ZHA, 20}}
	NearBombInfoMap[18] = []NearBombInfo{{6, 3, LIU_XING_SAN_LIAN_ZHA, 18}}
	NearBombInfoMap[16] = []NearBombInfo{{4, 4, SI_XING_SI_LIAN_ZHA, 16}, {8, 2, BA_XING_ER_LIAN_ZHA, 16}}
	NearBombInfoMap[15] = []NearBombInfo{{5, 3, WU_XING_SAN_LIAN_ZHA, 15}}
	NearBombInfoMap[14] = []NearBombInfo{{7, 2, QI_XING_ER_LIAN_ZHA, 14}}
	NearBombInfoMap[12] = []NearBombInfo{{4, 3, SI_XING_SAN_LIAN_ZHA, 12}, {6, 2, LIU_XING_ER_LIAN_ZHA, 12}, {12, 0, SHI_ER_WANG_ZHA, 12}}
	NearBombInfoMap[11] = []NearBombInfo{{11, 0, SHI_YI_WANG_ZHA, 11}}
	NearBombInfoMap[10] = []NearBombInfo{{5, 2, WU_XING_ER_LIAN_ZHA, 10}, {10, 0, SHI_WANG_ZHA, 10}}
	NearBombInfoMap[9] = []NearBombInfo{{9, 0, JIU_WANG_ZHA, 9}}
	NearBombInfoMap[8] = []NearBombInfo{{4, 2, SI_XING_ER_LIAN_ZHA, 8}, {8, 1, BA_XING_ZHA, 8}, {8, 0, BA_WANG_ZHA, 8}}
	NearBombInfoMap[7] = []NearBombInfo{{7, 0, QI_WANG_ZHA, 7}, {7, 1, QI_XING_ZHA, 7}}
	NearBombInfoMap[6] = []NearBombInfo{{6, 0, LIU_WANG_ZHA, 6}, {6, 1, LIU_XING_ZHA, 6}}
	NearBombInfoMap[5] = []NearBombInfo{{5, 0, WU_WANG_ZHA, 5}, {5, 1, WU_XING_ZHA, 5}}
	NearBombInfoMap[4] = []NearBombInfo{{4, 0, SI_WANG_ZHA, 4}, {4, 1, SI_XING_ZHA, 4}}
	NearBombInfoMap[3] = []NearBombInfo{{3, 0, SAN_WANG_ZHA, 3}}
	NearBombInfoMap[2] = []NearBombInfo{{2, 0, ER_WANG_ZHA, 2}}

	PassibleNormalCombineCount = []int{18, 16, 15, 14, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	NearNormalInfoMap = make(map[int][]NearNormalInfo)
	//NearNormalInfoMap[24] = []NearNormalInfo{{3, 8, NORMAL_BA_LIAN_FEI_JI, 24}}
	//NearNormalInfoMap[21] = []NearNormalInfo{{3, 7, NORMAL_QI_LIAN_FEI_JI, 21}}
	//NearNormalInfoMap[20] = []NearNormalInfo{{3, 6, NORMAL_LIU_LIAN_FEI_JI, 18}, {2, 10, NORMAL_SHI_LIAN_DUI, 20}}
	NearNormalInfoMap[18] = []NearNormalInfo{{3, 6, NORMAL_LIU_LIAN_FEI_JI, 18}, {2, 9, NORMAL_JIU_LIAN_DUI, 18}}
	NearNormalInfoMap[16] = []NearNormalInfo{{3, 5, NORMAL_WU_LIAN_FEI_JI, 15}, {2, 8, NORMAL_BA_LIAN_DUI, 16}}
	NearNormalInfoMap[15] = []NearNormalInfo{{3, 5, NORMAL_WU_LIAN_FEI_JI, 15}}
	NearNormalInfoMap[14] = []NearNormalInfo{{3, 4, NORMAL_SI_LIAN_FEI_JI, 12}, {2, 7, NORMAL_QI_LIAN_DUI, 14}}
	NearNormalInfoMap[13] = []NearNormalInfo{{3, 4, NORMAL_SI_LIAN_FEI_JI, 12}}
	NearNormalInfoMap[12] = []NearNormalInfo{{3, 4, NORMAL_SI_LIAN_FEI_JI, 12}, {2, 6, NORMAL_LIU_LIAN_DUI, 12}}
	NearNormalInfoMap[11] = []NearNormalInfo{{3, 3, NORMAL_SAN_LIAN_FEI_JI, 9}}
	NearNormalInfoMap[10] = []NearNormalInfo{{3, 3, NORMAL_SAN_LIAN_FEI_JI, 9}, {2, 5, NORMAL_WU_LIAN_DUI, 10}}
	NearNormalInfoMap[9] = []NearNormalInfo{{3, 3, NORMAL_SAN_LIAN_FEI_JI, 9}, {2, 4, NORMAL_SI_LIAN_DUI, 8}, {1, 8, NORMAL_JIU_LIAN_SHUN, 9}}
	NearNormalInfoMap[8] = []NearNormalInfo{{3, 2, NORMAL_ER_LIAN_FEI_JI, 6}, {2, 4, NORMAL_SI_LIAN_DUI, 8}, {1, 8, NORMAL_BA_LIAN_SHUN, 8}}
	NearNormalInfoMap[7] = []NearNormalInfo{{3, 2, NORMAL_ER_LIAN_FEI_JI, 6}, {2, 4, NORMAL_SAN_LIAN_DUI, 6}, {1, 7, NORMAL_QI_LIAN_SHUN, 7}}
	NearNormalInfoMap[6] = []NearNormalInfo{{3, 2, NORMAL_ER_LIAN_FEI_JI, 6}, {2, 4, NORMAL_SAN_LIAN_DUI, 6}, {1, 6, NORMAL_LIU_LIAN_SHUN, 6}}
	NearNormalInfoMap[5] = []NearNormalInfo{{1, 5, NORMAL_WU_LIAN_SHUN, 5}}
	NearNormalInfoMap[3] = []NearNormalInfo{{3, 1, NORMAL_TRIPLE, 3}}
	NearNormalInfoMap[2] = []NearNormalInfo{{2, 1, NORMAL_PAIR, 2}}
	NearNormalInfoMap[1] = []NearNormalInfo{{1, 1, NORMAL_SINGLE, 1}}
}

func Landlord4CombineGenerate() *CombineGenerator4 {
	cg4 := &CombineGenerator4{}
	cg4.init()
	cg4.selectFirstHandCards()
	loopTimes := 0
	for {
		info := cg4.GetNeedRepublicPlayerInfo()
		if info == nil {
			break
		}
		if loopTimes > 20 {
			return nil
		}

		cg4.RepublishHandCards(info)
		loopTimes++
	}
	for i := 0; i < len(cg4.PlayerHandCardsInfo); i++ {
		info := &cg4.PlayerHandCardsInfo[i]
		sort.Ints(info.cards)
		info.CalcScore()
	}
	return cg4
}

func (cg4 *CombineGenerator4) init() {
	for i := 0; i < 4; i++ {
		cg4.PlayerHandCardsInfo = append(cg4.PlayerHandCardsInfo, PlayerHandCardInfo{
			cards:     make([]int, 0),
			turnCount: 0,
			score:     0,
			multiple:  0,
			id:        i,
		})
	}
	cg4.firstHandChoices = [10]int{
		5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	}
	cg4.cardCounts = [MAX_CARD_COUNT_LENGTH]int{
		0, 0, 0, 0, 0, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 12,
	}
}

func RandStarCount(seqCount int) int {
	ints := StarTable[seqCount]
	return utils.RandInt(ints[0], ints[1])
}

func GetLianzhaMulti(seqCount int, starCount int) int {
	lianzhaType := -1
	if seqCount == 0 {
		switch starCount {
		case 2:
			lianzhaType = ER_WANG_ZHA
			break
		case 3:
			lianzhaType = SAN_WANG_ZHA
			break
		case 4:
			lianzhaType = SI_WANG_ZHA
			break
		case 5:
			lianzhaType = WU_WANG_ZHA
			break
		case 6:
			lianzhaType = LIU_WANG_ZHA
			break
		case 7:
			lianzhaType = QI_WANG_ZHA
			break
		case 8:
			lianzhaType = BA_WANG_ZHA
			break
		case 9:
			lianzhaType = JIU_WANG_ZHA
			break
		case 10:
			lianzhaType = SHI_WANG_ZHA
			break
		case 11:
			lianzhaType = SHI_YI_WANG_ZHA
			break
		case 12:
			lianzhaType = SHI_ER_WANG_ZHA
			break
		}
	} else {
		if seqCount == 1 {
			if starCount == 4 {
				lianzhaType = SI_XING_ZHA
			} else if starCount == 5 {
				lianzhaType = WU_XING_ZHA
			} else if starCount == 6 {
				lianzhaType = LIU_XING_ZHA
			} else if starCount == 7 {
				lianzhaType = QI_XING_ZHA
			} else if starCount == 8 {
				lianzhaType = BA_XING_ZHA
			}
		} else if seqCount == 2 {
			if starCount == 4 {
				lianzhaType = SI_XING_ER_LIAN_ZHA
			} else if starCount == 5 {
				lianzhaType = WU_XING_ER_LIAN_ZHA
			} else if starCount == 6 {
				lianzhaType = LIU_XING_ER_LIAN_ZHA
			} else if starCount == 7 {
				lianzhaType = QI_XING_ER_LIAN_ZHA
			} else if starCount == 8 {
				lianzhaType = BA_XING_ER_LIAN_ZHA
			}
		} else if seqCount == 3 {
			if starCount == 4 {
				lianzhaType = SI_XING_SAN_LIAN_ZHA
			} else if starCount == 5 {
				lianzhaType = WU_XING_SAN_LIAN_ZHA
			} else if starCount == 6 {
				lianzhaType = LIU_XING_SAN_LIAN_ZHA
			} else if starCount == 7 {
				lianzhaType = QI_XING_SAN_LIAN_ZHA
			} else if starCount == 8 {
				lianzhaType = BA_XING_SAN_LIAN_ZHA
			}
		} else if seqCount == 4 {
			if starCount == 4 {
				lianzhaType = SI_XING_SI_LIAN_ZHA
			} else if starCount == 5 {
				lianzhaType = WU_XING_SI_LIAN_ZHA
			} else if starCount == 6 {
				lianzhaType = LIU_XING_SI_LIAN_ZHA
			} else if starCount == 7 {
				lianzhaType = QI_XING_SI_LIAN_ZHA
			}
		} else if seqCount == 5 {
			if starCount == 4 {
				lianzhaType = SI_XING_WU_LIAN_ZHA
			} else if starCount == 5 {
				lianzhaType = WU_XING_WU_LIAN_ZHA
			}
		} else if seqCount == 6 {
			if starCount == 4 {
				lianzhaType = SI_XING_LIU_LIAN_ZHA
			}
		} else if seqCount == 7 {
			if starCount == 4 {
				lianzhaType = SI_XING_QI_LIAN_ZHA
			}
		}
	}
	if lianzhaType == -1 {
		return 0
	}
	return CombineMutiples[lianzhaType]
}

// 插入手牌
func (cg4 *CombineGenerator4) adaptCards(playerId int, cardVal int, count int) {
	for i := 0; i < count; i++ {
		tmpVal := 0
		if cardVal == 14 {
			tmpVal = 1
		} else if cardVal == 15 {
			tmpVal = 2
		} else {
			tmpVal = cardVal
		}
		cg4.PlayerHandCardsInfo[playerId].cards = append(cg4.PlayerHandCardsInfo[playerId].cards, tmpVal)
	}
	cg4.cardCounts[cardVal] -= count
}

func (cg4 *CombineGenerator4) selectFirstHandCards() {
	// 给第0个玩家随机王数
	jokerCount := RandStarCount(0) // 0是王
	cg4.adaptCards(0, 16, jokerCount)
	cg4.PlayerHandCardsInfo[0].multiple += GetLianzhaMulti(0, jokerCount)
	cg4.PlayerHandCardsInfo[0].turnCount++
	// 给第2个玩家随机值和长度
	midLength := utils.RandInt(2, 7)          // [2, 7)
	midIndex := utils.RandInt(2, 8-midLength) // [2, 8 - N)
	// 第二个玩家随机后,将数组截取为三份
	player1Choices := cg4.firstHandChoices[0:midIndex]
	player2Choices := cg4.firstHandChoices[midIndex : midIndex+midLength]
	player3Choices := cg4.firstHandChoices[midIndex+midLength:]
	// 给第1个玩家随机连炸
	length1 := utils.RandInt(2, len(player1Choices))
	index1 := utils.RandInt(0, len(player1Choices)-length1)
	starCount1 := RandStarCount(length1)
	for i := index1; i < index1+length1; i++ {
		val := player1Choices[i]
		cg4.adaptCards(1, val, starCount1)
	}
	cg4.PlayerHandCardsInfo[1].multiple += GetLianzhaMulti(length1, starCount1)
	cg4.PlayerHandCardsInfo[1].turnCount++

	// 给第2个玩家随机连炸
	starCount2 := RandStarCount(midLength)
	for i := 0; i < len(player2Choices); i++ {
		val := player2Choices[i]
		cg4.adaptCards(2, val, starCount2)
	}
	cg4.PlayerHandCardsInfo[2].multiple += GetLianzhaMulti(len(player2Choices), starCount2)
	cg4.PlayerHandCardsInfo[2].turnCount++

	// 给第3个玩家随机连炸
	length3 := utils.RandInt(2, len(player3Choices))
	index3 := utils.RandInt(0, len(player3Choices)-length3)
	starCount3 := RandStarCount(length3)
	for i := index3; i < index3+length3; i++ {
		val := player3Choices[i]
		cg4.adaptCards(3, val, starCount3)
	}
	cg4.PlayerHandCardsInfo[3].multiple += GetLianzhaMulti(length3, starCount3)
	cg4.PlayerHandCardsInfo[3].turnCount++

	fmt.Println("Player1Choices: ", player1Choices)
	fmt.Println("Player2Choices: ", player2Choices)
	fmt.Println("Player3Choices: ", player3Choices)
	fmt.Printf("Length1: %d, Index1: %d", length1, index1)
	fmt.Printf("Length3: %d, Index3: %d", length3, index3)
}

func (cg4 *CombineGenerator4) IsValid() bool {
	valid := true
	for i := 0; i < len(cg4.PlayerHandCardsInfo); i++ {
		if cg4.PlayerHandCardsInfo[i].turnCount >= 5 {
			valid = false
		}
	}
	return valid
}

func (cg4 *CombineGenerator4) GetNeedRepublicPlayerInfo() *PlayerHandCardInfo {
	maxNeedReplenishCount := 0
	needReplenishIndex := 0
	for i := 0; i < len(cg4.PlayerHandCardsInfo); i++ {
		if len(cg4.PlayerHandCardsInfo[i].cards) < MAX_HAND_CARD_COUNT {
			needCount := MAX_HAND_CARD_COUNT - len(cg4.PlayerHandCardsInfo[i].cards)
			if needCount >= maxNeedReplenishCount {
				maxNeedReplenishCount = needCount
				needReplenishIndex = i
			}
		}
	}
	if maxNeedReplenishCount == 0 {
		return nil
	}
	return &cg4.PlayerHandCardsInfo[needReplenishIndex]
}

// RepublishHandCards 插入离差牌数量最近的牌型
func (cg4 *CombineGenerator4) RepublishHandCards(info *PlayerHandCardInfo) bool {
	if cg4.RepublishBombCards(info) {
		return true
	}
	if cg4.RepublicNormalCards(info) {
		return true
	}
	return false
}

func (cg4 *CombineGenerator4) RepublishBombCards(playerInfo *PlayerHandCardInfo) bool {
	cardsAnalyzer := CardsAnalyzer{}
	cardsAnalyzer.Init(cg4.cardCounts)
	needCount := MAX_HAND_CARD_COUNT - len(playerInfo.cards)
	found := false
	for i := 0; i < len(PassibleBombCardCount); i++ {
		count := PassibleBombCardCount[i]
		if count <= needCount {
			bombInfos := NearBombInfoMap[count]
			for j := 0; j < len(bombInfos); j++ {
				info := bombInfos[j]
				if info.combineType == WU_XING_ZHA || info.combineType == LIU_XING_ZHA || info.combineType == QI_XING_ZHA || info.combineType == BA_XING_ZHA {
					beginIndex := cardsAnalyzer.FindRepeatSingle(info.starCount)
					if beginIndex != -1 {
						found = true
						cg4.adaptCards(playerInfo.id, beginIndex, info.cardCount)
						playerInfo.turnCount++
						break
					}
				} else {
					beginIndex := cardsAnalyzer.FindSequence(info.starCount, info.seqCount)
					if beginIndex != -1 {
						found = true
						// 王
						if beginIndex == 0 {
							cg4.adaptCards(playerInfo.id, 16, info.starCount)

						} else {
							for k := beginIndex; k < beginIndex+info.seqCount; k++ {
								cg4.adaptCards(playerInfo.id, k, info.starCount)
							}
						}
						playerInfo.multiple += CombineMutiples[info.combineType]
						playerInfo.turnCount++
						break
					}
				}

			}
		}
		if found {
			break
		}
	}
	return found
}

func (cg4 *CombineGenerator4) RepublicNormalCards(playerInfo *PlayerHandCardInfo) bool {
	cardsAnalyzer := CardsAnalyzer{}
	cardsAnalyzer.Init(cg4.cardCounts)
	needCount := MAX_HAND_CARD_COUNT - len(playerInfo.cards)
	found := false
	for i := 0; i < len(PassibleNormalCombineCount); i++ {
		count := PassibleNormalCombineCount[i]
		if count <= needCount {
			NormalInfos := NearNormalInfoMap[count]
			for j := 0; j < len(NormalInfos); j++ {
				info := NormalInfos[j]
				if count <= 3 {
					beginIndex := cardsAnalyzer.FindRepeatSingle(info.starCount)
					if beginIndex != -1 {
						found = true
						cg4.adaptCards(playerInfo.id, beginIndex, info.starCount)
						playerInfo.turnCount++
						break
					}
				} else {
					beginIndex := cardsAnalyzer.FindSequence(info.starCount, info.seqCount)
					if beginIndex != -1 {
						found = true
						// 王
						if beginIndex == 0 {
							cg4.adaptCards(playerInfo.id, 16, info.starCount)
						} else {
							for k := beginIndex; k < beginIndex+info.seqCount; k++ {
								cg4.adaptCards(playerInfo.id, k, info.starCount)
							}
						}
						playerInfo.turnCount++
						break
					}
				}
			}
		}
		if found {
			break
		}
	}
	return found
}

func (cg4 *CombineGenerator4) FillProto(config *landlord_4p_combines.Landlord_4PCombineConfig) {
	for i := 0; i < len(cg4.PlayerHandCardsInfo); i++ {
		info := cg4.PlayerHandCardsInfo[i]
		for j := 0; j < len(info.cards); j++ {
			if i == 0 {
				config.Player1HandCards = append(config.Player1HandCards, uint32(info.cards[j]))
			} else if i == 1 {
				config.Player2HandCards = append(config.Player2HandCards, uint32(info.cards[j]))
			} else if i == 2 {
				config.Player3HandCards = append(config.Player3HandCards, uint32(info.cards[j]))
			} else if i == 3 {
				config.Player4HandCards = append(config.Player4HandCards, uint32(info.cards[j]))
			}
		}
		config.Scores = append(config.Scores, uint32(info.score))
		config.Turns = append(config.Turns, uint32(info.turnCount))
		config.BombMultis = append(config.BombMultis, uint32(info.multiple))
	}
}
