package utils

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
)

func Image2Base64(img image.Image) string {
	var b bytes.Buffer
	png.Encode(&b, img)
	return base64.StdEncoding.EncodeToString(b.Bytes())
}