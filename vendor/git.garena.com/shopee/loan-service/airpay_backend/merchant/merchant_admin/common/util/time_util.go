package util

import (
	"time"
)

func TimeNow() uint64 {
	return uint64(time.Now().Unix())
}

func TimeNowNano() uint64 {
	return uint64(time.Now().UnixNano())
}

func TimeNowMs() uint64 {
	return uint64(time.Now().UnixNano() / 1000000)
}

func Time2RoundDay(t uint64) uint64 {
	tt := time.Unix(int64(t), 0)
	//y, m, d := tt.Date()
	//return uint64(time.Date(y, m, d, 0, 0, 0, 0, time.UTC).Unix())
	//return uint64(tt.Round(24*time.Hour).Unix())
	h, m, s := tt.Clock()
	return uint64(t - uint64(h*3600+m*60+s))
}

func Time2RoundHour(t uint64) uint64 {
	tt := time.Unix(int64(t), 0)
	_, m, s := tt.Clock()
	return uint64(t - uint64(m*60+s))
}
