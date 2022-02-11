## Easier to record fields with logger

```bash
go get github.com/fangx-packages/logkv@v0.1.0
```

### Usage

- `github.com/go-kit/kit/log`

```go
package logger

import (
	logkv "github.com/fangx-packages/logkv"
	kitlog "github.com/go-kit/kit/log"
	kitlevel "github.com/go-kit/kit/log/level"
)

type kitLogger struct {
	next kitlog.Logger
}

// NewKitLogKv Register this func with go-wire or other tools.
func NewKitLogKv(log kitlog.Logger) logkv.LogKv {
	return &kitLogger{next: log}
}

func (l *kitLogger) With(kvs logkv.Kvs) logkv.LogKv {
	return &kitLogger{next: kitlog.With(l.next, l.format(kvs)...)}
}

func (l *kitLogger) WithFunc(fn func() logkv.Kvs) logkv.LogKv {
	return l.With(fn())
}

func (l *kitLogger) Error(msg string, err error, kvs logkv.Kvs) {
	_ = kitlevel.Error(l.next).Log(l.format(kvs.Add(logkv.Kvs{"msg": msg, "error": err}))...)
}

func (l *kitLogger) Info(msg string, kvs logkv.Kvs) {
	_ = kitlevel.Info(l.next).Log(l.format(kvs.Append("msg", msg))...)
}

func (l *kitLogger) Debug(msg string, kvs logkv.Kvs) {
	_ = kitlevel.Debug(l.next).Log(l.format(kvs.Append("msg", msg))...)
}

func (l *kitLogger) format(kvs logkv.Kvs) []interface{} {
	format := make([]interface{}, len(kvs)*2)
	for k, v := range kvs {
		format = append(format, k, v)
	}
	return format
}

```

- `go.uber.org/zap`

```go
package logger

import (
	logkv "github.com/fangx-packages/logkv"
	zap "go.uber.org/zap"
)

type zapLogger struct {
	next *zap.SugaredLogger
}

// NewZapLogKv Register this func with go-wire or other tools.
func NewZapLogKv(log *zap.SugaredLogger) logkv.LogKv {
	return &zapLogger{next: log}
}

func (l *zapLogger) With(kvs logkv.Kvs) logkv.LogKv {
	return &zapLogger{next: l.next.With(l.format(kvs)...)}
}

func (l *zapLogger) WithFunc(fn func() logkv.Kvs) logkv.LogKv {
	return l.With(fn())
}

func (l *zapLogger) Error(msg string, err error, kvs logkv.Kvs) {
	l.next.Errorw(msg, l.format(kvs.Add(logkv.Kvs{"msg": msg, "error": err}))...)
}

func (l *zapLogger) Info(msg string, kvs logkv.Kvs) {
	l.next.Infow(msg, l.format(kvs.Append("msg", msg))...)
}

func (l *zapLogger) Debug(msg string, kvs logkv.Kvs) {
	l.next.Debugw(msg, l.format(kvs.Append("msg", msg))...)
}

func (l *zapLogger) format(kvs logkv.Kvs) []interface{} {
	format := make([]interface{}, len(kvs)*2)
	for k, v := range kvs {
		format = append(format, k, v)
	}
	return format
}

```
