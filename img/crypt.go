package img

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
)

const randomCharacterSet = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var CryptUtil cryptUtil

type cryptUtil struct {
}

func (r *cryptUtil) RandomString(length int, allowedChars string) string {
	t := []rune(allowedChars)
	maxIndex := len(t) - 1
	var buf bytes.Buffer
	for i :=0; i<length; i++ {
		buf.WriteRune(t[rand.Intn(maxIndex)])
	}
	return buf.String()
}

func (r *cryptUtil) RandomStringUseDefaultAllowedChars(length int) string {
	return r.RandomString(length, randomCharacterSet)
}

func (r *cryptUtil) HmacSha256(plainText string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(plainText))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func (r *cryptUtil) Sha256(plainText string) string {
	hash := sha256.New()
	hash.Write([]byte(plainText))
	return hex.EncodeToString(hash.Sum(nil))
}