package utils

import "math/rand"

func RemoveSliceItem(ids []uint32, id uint32) {
	var newIds []uint32
	for idx, _id := range ids {
		if _id == id {
			newIds = append(newIds, ids[:idx]...)
			newIds = append(newIds, ids[idx+1:]...)
			ids = make([]uint32, len(newIds))
			copy(ids, newIds)
			break
		}
	}
}

// RandInt [min, max) 左闭右开区间
func RandInt(min int, max int) int {
	if max-min == 0 {
		return min
	}
	return rand.Intn(max-min) + min
}
