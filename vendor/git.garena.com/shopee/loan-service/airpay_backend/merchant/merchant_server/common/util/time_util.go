package common_util

import (
	"strconv"
	"time"
)

type DateLayout string

const (
	DateLayout_mY DateLayout = "01-2006"
	// %Y-%m-%d %H:%M:%S
	DateLayout_YmdHMS   DateLayout = "2006-01-02 15:04:05"
	DateLayout_yyyyMMdd DateLayout = "20060102"
)

func NowTimestamp() int64 {
	return time.Now().Unix()
}

func GetCurrentDatetimeTimestamp(hour, minute, second int) int64 {
	now := time.Now()
	t := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, second, now.Nanosecond(), now.Location())
	return t.Unix()
}

func RoundTimeString(format DateLayout, str string) (time.Time, error) {
	date, err := time.Parse(string(format), str)
	if err != nil {
		return time.Now(), err
	}
	t := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	return t, nil
}

func Format(time time.Time, layout DateLayout) string {
	return time.Format(string(layout))
}

func FormatDefaultEmptyString(timestamp int64, layout DateLayout) string {
	if timestamp > 0 {
		t := time.Unix(timestamp, 0)
		return t.Format(string(layout))
	}
	return ""
}

// 毫秒 字符串
func NowTimestampMsStr() string {
	return int64ToString(NowTimestampMs())
}

// 毫秒
func NowTimestampMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func int64PToString(i *int64, defaultV string) string {
	if i == nil {
		return defaultV
	}
	return strconv.FormatInt(*i, 10)
}

func int64ToString(i int64) string {
	return int64PToString(&i, "")
}

func AddDaysUint64(t time.Time, days int) uint64 {
	return uint64(t.Add(time.Duration(days) * 24 * time.Hour).Unix())
}

func AddDays(t time.Time, days int) int64 {
	return t.Add(time.Duration(days) * 24 * time.Hour).Unix()
}
