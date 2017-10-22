package cache

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io/ioutil"

	"github.com/bradfitz/gomemcache/memcache"
)

type OptionsMemcached struct {
	Server    []string
	PrefixKey string
}

type Memcached struct {
	client *memcache.Client
	prefix string
}

func NewMemcached(opt OptionsMemcached) *Memcached {
	return &Memcached{
		client: memcache.New(opt.Server...),
		prefix: opt.PrefixKey,
	}
}

func (m *Memcached) composeKey(key string) string {
	return fmt.Sprintf("%s:%s", m.prefix, key)
}

func (m *Memcached) prepareItem(key string, data []byte) *memcache.Item {

	var b bytes.Buffer

	k := m.composeKey(key)

	if len(data) > 1024 {
		w := zlib.NewWriter(&b)
		w.Write(data)
		w.Close()

		return &memcache.Item{
			Key:   k,
			Value: b.Bytes(),
		}
	}

	return &memcache.Item{
		Key:        k,
		Value:      data,
		Expiration: 3600,
	}
}

func (m *Memcached) readValue(data []byte) ([]byte, error) {
	b := bytes.NewBuffer(data)

	r, err := zlib.NewReader(b)
	if err != nil {
		return data, nil
	}

	return ioutil.ReadAll(r)
}

func (m *Memcached) Set(key string, data []byte) error {
	item := m.prepareItem(key, data)

	return m.client.Set(item)
}

func (m *Memcached) Get(key string) ([]byte, error) {
	k := m.composeKey(key)

	item, err := m.client.Get(k)
	if err != nil {
		return []byte{}, err
	}

	return m.readValue(item.Value)
}
