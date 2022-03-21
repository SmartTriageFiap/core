package services

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
)

func Salt(doc string) string {
	h := sha256.New()
	h.Write([]byte(doc + os.Getenv("GO_CRIPYT")))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}
