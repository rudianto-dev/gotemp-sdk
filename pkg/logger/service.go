package logger

import (
	"context"

	"github.com/rudianto-dev/gotemp-sdk/pkg/helper"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *Logger) Info(args ...interface{}) {
	s.Logrus.Info(args...)
}

func (s *Logger) Warning(args ...interface{}) {
	s.Logrus.Warning(args...)
}

func (s *Logger) Error(args ...interface{}) {
	s.Logrus.Error(args...)
}

func (s *Logger) Fatal(args ...interface{}) {
	s.Logrus.Fatal(args...)
}

func (s *Logger) Panic(args ...interface{}) {
	s.Logrus.Panic(args...)
}

func (s *Logger) Infof(format string, args ...interface{}) {
	s.Logrus.Infof(format, args...)
}

func (s *Logger) Warningf(format string, args ...interface{}) {
	s.Logrus.Warningf(format, args...)
}

func (s *Logger) Errorf(format string, args ...interface{}) {
	s.Logrus.Errorf(format, args...)
}

func (s *Logger) Fatalf(format string, args ...interface{}) {
	s.Logrus.Fatalf(format, args...)
}

func (s *Logger) Panicf(format string, args ...interface{}) {
	s.Logrus.Panicf(format, args...)
}

func (s *Logger) InfoWithField(fields Fields) {
	var arg map[string]interface{} = fields
	s.Logrus.WithFields(arg).Println()

}

func (s *Logger) ErrorWithField(fields Fields) {
	var arg map[string]interface{} = fields
	s.Logrus.WithFields(arg).Errorln()
}

func (s *Logger) InfoWithContext(ctx context.Context, stage utils.ErrorStage, args ...interface{}) {
	s.InfoContextWithField(ctx, map[string]interface{}{
		"stage":   stage,
		"message": args,
	})
}

func (s *Logger) ErrorWithContext(ctx context.Context, stage utils.ErrorStage, args ...interface{}) {
	s.ErrorContextWithField(ctx, map[string]interface{}{
		"stage":   stage,
		"message": args,
	})
}

func (s *Logger) InfoContextWithField(ctx context.Context, fields Fields) {
	var arg map[string]interface{} = fields
	arg["req_id"] = helper.GetReqID(ctx)
	s.Logrus.WithFields(arg).Println()
}

func (s *Logger) ErrorContextWithField(ctx context.Context, fields Fields) {
	var arg map[string]interface{} = fields
	arg["req_id"] = helper.GetReqID(ctx)
	s.Logrus.WithFields(arg).Errorln()
}

func (s *Logger) CloseStream() {
	s.Logrus.Writer().Close()
}
