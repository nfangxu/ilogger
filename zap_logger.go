package ilogger

import (
	"go.uber.org/zap"
)

type zapLogger struct {
	next *zap.SugaredLogger
}

// NewZapLogger ilogger Register this func with go-wire or other tools.
func NewZapLogger(log *zap.SugaredLogger) ILogger {
	return &zapLogger{next: log}
}

func (l *zapLogger) With(kvs M) ILogger {
	return &zapLogger{next: l.next.With(l.format(kvs)...)}
}

func (l *zapLogger) WithFunc(fn func() M) ILogger {
	return l.With(fn())
}

func (l *zapLogger) Error(msg string, err error, kvs M) {
	l.next.Errorw(msg, l.format(kvs.Add(M{"msg": msg, "error": err}))...)
}

func (l *zapLogger) Info(msg string, kvs M) {
	l.next.Infow(msg, l.format(kvs.Append("msg", msg))...)
}

func (l *zapLogger) Debug(msg string, kvs M) {
	l.next.Debugw(msg, l.format(kvs.Append("msg", msg))...)
}

func (l *zapLogger) format(kvs M) []any {
	format := make([]any, len(kvs)*2)
	for k, v := range kvs {
		format = append(format, k, v)
	}
	return format
}
