package model

type ModulationSequence struct {
	Key   int
	Value string
}

type ModulationSequenceRequest struct {
	ModulationSequences []ModulationSequence
}

type ModulationSequenceResponse struct {
	UpdateCount         int
	ModulationSequences []ModulationSequence
}
