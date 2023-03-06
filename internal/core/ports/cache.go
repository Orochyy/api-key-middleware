package ports

type CacheItem interface {
	GetKey() string
	GetValue() []byte
	GetExpiration() int32
}

type Cache interface {
	Set(key string, value []byte, expiration int32) error
	Get(key string) (CacheItem, error)
	Exists(key string) bool
	Delete(key string) error
	Ping() error
}
