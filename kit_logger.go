package ilogger

import (
	kitlog "github.com/go-kit/kit/log"
	kitlevel "github.com/go-kit/kit/log/level"
)

type kitLogger struct {
	next kitlog.Logger
}

// NewKitLogger ilogger Register this func with go-wire or other tools.
func NewKitLogger(log kitlog.Logger) ILogger {
	return &kitLogger{next: log}
}

func (l *kitLogger) With(kvs M) ILogger {
	return &kitLogger{next: kitlog.With(l.next, l.format(kvs)...)}
}

func (l *kitLogger) WithFunc(fn func() M) ILogger {
	return l.With(fn())
}

func (l *kitLogger) Error(msg string, err error, kvs M) {
	_ = kitlevel.Error(l.next).Log(l.format(kvs.Add(M{"msg": msg, "error": err.Error()}))...)
}

func (l *kitLogger) Info(msg string, kvs M) {
	_ = kitlevel.Info(l.next).Log(l.format(kvs.Append("msg", msg))...)
}

func (l *kitLogger) Debug(msg string, kvs M) {
	_ = kitlevel.Debug(l.next).Log(l.format(kvs.Append("msg", msg))...)
}

func (l *kitLogger) format(kvs M) []any {
	format := make([]any, len(kvs)*2)
	for k, v := range kvs {
		format = append(format, k, v)
	}
	return format
}
