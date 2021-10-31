package enc

import (
	"encoding/base64"
	"strings"
)

/*
	this.Wm = function (e) {
		return s(5) + n.Zk(e) + s(3);
	}
*/
func Wm(jsonString []byte) string {
	var (
		sb strings.Builder
	)
	// Длина закодированого json в base64
	n := base64.StdEncoding.EncodedLen(len(jsonString))

	sb.Grow(n + 5 + 3)

	jsonB64 := make([]byte, n)
	base64.StdEncoding.Encode(jsonB64, jsonString)

	// метод zk() оказался методом base64

	sb.Write(randomString(5))
	sb.Write(jsonB64)
	sb.Write(randomString(3))

	return sb.String()
}