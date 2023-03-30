package blake3

import (
	"encoding/hex"

	"github.com/zeebo/blake3"
)

//go:generate mockgen -source=./blake3.go -package=blake3 -destination=./blake3_mock.go
type Hasher interface {
	Hash(value string) (string, error)
}

type blake struct {
	key string
}

func New(key string) Hasher {
	return &blake{
		key: key,
	}
}

func (b *blake) Hash(value string) (string, error) {
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
