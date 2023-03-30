package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{
			name: "simple random string",
			size: 32,
		},
	}

	for _, test := range tests {
		value := String(test.size)

		a := assert.New(t)

		a.Equal(test.size, len(value))
	}
}

func TestBase64(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{
			name: "simple base64",
			size: 32,
		},
	}

	for _, test := range tests {
		value := Base64(test.size)

		a := assert.New(t)

		a.Equal(test.size, len(value))
	}
}

func TestBase62(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{
			name: "simple base62",
			size: 32,
		},
	}

	for _, test := range tests {
		value := Base62(test.size)

		a := assert.New(t)

		a.Equal(test.size, len(value))
	}
}

func TestHex(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{
			name: "simple hex",
			size: 32,
		},
	}

	for _, test := range tests {
		value := Hex(test.size)

		a := assert.New(t)

		a.Equal(test.size*2, len(value))
	}
}

func TestBytes(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{
			name: "simple hex",
			size: 32,
		},
	}

	for _, test := range tests {
		value := Bytes(test.size)

		a := assert.New(t)

		a.Equal(test.size, len(value))
	}
}
