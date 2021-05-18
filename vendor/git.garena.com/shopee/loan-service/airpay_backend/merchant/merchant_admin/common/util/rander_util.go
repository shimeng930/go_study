package util

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	numericSample     = "0123456789"                 // 数字
	lowerLetterSample = "qwertyuiopasdfghjklzxcvbnm" // 小写字母
	upperLetterSample = "QWERTYUIOPASDFGHJKLZXCVBNM" // 大写字母
)

var Rander rander

type rander struct {
}

// 生成随机字符串（字母数字）
func (r *rander) CreateRandomString(randLen int) string {
	return r._genRandom(numericSample+lowerLetterSample+upperLetterSample, randLen)
}

// 生成随机字符串（数字）
func (r *rander) CreateRandomNumber(randLen int) string {
	return r._genRandom(numericSample, randLen)
}

func (r *rander) _genRandom(sampleStr string, randLen int) string {
	sample := []byte(sampleStr)
	if len(sample) == 0 || randLen == 0 {
		return ""
	}

	var container string
	length := bytes.NewReader(sample).Len()

	for i := 1; i <= randLen; i++ {
		random, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {
		}
		container += fmt.Sprintf("%c", sample[random.Int64()])
	}
	return container
}
