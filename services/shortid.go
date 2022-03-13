package services

import (
	"hmv-rest-api/database"

	qrcode "github.com/skip2/go-qrcode"
)

func GenerateCodes(cpf string) (string, []byte) {
	var png []byte
	shortId := database.GetShortId(cpf)

	png, err := qrcode.Encode(shortId, qrcode.Medium, 256)

	if err != nil {
		panic(err)
	}

	return shortId, png
}
