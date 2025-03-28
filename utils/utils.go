package utils

import "strconv"

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

func ToHex(arr []int) []string {
	hexSlice := make([]string, len(arr))
	for i, num := range arr {
		// 使用strconv包将整数转换为16进制字符串，不带前缀0x
		hexStr := strconv.FormatInt(int64(num), 16)
		hexSlice[i] = hexStr
	}
	return hexSlice
}
