package random

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
)

var (
	DefaultLetters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	Base64Letters  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/"
	Base62Letters  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func String(size int, letters ...string) string {
	letterRunes := DefaultLetters

	if len(letters) > 0 {
		letterRunes = []rune(letters[0])
	}

	var bb bytes.Buffer
	bb.Grow(size)
	l := uint32(len(letterRunes))
	// on each loop, generate one random rune and append to output
	for i := 0; i < size; i++ {
		bb.WriteRune(letterRunes[binary.BigEndian.Uint32(Bytes(4))%l])
	}

	return bb.String()
}

func Base64(n int) string {
	return String(n, Base64Letters)
}

func Base62(s int) string {
	return String(s, Base62Letters)
}

func Hex(size int) string {
	return hex.EncodeToString(Bytes(size))
}

func Bytes(n int) []byte {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return b
}
