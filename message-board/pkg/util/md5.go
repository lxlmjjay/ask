package util

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(time.Now().String()+value))
	return hex.EncodeToString(m.Sum(nil))
}
