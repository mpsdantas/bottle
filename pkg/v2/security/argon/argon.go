package argon

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type Argon struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

type decodeHashOptions struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
	hash        []byte
	salt        []byte
}

func New(opts ...Option) *Argon {
	defaults := &options{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	for _, opt := range opts {
		opt(defaults)
	}

	return &Argon{
		memory:      defaults.memory,
		iterations:  defaults.iterations,
		parallelism: defaults.parallelism,
		saltLength:  defaults.saltLength,
		keyLength:   defaults.keyLength,
	}
}

func (a *Argon) Verify(password string, hash string) (bool, error) {
	p, err := decodeHash(hash)
	if err != nil {
		return false, err
	}

	otherHash := argon2.IDKey([]byte(password), p.salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	if subtle.ConstantTimeCompare(p.hash, otherHash) == 1 {
		return true, nil
	}

	return false, nil
}

func decodeHash(encodedHash string) (*decodeHashOptions, error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, ErrInvalidHash
	}

	var version int
	if _, err := fmt.Sscanf(vals[2], "v=%d", &version); err != nil {
		return nil, err
	}

	if version != argon2.Version {
		return nil, ErrIncompatibleVersion
	}

	p := &decodeHashOptions{}
	if _, err := fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism); err != nil {
		return nil, err
	}

	salt, err := base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, err
	}

	p.saltLength = uint32(len(salt))

	hash, err := base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, err
	}

	p.keyLength = uint32(len(hash))
	p.salt = salt
	p.hash = hash

	return p, nil
}

func (a *Argon) Hash(password string) (string, error) {
	salt, err := generateRandomBytes(a.saltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, a.iterations, a.memory, a.parallelism, a.keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, a.memory, a.iterations, a.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
