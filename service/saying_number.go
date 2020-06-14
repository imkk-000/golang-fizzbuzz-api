package service

import (
	"fizzbuzz-api/model"
	"fizzbuzz-api/repository"
)

type SayingNumberService interface {
	getCache() (model.ModulationSequenceResponse, error)
	saveCache(data model.ModulationSequenceRequest) error
}

type SayingNumberServiceImpl struct {
	Repository repository.SayingNumberRepository
}

func (service SayingNumberServiceImpl) getCache() (model.ModulationSequenceResponse, error) {
	rawModulationSequences, err := service.Repository.ReadCache("modulationSequence")
	if err != nil {
		return model.ModulationSequenceResponse{}, err
	}
	rawUpdateCount, err := service.Repository.ReadCache("updateCount")
	if err != nil {
		return model.ModulationSequenceResponse{}, err
	}

	return model.ModulationSequenceResponse{
		UpdateCount:         rawUpdateCount.(int),
		ModulationSequences: rawModulationSequences.(model.ModulationSequenceRequest).ModulationSequences,
	}, nil
}

func (service SayingNumberServiceImpl) saveCache(data model.ModulationSequenceRequest) error {
	updateCount, err := service.Repository.ReadCache("updateCount")
	if err != nil {
		return err
	}
	err = service.Repository.WriteCache("modulationSequence", data)
	if err != nil {
		return err
	}
	err = service.Repository.WriteCache("updateCount", updateCount.(int)+1)

	return err
}
