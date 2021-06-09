package logkv

type Kvs map[string]interface{}

func (kv Kvs) Append(key string, value interface{}) Kvs {
	return kv.Add(Kvs{key: value})
}

func (kv Kvs) Add(kvs Kvs) Kvs {
	newKvs := make(Kvs, len(kv)+len(kvs))

	for field, value := range kv {
		newKvs[field] = value
	}
	for field, value := range kvs {
		newKvs[field] = value
	}

	return newKvs
}

func (kv Kvs) Copy() Kvs {
	cpy := make(Kvs, len(kv))
	for k, v := range kv {
		cpy[k] = v
	}
	return cpy
}
