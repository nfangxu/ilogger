package ilogger

type ILogger interface {
	Error(msg string, err error, kvs M)
	Info(msg string, kvs M)
	Debug(msg string, kvs M)
	With(kvs M) ILogger
	WithFunc(fn func() M) ILogger
}

func None() ILogger {
	return &none{}
}

type none struct{}

func (l none) Error(msg string, err error, kvs M) {}

func (l none) Info(msg string, kvs M) {}

func (l none) Debug(msg string, kvs M) {}

func (l none) With(kvs M) ILogger { return l }

func (l none) WithFunc(fn func() M) ILogger { return l }

type M map[string]any

func (m M) Append(key string, value interface{}) M {
	return m.Add(M{key: value})
}

func (m M) Add(keyvals M) M {
	kvs := make(M, len(m)+len(keyvals))

	for field, value := range m {
		kvs[field] = value
	}
	for field, value := range keyvals {
		kvs[field] = value
	}

	return kvs
}

func (m M) Copy() M {
	cpy := make(M, len(m))
	for k, v := range m {
		cpy[k] = v
	}
	return cpy
}
