package common_util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"math/rand"
)

const randomCharacterSet = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const AES_BLOCK_SIZE = 16

var AES_CBC_IV []byte //= '\0' * AES_BLOCK_SIZE
const GARENA_PADDING_MIN_BLOCK = 3

func init() {
	AES_CBC_IV = make([]byte, AES_BLOCK_SIZE)
	for i := 0; i < AES_BLOCK_SIZE; i++ {
		AES_CBC_IV[i] = '\x00'
	}
}

var CryptUtil cryptUtil

type cryptUtil struct {
}

func (r *cryptUtil) RandomString(length int, allowedChars string) string {
	t := []rune(allowedChars)
	maxIndex := len(t) - 1
	var buf bytes.Buffer
	for i := 0; i < length; i++ {
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

func (r *cryptUtil) GarenaAesEncrypt(data string, key []byte) ([]byte, error) {
	paddingData := r.garenaCbcPadding([]byte(data), AES_BLOCK_SIZE)
	return r.aesCBCEncrypt(paddingData, key)
}

func (r *cryptUtil) GarenaAesDecrypt(data []byte, key []byte) (string, error) {
	dataSize := len(data)
	if dataSize%AES_BLOCK_SIZE != 0 {
		return "", nil
	}
	if dataSize < AES_BLOCK_SIZE*GARENA_PADDING_MIN_BLOCK {
		return "", nil
	}
	paddingData, err := r.aesCbcDecrypt(data, key)
	if err != nil {
		return "", fmt.Errorf("aesCbcDecrypt error:%s", err)
	}
	return string(r.garenaCbcUnpadding(paddingData, AES_BLOCK_SIZE)), nil
}

func randomBytes(length int) []byte {
	var res []byte
	for i := 0; i < length; i++ {
		res = append(res, uint8(rand.Intn(255)))
	}
	return res
}

func (r *cryptUtil) garenaCbcPadding(rawData []byte, blockSize int) []byte {
	data := make([]byte, 0)
	data = append(data, randomBytes(blockSize)...)
	data = append(data, rawData...)
	data = r.pkcs7Padding(data, blockSize)
	checkSum := make([]uint8, blockSize)
	i := 0
	for _, ch := range data {
		checkSum[i] ^= ch
		i += 1
		if i >= blockSize {
			i = 0
		}
	}
	res := make([]byte, 0)
	res = append(res, data...)
	res = append(res, checkSum...)
	return res

}
func (r *cryptUtil) garenaCbcUnpadding(data []byte, blockSize int) []byte {
	dataSize := len(data)
	if dataSize%AES_BLOCK_SIZE != 0 {
		return nil
	}
	if dataSize < AES_BLOCK_SIZE*GARENA_PADDING_MIN_BLOCK {
		return nil
	}
	checkSum := make([]uint8, blockSize)
	i := 0
	for _, ch := range data {
		checkSum[i] ^= ch
		i += 1
		if i >= blockSize {
			i = 0
		}
	}
	for _, v := range checkSum {
		if v != 0 {
			return nil
		}
	}
	return r.pkcs7UnPadding(data[blockSize:len(data)-blockSize], blockSize)
}

func (r *cryptUtil) pkcs7Padding(data []byte, blockSize int) []byte {
	pad := blockSize - (len(data) % blockSize)
	padding := make([]byte, pad)
	for i := 0; i < pad; i++ {
		padding[i] = uint8(pad)
	}
	return append(data, padding...)
}

func (r *cryptUtil) pkcs7UnPadding(data []byte, blockSize int) []byte {
	size := len(data)
	if size < blockSize || size%blockSize != 0 {
		return nil
	}
	pad := data[len(data)-1]
	if pad <= 0 || int(pad) > blockSize {
		return nil
	}
	for i := 2; i < int(pad)+1; i++ {
		if data[len(data)-i] != pad {
			return nil
		}
	}
	return data[0 : len(data)-int(pad)]
}

func (r *cryptUtil) aesCbcDecrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, AES_CBC_IV)
	decrypted := make([]byte, len(data))
	mode.CryptBlocks(decrypted, data)
	return decrypted, nil
}

//aes加密，填充秘钥key的16位，24,32分别对应AES-128, AES-192, or AES-256.
func (r *cryptUtil) aesCBCEncrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCEncrypter(block, AES_CBC_IV)
	decrypted := make([]byte, len(data))
	mode.CryptBlocks(decrypted, data)
	return decrypted, nil
}

func (r *cryptUtil) Sha512(plainText string) string {
	hash := sha512.New()
	hash.Write([]byte(plainText))
	return hex.EncodeToString(hash.Sum(nil))
}

func (r *cryptUtil) MD5AndHexDigest(plainText string) string {
	hash := md5.New()
	hash.Write([]byte(plainText))
	return hex.EncodeToString(hash.Sum(nil))
}
