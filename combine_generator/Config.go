package combine_generator

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type CombineConfig struct {
	ID       uint32 `json:"ID"`
	Bomb     uint32 `json:"Bomb"`
	KingBomb uint32 `json:"KingBomb"`
	Triple   uint32 `json:"Triple"`
	Pair     uint32 `json:"Pair"`
	Single   uint32 `json:"Single"`
	EventID  string `json:"EventID"`
	Weight   uint32 `json:"Weight"`
}

type CombineConfigs struct {
	ConfigList     []CombineConfig          // 所有配置列表
	ConfigIDMap    map[uint32]CombineConfig //
	TripleIDList   []uint32                 // 带三张的配置ID列表
	PairIDList     []uint32                 // 带对子的配置ID列表
	SingleIDList   []uint32                 // 带单张的配置ID列表
	BombIDList     []uint32                 // 带炸弹的配置ID列表
	KingBombIDList []uint32                 // 带王炸的配置ID列表
	AllIDList      []uint32                 // 所有配置ID列表
}

// ParseConfig 配置解析
func ParseConfig(path string, conf interface{}) bool {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(conf)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func (cc *CombineConfigs) GetOneConfig(configList []uint32) CombineConfig {
	idx := rand.Intn(len(configList))
	id := configList[idx]
	config := cc.ConfigIDMap[id]
	return config
}

func (cc *CombineConfigs) GetConfigsByCombineType(combineType uint32) []uint32 {
	var configIDList []uint32
	switch combineType {
	case CombineTypeTriple:
		configIDList = make([]uint32, len(cc.TripleIDList))
		copy(configIDList, cc.TripleIDList)
	case CombineTypePair:
		configIDList = make([]uint32, len(cc.PairIDList))
		copy(configIDList, cc.PairIDList)
	case CombineTypeSingle:
		configIDList = make([]uint32, len(cc.SingleIDList))
		copy(configIDList, cc.SingleIDList)
	case CombineTypeBomb:
		configIDList = make([]uint32, len(cc.BombIDList))
		copy(configIDList, cc.BombIDList)
	case CombineTypeKingBomb:
		configIDList = make([]uint32, len(cc.KingBombIDList))
		copy(configIDList, cc.KingBombIDList)
	default:
		break
	}
	return configIDList
}

func NewCombineConfigs() *CombineConfigs {
	return &CombineConfigs{
		ConfigList:     make([]CombineConfig, 0),
		ConfigIDMap:    make(map[uint32]CombineConfig),
		TripleIDList:   make([]uint32, 0),
		PairIDList:     make([]uint32, 0),
		SingleIDList:   make([]uint32, 0),
		BombIDList:     make([]uint32, 0),
		KingBombIDList: make([]uint32, 0),
		AllIDList:      make([]uint32, 0),
	}
}

func (cc *CombineConfigs) Init(path string) {
	ParseConfig(path /*"./res/config/config_BaseCombine.json"*/, &cc.ConfigList)
	for _, config := range cc.ConfigList {
		cc.ConfigIDMap[config.ID] = config
		if config.Triple > 0 {
			cc.TripleIDList = append(cc.TripleIDList, config.ID)
		}
		if config.Pair > 0 {
			cc.PairIDList = append(cc.PairIDList, config.ID)
		}
		if config.Single > 0 {
			cc.SingleIDList = append(cc.SingleIDList, config.ID)
		}
		if config.Bomb > 0 {
			cc.BombIDList = append(cc.BombIDList, config.ID)
		}
		if config.KingBomb > 0 {
			cc.KingBombIDList = append(cc.KingBombIDList, config.ID)
		}
		cc.AllIDList = append(cc.AllIDList, config.ID)
	}
}

func (cc *CombineConfigs) GetIDListByCT(combineType uint32) []uint32 {
	var configIDList []uint32
	switch combineType {
	case CombineTypeTriple:
		configIDList = make([]uint32, len(cc.TripleIDList))
		copy(configIDList, cc.TripleIDList)
	case CombineTypePair:
		configIDList = make([]uint32, len(cc.PairIDList))
		copy(configIDList, cc.PairIDList)
	case CombineTypeSingle:
		configIDList = make([]uint32, len(cc.SingleIDList))
		copy(configIDList, cc.SingleIDList)
	case CombineTypeBomb:
		configIDList = make([]uint32, len(cc.BombIDList))
		copy(configIDList, cc.BombIDList)
	case CombineTypeKingBomb:
		configIDList = make([]uint32, len(cc.KingBombIDList))
		copy(configIDList, cc.KingBombIDList)
	}
	return configIDList
}

func (cc *CombineConfigs) GetConfigByID(ID uint32) *CombineConfig {
	if conf, ok := cc.ConfigIDMap[ID]; ok {
		return &conf
	}
	return nil
}

func RandomOneID(IDList []uint32) uint32 {
	idx := rand.Intn(len(IDList))
	return IDList[idx]
}
