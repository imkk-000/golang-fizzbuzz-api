package repository

type SayingNumberRepository interface {
	WriteCache(key string, value interface{}) error
	ReadCache(key string) (data interface{}, err error)
}
