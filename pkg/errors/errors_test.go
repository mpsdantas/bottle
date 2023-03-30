package errors

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestNotFound(t *testing.T) {
	type args struct {
		msg string
	}

	type want struct {
		err error
	}

	tests := []struct {
		name string
		args func() *args
		want func(args *args) *want
	}{
		{
			name: "simple test not found",
			args: func() *args {
				return &args{
					msg: "user not found",
				}
			},
			want: func(args *args) *want {
				return &want{
					err: NotFound(args.msg),
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.args()
			want := test.want(args)

			err := NotFound(args.msg)

			a := assert.New(t)

			a.Equal(want.err.Error(), err.Error())
		})
	}
}

func TestInternal(t *testing.T) {
	type args struct {
		msg string
	}

	type want struct {
		err error
	}

	tests := []struct {
		name string
		args func() *args
		want func(args *args) *want
	}{
		{
			name: "simple test internal",
			args: func() *args {
				return &args{
					msg: "cannot get user by id",
				}
			},
			want: func(args *args) *want {
				return &want{
					err: Internal(args.msg),
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.args()
			want := test.want(args)

			err := Internal(args.msg)

			a := assert.New(t)

			a.Equal(want.err.Error(), err.Error())
		})
	}
}

func TestValidation(t *testing.T) {
	type args struct {
		msg string
	}

	type want struct {
		err error
	}

	tests := []struct {
		name string
		args func() *args
		want func(args *args) *want
	}{
		{
			name: "simple test validation",
			args: func() *args {
				return &args{
					msg: "invalid email",
				}
			},
			want: func(args *args) *want {
				return &want{
					err: Validation(args.msg),
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.args()
			want := test.want(args)

			err := Validation(args.msg)

			a := assert.New(t)

			a.Equal(want.err.Error(), err.Error())
		})
	}
}

func TestFailureCause(t *testing.T) {
	type args struct {
		cause string
		msg   string
	}

	type want struct {
		err error
	}

	tests := []struct {
		name string
		args func() *args
		want func(args *args) *want
	}{
		{
			name: "simple test failure acuse",
			args: func() *args {
				return &args{
					cause: "INVALID_EMAIL",
					msg:   "invalid email",
				}
			},
			want: func(args *args) *want {
				return &want{
					err: FailureCause(args.cause, args.msg),
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.args()
			want := test.want(args)

			err := FailureCause(args.cause, args.msg)

			a := assert.New(t)

			a.Equal(want.err.Error(), err.Error())
		})
	}
}

func TestUnauthorized(t *testing.T) {
	type args struct {
		msg string
	}

	type want struct {
		err error
	}

	tests := []struct {
		name string
		args func() *args
		want func(args *args) *want
	}{
		{
			name: "simple test unauthorized",
			args: func() *args {
				return &args{
					msg: "unauthorized",
				}
			},
			want: func(args *args) *want {
				return &want{
					err: Unauthorized(args.msg),
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.args()
			want := test.want(args)

			err := Unauthorized(args.msg)

			a := assert.New(t)

			a.Equal(want.err.Error(), err.Error())
		})
	}
}

func TestForbidden(t *testing.T) {
	type args struct {
		msg string
	}

	type want struct {
		err error
	}

	tests := []struct {
		name string
		args func() *args
		want func(args *args) *want
	}{
		{
			name: "simple test forbidden",
			args: func() *args {
				return &args{
					msg: "forbidden",
				}
			},
			want: func(args *args) *want {
				return &want{
					err: Forbidden(args.msg),
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.args()
			want := test.want(args)

			err := Forbidden(args.msg)

			a := assert.New(t)

			a.Equal(want.err.Error(), err.Error())
		})
	}
}

func TestIs(t *testing.T) {
	type args struct {
		err    error
		target error
	}

	type want struct {
		ok bool
	}

	tests := []struct {
		name string
		args func() *args
		want func(args *args) *want
	}{
		{
			name: "simple test forbidden",
			args: func() *args {
				return &args{
					err:    NotFound("test"),
					target: fiber.ErrNotFound,
				}
			},
			want: func(args *args) *want {
				return &want{
					ok: true,
				}
			},
		},
		{
			name: "simple test forbidden",
			args: func() *args {
				return &args{
					err:    NotFound("test"),
					target: fiber.ErrInternalServerError,
				}
			},
			want: func(args *args) *want {
				return &want{
					ok: false,
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.args()
			want := test.want(args)

			ok := Is(args.err, args.target)

			a := assert.New(t)

			a.Equal(want.ok, ok)
		})
	}
}
