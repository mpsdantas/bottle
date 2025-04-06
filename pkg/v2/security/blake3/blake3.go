package blake3

import (
	"encoding/hex"

	"github.com/zeebo/blake3"
)

type Blake3 struct {
	key string
}

func New(key string) *Blake3 {
	return &Blake3{
		key: key,
	}
}

func (b *Blake3) Hash(value string) (string, error) {
	h, err := blake3.NewKeyed([]byte(b.key))
	if err != nil {
		return "", err
	}

	if _, err := h.Write([]byte(value)); err != nil {
		return "", err
	}

	encoded := hex.EncodeToString(h.Sum(nil))

	return encoded, nil
}
