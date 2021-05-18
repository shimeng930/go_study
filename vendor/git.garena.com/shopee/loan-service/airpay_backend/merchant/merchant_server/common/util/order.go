package common_util

import "time"

func IsSameDayOrder(createTime1 int64, createTime2 int64) bool {
	cTime1 := time.Unix(createTime1, 0)
	cTime2 := time.Unix(createTime2, 0)
	curDay1 := time.Date(cTime1.Year(), cTime1.Month(), cTime1.Day(), 0, 0, 0, 0, cTime1.Location())
	curDay2 := time.Date(cTime2.Year(), cTime2.Month(), cTime2.Day(), 0, 0, 0, 0, cTime2.Location())
	if curDay1.Unix() == curDay2.Unix() {
		return true
	}
	return false
}
