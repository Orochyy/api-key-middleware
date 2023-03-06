package cache

import (
	"api-key-middleware/internal/core/ports"
	"errors"
	"github.com/bradfitz/gomemcache/memcache"
)

type Item struct {
	Key        string
	Value      []byte
	Expiration int32
}

func (i Item) GetKey() string {
	return i.Key
}

func (i Item) GetValue() []byte {
	return i.Value
}

func (i Item) GetExpiration() int32 {
	return i.Expiration
}

type Cache struct {
	connection *memcache.Client
}

func NewCache(cacheHost string) *Cache {
	cache := memcache.New(cacheHost)

	return &Cache{connection: cache}
}

func (c *Cache) Set(key string, value []byte, expiration int32) error {
	return c.connection.Set(&memcache.Item{Key: key, Value: value, Expiration: expiration})
}

func (c *Cache) Get(key string) (ports.CacheItem, error) {
	memItem, err := c.connection.Get(key)
	if err != nil {
		return nil, err
	}

	item := Item{
		Key:        memItem.Key,
		Value:      memItem.Value,
		Expiration: memItem.Expiration,
	}

	return item, nil
}

func (c *Cache) Delete(key string) error {
	return c.connection.Delete(key)
}

func (c *Cache) Exists(key string) bool {
	_, err := c.connection.Get(key)
	if errors.Is(err, memcache.ErrCacheMiss) {
		return false
	}

	return true
}

func (c *Cache) Ping() error {
	return c.connection.Ping()
}
