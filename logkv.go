package logkv

type LogKv interface {
	Error(msg string, err error, kvs Kvs)
	Info(msg string, kvs Kvs)
	Debug(msg string, kvs Kvs)
	With(kvs Kvs) LogKv
}

type NoLogKv struct{}

func (kv NoLogKv) Error(msg string, err error, kvs Kvs) {}

func (kv NoLogKv) Info(msg string, kvs Kvs) {}

func (kv NoLogKv) Debug(msg string, kvs Kvs) {}

func (kv NoLogKv) With(kvs Kvs) LogKv { return kv }
