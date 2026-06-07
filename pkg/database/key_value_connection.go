package database

type KeyValueConnection interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Remove(key string) error
	Increment(key string) (int64, error)
}
