package scrypt

import (
	"bytes"
	"encoding/hex"
	"io"

	"golang.org/x/crypto/scrypt"
)

func Encode(pass, salt io.Reader) string {
	passBuf := new(bytes.Buffer)
	passBuf.ReadFrom(pass)
	passBytes := passBuf.Bytes()

	saltBuf := new(bytes.Buffer)
	saltBuf.ReadFrom(salt)
	saltBytes := saltBuf.Bytes()

	data, _ := scrypt.Key(passBytes, saltBytes, 16384, 8, 1, 32)
	return hex.EncodeToString(data)
}
