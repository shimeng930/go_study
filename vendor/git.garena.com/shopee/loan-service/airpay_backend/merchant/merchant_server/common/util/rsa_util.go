package common_util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

/**
  私钥签名 ，SHA256，需要私钥格式为为PKCS1，签名后做base64处理
*/
func RsaSignWithSha256Base64(data string, prvKey string) (string, error) {
	sign, err := base64.StdEncoding.DecodeString(prvKey)
	privateKey, err := x509.ParsePKCS1PrivateKey([]byte(sign))
	if err != nil {
		fmt.Println("ParsePKCS1PrivateKey err", err)
		return "", err
	}

	h := sha256.New()
	h.Write([]byte([]byte(data)))
	hash := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		return "", err
	}
	out := base64.StdEncoding.EncodeToString(signature)
	return out, nil
}

/**
  公钥验签，SHA256，需要公钥格式为为PKCS1，signData为base64数据
  err为nil 为验签成功
*/
func RsaVerySignWithSha256Base64(originalData, signData, pubKey string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	public, _ := base64.StdEncoding.DecodeString(pubKey)
	pub, err := x509.ParsePKIXPublicKey(public)
	if err != nil {
		return err
	}
	hash := sha256.New()
	hash.Write([]byte(originalData))
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA256, hash.Sum(nil), sign)
}
