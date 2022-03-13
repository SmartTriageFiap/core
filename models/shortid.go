package models

type ShortId struct {
	QrCode  []byte `json:qrcode`
	ShortId string `json:short_id`
}
