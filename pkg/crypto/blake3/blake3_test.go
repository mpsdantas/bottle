package blake3

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlake_Hash(t *testing.T) {
	type args struct {
		value string
		hash  string
	}

	type want struct {
		value string
		err   error
	}

	tests := []struct {
		name string
		args func() *args
		want func(args *args) *want
	}{
		{
			name: "invalid key size",
			args: func() *args {
				return &args{
					value: strings.Repeat("a", 20),
					hash:  "test-123",
				}
			},
			want: func(args *args) *want {
				return &want{
					err: fmt.Errorf("invalid key size"),
				}
			},
		},
		{
			name: "success",
			args: func() *args {
				return &args{
					value: strings.Repeat("a", 32),
					hash:  "test-123",
				}
			},
			want: func(args *args) *want {
				return &want{
					value: "eedae427030802c2b4525d6eae96efadc829be75d0eb8fd646e1105a134d72ea",
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.args()
			want := test.want(args)

			bl := New(args.value)

			value, err := bl.Hash(args.hash)

			a := assert.New(t)

			a.Equal(want.err, err)
			a.Equal(want.value, value)
		})
	}
}
