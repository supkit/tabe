package log

import (
	rotate "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"time"
)

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getWriter(filename string) io.Writer {
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotate.New(
		filename+".%Y-%m-%d",
		rotate.WithMaxAge(time.Hour*24*7),
		rotate.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}

	return hook
}
