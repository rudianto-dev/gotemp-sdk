package logger

import (
	context "context"

	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

type Fields map[string]interface{}

type ILogger interface {
	Info(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	Infof(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	InfoWithField(fields Fields)
	ErrorWithField(fields Fields)

	InfoWithContext(ctx context.Context, stage utils.ErrorStage, args ...interface{})
	ErrorWithContext(ctx context.Context, stage utils.ErrorStage, args ...interface{})
	InfoContextWithField(ctx context.Context, fields Fields)
	ErrorContextWithField(ctx context.Context, fields Fields)

	CloseStream()
}
