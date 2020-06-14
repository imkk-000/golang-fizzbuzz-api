package service

import (
	"errors"
	"fizzbuzz-api/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCacheWithObject(t *testing.T) {
	modulationSequences := []model.ModulationSequence{
		{
			Key:   15,
			Value: "FizzBuzz",
		},
		{
			Key:   3,
			Value: "Fizz",
		},
		{
			Key:   5,
			Value: "Buzz",
		},
	}
	expectedData := model.ModulationSequenceResponse{
		UpdateCount:         1,
		ModulationSequences: modulationSequences,
	}
	sayingNumberRepository := &SayingNumberRepositoryMock{}
	sayingNumberRepository.On("ReadCache", "modulationSequence").Return(
		model.ModulationSequenceRequest{
			ModulationSequences: modulationSequences,
		}, nil)
	sayingNumberRepository.On("ReadCache", "updateCount").Return(1, nil)

	service := &SayingNumberServiceImpl{
		Repository: sayingNumberRepository,
	}

	actualData, _ := service.getCache()

	assert.Equal(t, expectedData, actualData)
}

func TestGetCacheWithErrorModulationSequence(t *testing.T) {
	expectedError := errors.New("not existing key: modulationSequence")
	sayingNumberRepository := &SayingNumberRepositoryMock{}
	sayingNumberRepository.On("ReadCache", "modulationSequence").Return(nil, expectedError)
	service := &SayingNumberServiceImpl{
		Repository: sayingNumberRepository,
	}

	actualData, actualError := service.getCache()

	assert.Equal(t, model.ModulationSequenceResponse{}, actualData)
	assert.Equal(t, expectedError, actualError)
}

func TestGetCacheWithErrorUpdateCount(t *testing.T) {
	expectedError := errors.New("not existing key: updateCount")
	sayingNumberRepository := &SayingNumberRepositoryMock{}
	sayingNumberRepository.On("ReadCache", "modulationSequence").Return([]model.ModulationSequenceRequest{}, nil)
	sayingNumberRepository.On("ReadCache", "updateCount").Return(nil, expectedError)
	service := &SayingNumberServiceImpl{
		Repository: sayingNumberRepository,
	}

	actualData, actualError := service.getCache()

	assert.Equal(t, model.ModulationSequenceResponse{}, actualData)
	assert.Equal(t, expectedError, actualError)
}

func TestSaveCacheWithObject(t *testing.T) {
	modulationSequenceRequest := model.ModulationSequenceRequest{
		ModulationSequences: []model.ModulationSequence{
			{
				Key:   15,
				Value: "FizzBuzz",
			},
			{
				Key:   3,
				Value: "Fizz",
			},
			{
				Key:   5,
				Value: "Buzz",
			},
		},
	}
	sayingNumberRepository := &SayingNumberRepositoryMock{}
	sayingNumberRepository.On("ReadCache", "updateCount").Return(0, nil)
	sayingNumberRepository.On("WriteCache", "modulationSequence", modulationSequenceRequest).Return(nil)
	sayingNumberRepository.On("WriteCache", "updateCount", 1).Return(nil)
	service := &SayingNumberServiceImpl{
		Repository: sayingNumberRepository,
	}

	actualError := service.saveCache(modulationSequenceRequest)

	assert.NoError(t, actualError)
}

func TestSaveCacheWithReadCacheError(t *testing.T) {
	expectedError := errors.New("not existing key: updateCount")
	sayingNumberRepository := &SayingNumberRepositoryMock{}
	sayingNumberRepository.On("ReadCache", "updateCount").Return(0, expectedError)
	service := &SayingNumberServiceImpl{
		Repository: sayingNumberRepository,
	}

	actualError := service.saveCache(model.ModulationSequenceRequest{})

	assert.Error(t, actualError)
	assert.Equal(t, expectedError, actualError)
}

func TestSaveCacheWithWriteCacheModulationSequenceError(t *testing.T) {
	expectedError := errors.New("cannot write cache with key: modulationSequence")
	sayingNumberRepository := &SayingNumberRepositoryMock{}
	sayingNumberRepository.On("ReadCache", "updateCount").Return(0, nil)
	sayingNumberRepository.On("WriteCache", "modulationSequence", model.ModulationSequenceRequest{}).Return(expectedError)
	service := &SayingNumberServiceImpl{
		Repository: sayingNumberRepository,
	}

	actualError := service.saveCache(model.ModulationSequenceRequest{})

	assert.Equal(t, expectedError, actualError)
}

func TestSaveCacheWithWriteCacheIncreasingUpdateCountError(t *testing.T) {
	expectedError := errors.New("cannot write cache with key: updateCount")
	sayingNumberRepository := &SayingNumberRepositoryMock{}
	sayingNumberRepository.On("ReadCache", "updateCount").Return(0, nil)
	sayingNumberRepository.On("WriteCache", "modulationSequence", model.ModulationSequenceRequest{}).Return(nil)
	sayingNumberRepository.On("WriteCache", "updateCount", 1).Return(expectedError)
	service := &SayingNumberServiceImpl{
		Repository: sayingNumberRepository,
	}

	actualError := service.saveCache(model.ModulationSequenceRequest{})

	assert.Equal(t, expectedError, actualError)
}
