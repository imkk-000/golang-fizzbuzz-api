package service

import "github.com/stretchr/testify/mock"

type SayingNumberRepositoryMock struct {
	mock.Mock
}

func (r *SayingNumberRepositoryMock) WriteCache(key string, value interface{}) error {
	o := r.Called(key, value)
	return o.Error(0)
}

func (r *SayingNumberRepositoryMock) ReadCache(key string) (data interface{}, err error) {
	o := r.Called(key)
	return o.Get(0), o.Error(1)
}
