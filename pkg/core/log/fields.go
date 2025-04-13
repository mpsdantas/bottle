package log

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Skip() zap.Field {
	return zap.Skip()
}

func Binary(key string, val []byte) zap.Field {
	return zap.Binary(key, val)
}

func Bool(key string, val bool) zap.Field {
	return zap.Bool(key, val)
}

func Boolp(key string, val *bool) zap.Field {
	return zap.Boolp(key, val)
}

func ByteString(key string, val []byte) zap.Field {
	return zap.ByteString(key, val)
}

func Complex128(key string, val complex128) zap.Field {
	return zap.Complex128(key, val)
}

func Complex128p(key string, val *complex128) zap.Field {
	return zap.Complex128p(key, val)
}

func Complex64(key string, val complex64) zap.Field {
	return zap.Complex64(key, val)
}

func Complex64p(key string, val *complex64) zap.Field {
	return zap.Complex64p(key, val)
}

func Float64(key string, val float64) zap.Field {
	return zap.Float64(key, val)
}

func Float64p(key string, val *float64) zap.Field {
	return zap.Float64p(key, val)
}

func Float32(key string, val float32) zap.Field {
	return zap.Float32(key, val)
}

func Float32p(key string, val *float32) zap.Field {
	return zap.Float32p(key, val)
}

func Int(key string, val int) zap.Field {
	return Int64(key, int64(val))
}

func Intp(key string, val *int) zap.Field {
	return zap.Intp(key, val)
}

func Int64(key string, val int64) zap.Field {
	return zap.Int64(key, val)
}

func Int64p(key string, val *int64) zap.Field {
	return zap.Int64p(key, val)
}

func Int32(key string, val int32) zap.Field {
	return zap.Int32(key, val)
}

func Int32p(key string, val *int32) zap.Field {
	return zap.Int32p(key, val)
}

func Int16(key string, val int16) zap.Field {
	return zap.Int16(key, val)
}

func Int16p(key string, val *int16) zap.Field {
	return zap.Int16p(key, val)
}

func Int8(key string, val int8) zap.Field {
	return zap.Int8(key, val)
}

func Int8p(key string, val *int8) zap.Field {
	return zap.Int8p(key, val)
}

func String(key string, val string) zap.Field {
	return zap.String(key, val)
}

func Stringp(key string, val *string) zap.Field {
	return zap.Stringp(key, val)
}

func Uint(key string, val uint) zap.Field {
	return Uint64(key, uint64(val))
}

func Uintp(key string, val *uint) zap.Field {
	return zap.Uintp(key, val)
}

func Uint64(key string, val uint64) zap.Field {
	return zap.Uint64(key, val)
}

func Uint64p(key string, val *uint64) zap.Field {
	return zap.Uint64p(key, val)
}

func Uint32(key string, val uint32) zap.Field {
	return zap.Uint32(key, val)
}

func Uint32p(key string, val *uint32) zap.Field {
	return zap.Uint32p(key, val)
}

func Uint16(key string, val uint16) zap.Field {
	return zap.Uint16(key, val)
}

func Uint16p(key string, val *uint16) zap.Field {
	return zap.Uint16p(key, val)
}

func Uint8(key string, val uint8) zap.Field {
	return zap.Uint8(key, val)
}

func Uint8p(key string, val *uint8) zap.Field {
	return zap.Uint8p(key, val)
}

func Uintptr(key string, val uintptr) zap.Field {
	return zap.Uintptr(key, val)
}

func Uintptrp(key string, val *uintptr) zap.Field {
	return zap.Uintptrp(key, val)
}

func Reflect(key string, val interface{}) zap.Field {
	return zap.Reflect(key, val)
}

func Namespace(key string) zap.Field {
	return zap.Namespace(key)
}

func Stringer(key string, val fmt.Stringer) zap.Field {
	return zap.Stringer(key, val)
}

func Time(key string, val time.Time) zap.Field {
	return zap.Time(key, val)
}

func Timep(key string, val *time.Time) zap.Field {
	return zap.Timep(key, val)
}

func Stack(key string) zap.Field {
	return zap.Stack(key)
}

func StackSkip(key string, skip int) zap.Field {
	return zap.StackSkip(key, skip)
}

func Duration(key string, val time.Duration) zap.Field {
	return zap.Duration(key, val)
}

func Durationp(key string, val *time.Duration) zap.Field {
	return zap.Durationp(key, val)
}

func Object(key string, val zapcore.ObjectMarshaler) zap.Field {
	return zap.Object(key, val)
}

func Inline(val zapcore.ObjectMarshaler) zap.Field {
	return zap.Inline(val)
}

func Any(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

func Err(err error) zap.Field {
	return zap.Error(err)
}
