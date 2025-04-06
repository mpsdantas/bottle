package blake3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlake3HashWithSuccess(t *testing.T) {
	key := "12345678901234567890123456789012"
	b := New(key)

	hashed, err := b.Hash("hello world")

	a := assert.New(t)
	a.NoError(err)
	a.NotEmpty(hashed)
}

func TestBlake3HashWithInvalidKey(t *testing.T) {
	key := "short"
	b := New(key)

	hashed, err := b.Hash("hello world")

	a := assert.New(t)
	a.Error(err)
	a.Empty(hashed)
}

func TestBlake3HashWithEmptyValue(t *testing.T) {
	key := "12345678901234567890123456789012"
	b := New(key)

	hashed, err := b.Hash("")

	a := assert.New(t)
	a.NoError(err)
	a.NotEmpty(hashed)
}

func TestBlake3HashIsDeterministic(t *testing.T) {
	key := "12345678901234567890123456789012"
	b := New(key)

	h1, err1 := b.Hash("some value")
	h2, err2 := b.Hash("some value")

	a := assert.New(t)
	a.NoError(err1)
	a.NoError(err2)
	a.Equal(h1, h2)
}

func TestBlake3HashWithDifferentKeys(t *testing.T) {
	k1 := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	k2 := "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"

	b1 := New(k1)
	b2 := New(k2)

	h1, err1 := b1.Hash("same data")
	h2, err2 := b2.Hash("same data")

	a := assert.New(t)
	a.NoError(err1)
	a.NoError(err2)
	a.NotEqual(h1, h2)
}

func TestBlake3HashWithDifferentInputs(t *testing.T) {
	key := "12345678901234567890123456789012"
	b := New(key)

	h1, err1 := b.Hash("data 1")
	h2, err2 := b.Hash("data 2")

	a := assert.New(t)
	a.NoError(err1)
	a.NoError(err2)
	a.NotEqual(h1, h2)
}
