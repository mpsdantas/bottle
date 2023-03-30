package argon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgon_Hash(t *testing.T) {
	type args struct {
		password string
	}

	type want struct {
		hash string
		size int
		err  error
	}

	tests := []struct {
		name string
		args *args
		want *want
	}{
		{
			name: "simple hash password",
			args: &args{
				password: "MinhaSenhaLegal",
			},
			want: &want{
				hash: "$argon2id$v=19$m=65536,t=3,p=2",
				size: 97,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args, want := test.args, test.want

			arg := New()

			hash, err := arg.Hash(args.password)

			a := assert.New(t)

			if want.err != nil {
				a.Equal(want.err, err)
			}

			if want.err == nil {
				a.Contains(hash, want.hash)
				a.Equal(want.size, len(hash))
			}
		})
	}
}

func TestArgon_Verify(t *testing.T) {
	type args struct {
		password string
		hash     string
	}

	type want struct {
		ok  bool
		err error
	}

	tests := []struct {
		name string
		args func() *args
		want *want
	}{
		{
			name: "simple verify password",
			args: func() *args {
				password := "MinhaSenhaLegal"
				hash, _ := New().Hash(password)

				return &args{
					password: password,
					hash:     hash,
				}
			},
			want: &want{
				ok:  true,
				err: nil,
			},
		},
		{
			name: "invalid password",
			args: func() *args {
				password := "MinhaSenhaLegal"
				hash, _ := New().Hash(password)

				return &args{
					password: password + "123",
					hash:     hash,
				}
			},
			want: &want{
				ok:  false,
				err: nil,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args, want := test.args(), test.want

			arg := New()

			ok, err := arg.Verify(args.password, args.hash)

			a := assert.New(t)

			if want.err != nil {
				a.Equal(want.err, err)
			}

			if want.err == nil {
				a.Equal(want.ok, ok)
			}
		})
	}
}
