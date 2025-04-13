package argon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgonHashAndVerifyWithSuccess(t *testing.T) {
	password := "supersecret"
	arg := New()

	hash, err := arg.Hash(password)

	a := assert.New(t)
	a.NoError(err)
	a.NotEmpty(hash)

	ok, err := arg.Verify(password, hash)
	a.NoError(err)
	a.True(ok)
}

func TestArgonVerifyWithWrongPassword(t *testing.T) {
	password := "correct-password"
	wrongPassword := "wrong-password"

	arg := New()

	hash, err := arg.Hash(password)

	a := assert.New(t)
	a.NoError(err)
	a.NotEmpty(hash)

	ok, err := arg.Verify(wrongPassword, hash)
	a.NoError(err)
	a.False(ok)
}

func TestArgonVerifyWithInvalidHashFormat(t *testing.T) {
	arg := New()
	invalidHash := "invalid$hash$format"

	ok, err := arg.Verify("any-password", invalidHash)

	a := assert.New(t)
	a.Error(err)
	a.False(ok)
	a.ErrorIs(err, ErrInvalidHash)
}

func TestArgonVerifyWithInvalidVersion(t *testing.T) {
	invalidHash := "$argon2id$v=999$m=65536,t=3,p=2$" +
		"c2FsdGhlcmU$" +
		"aGFzaGVycmVy"

	arg := New()
	ok, err := arg.Verify("any-password", invalidHash)

	a := assert.New(t)
	a.Error(err)
	a.False(ok)
	a.ErrorIs(err, ErrIncompatibleVersion)
}

func TestArgonVerifyWithCorruptedBase64(t *testing.T) {
	invalidHash := "$argon2id$v=19$m=65536,t=3,p=2$" +
		"invalid-base64-salt$" +
		"invalid-base64-hash"

	arg := New()
	ok, err := arg.Verify("any-password", invalidHash)

	a := assert.New(t)
	a.Error(err)
	a.False(ok)
}

func TestArgonWithCustomOptions(t *testing.T) {
	arg := New(
		WithMemory(128*1024),
		WithIterations(4),
		WithParallelism(4),
		WithSaltLength(32),
		WithKeyLength(64),
	)

	hash, err := arg.Hash("password")
	a := assert.New(t)
	a.NoError(err)
	a.NotEmpty(hash)

	ok, err := arg.Verify("password", hash)
	a.NoError(err)
	a.True(ok)
}
