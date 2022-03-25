package services

import (
	"strings"

	qrcode "github.com/skip2/go-qrcode"
)

func GenerateCodes(salt string) (string, []byte) {
	var png []byte
	shortId := strings.ToUpper(salt[10:19])

	png, err := qrcode.Encode(shortId, qrcode.Medium, 256)

	if err != nil {
		panic(err)
	}

	return shortId, png
}
