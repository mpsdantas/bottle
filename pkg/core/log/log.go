package log

import (
	"context"

	"github.com/mpsdantas/bottle/pkg/core/application"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	Init()
}

func Init() {
	var config zap.Config

	fields := zap.Fields(
		zap.Field{
			Key:    "application",
			Type:   zapcore.StringType,
			String: application.Name,
		},
		zap.Field{
			Key:    "environment",
			Type:   zapcore.StringType,
			String: application.Environment,
		},
		zap.Field{
			Key:    "version",
			Type:   zapcore.StringType,
			String: application.Version,
		},
		zap.Field{
			Key:    "port",
			Type:   zapcore.StringType,
			String: application.Port,
		},
	)

	config = zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	if application.Environment != application.Local {
		config = zap.Config{
			Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
			Encoding:         "json",
			EncoderConfig:    encoderConfig,
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}

		config.EncoderConfig.StacktraceKey = "error.stack"
	}

	l, _ := config.Build(zap.AddCallerSkip(1), zap.AddStacktrace(zap.DPanicLevel), fields)

	zap.ReplaceGlobals(l)
}

var encoderConfig = zapcore.EncoderConfig{
	TimeKey:        "time",
	LevelKey:       "severity",
	NameKey:        "logger",
	CallerKey:      "caller",
	MessageKey:     "message",
	StacktraceKey:  "stacktrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    encodeLevel(),
	EncodeTime:     zapcore.RFC3339TimeEncoder,
	EncodeDuration: zapcore.MillisDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

func encodeLevel() zapcore.LevelEncoder {
	return func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		switch l {
		case zapcore.DebugLevel:
			enc.AppendString("DEBUG")
		case zapcore.InfoLevel:
			enc.AppendString("INFO")
		case zapcore.WarnLevel:
			enc.AppendString("WARNING")
		case zapcore.ErrorLevel:
			enc.AppendString("ERROR")
		case zapcore.DPanicLevel:
			enc.AppendString("CRITICAL")
		case zapcore.PanicLevel:
			enc.AppendString("ALERT")
		case zapcore.FatalLevel:
			enc.AppendString("EMERGENCY")
		}
	}
}

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	zap.L().Debug(msg, MergeDefaultFields(ctx, fields...)...)
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	zap.L().Info(msg, MergeDefaultFields(ctx, fields...)...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	zap.L().Warn(msg, MergeDefaultFields(ctx, fields...)...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	zap.L().Error(msg, MergeDefaultFields(ctx, fields...)...)
}

func DPanic(ctx context.Context, msg string, fields ...zap.Field) {
	zap.L().DPanic(msg, MergeDefaultFields(ctx, fields...)...)
}

func Panic(ctx context.Context, msg string, fields ...zap.Field) {
	zap.L().Panic(msg, MergeDefaultFields(ctx, fields...)...)
}

func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	zap.L().Fatal(msg, MergeDefaultFields(ctx, fields...)...)
}

func Sync() error {
	return zap.L().Sync()
}

func MergeDefaultFields(ctx context.Context, fields ...zap.Field) []zap.Field {
	var f []zap.Field

	requestid, ok := ctx.Value("requestid").(string)
	if requestid != "" && ok {
		f = append(f, String("x-request-id", requestid))
	}

	return append(f, fields...)
}
