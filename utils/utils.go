package utils

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
