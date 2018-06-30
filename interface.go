package steppingdb

type Storage interface {
	Get(key string) Value
	Set(key string, value Value)

	HMGet(key, field string) Value
	HMSet(key, field string, value Value)
	HMLen(key string) int
	HMKeys(key string) []string

	ArrayGet(key string, i int) Value
	ArraySet(key string, i int, value Value)
	ArrayLen(key string) int
	ArrayResize(key string, length int)
}
