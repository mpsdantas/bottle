package log

import (
	"context"

	"github.com/mpsdantas/bottle/pkg/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
)

func init() {
	var config zap.Config
	fields := zap.Fields(
		zap.Field{
			Key:    "application",
			Type:   zapcore.StringType,
			String: env.Application,
		},
		zap.Field{
			Key:    "environment",
			Type:   zapcore.StringType,
			String: env.Environment,
		},
		zap.Field{
			Key:    "scope",
			Type:   zapcore.StringType,
			String: env.Scope,
		},
		zap.Field{
			Key:    "version",
			Type:   zapcore.StringType,
			String: env.Version,
		},
		zap.Field{
			Key:    "port",
			Type:   zapcore.StringType,
			String: env.Port,
		},
	)

	if env.Environment == env.Prod || env.Environment == env.Stage {
		config = zap.NewProductionConfig()
		config.EncoderConfig.StacktraceKey = "error.stack"
	} else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	logger, _ = config.Build(zap.AddCallerSkip(1), fields)
}

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Debug(msg, MergeDefaultFields(ctx, fields...)...)
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Info(msg, MergeDefaultFields(ctx, fields...)...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Warn(msg, MergeDefaultFields(ctx, fields...)...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Error(msg, MergeDefaultFields(ctx, fields...)...)
}

func DPanic(ctx context.Context, msg string, fields ...zap.Field) {
	logger.DPanic(msg, MergeDefaultFields(ctx, fields...)...)
}

func Panic(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Panic(msg, MergeDefaultFields(ctx, fields...)...)
}

func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Fatal(msg, MergeDefaultFields(ctx, fields...)...)
}

func Sync() error {
	return logger.Sync()
}

func MergeDefaultFields(ctx context.Context, fields ...zap.Field) []zap.Field {
	var f []zap.Field

	requestid, ok := ctx.Value("requestid").(string)
	if requestid != "" && ok {
		f = append(f, String("x-request-id", requestid))
	}

	return append(f, fields...)
}
