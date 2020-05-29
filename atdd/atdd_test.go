package atdd_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"fizzbuzz-api/model"
)

const HOST = "http://localhost:8080"

func TestATDDSendToSetModulationSequence(t *testing.T) {
	assert := assert.New(t)
	modulationSequence := &model.ModulationSequenceRequest{
		ModulationSequences: []model.ModulationSequence{
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
	url := fmt.Sprintf("%s/setModulationSequence", HOST)
	reqBody, err := json.Marshal(modulationSequence)
	assert.NoError(err)
	assert.NotEmpty(reqBody)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	assert.NoError(err)
	assert.NotNil(resp)
	assert.Equal(http.StatusOK, resp.StatusCode)

	respBody, err := ioutil.ReadAll(resp.Body)
	assert.NoError(err)
	assert.NotEmpty(respBody)

	actualResp := &model.ModulationSequenceResponse{}
	err = json.Unmarshal(respBody, actualResp)
	assert.NoError(err)

	assert.NotZero(actualResp.UpdateCount)
	assert.NotEmpty(actualResp.ModulationSequences)
	assert.Equal(modulationSequence.ModulationSequences, actualResp.ModulationSequences)
}

func TestATDDSendNumberToSayWithSuccess(t *testing.T) {
	assert := assert.New(t)
	sayNumbers := map[int]string{
		1:  "1",
		3:  "Fizz",
		5:  "Buzz",
		6:  "Fizz",
		10: "Buzz",
		15: "FizzBuzz",
	}

	for sayNumber, expectedSayNumber := range sayNumbers {
		url := fmt.Sprintf("%s/say/%d", HOST, sayNumber)
		resp, err := http.Get(url)
		assert.NoError(err)
		assert.NotNil(resp)
		assert.Equal(http.StatusOK, resp.StatusCode)

		respBody, err := ioutil.ReadAll(resp.Body)
		assert.NoError(err)
		assert.NotEmpty(respBody)

		actualSayNumber := &model.SayingNumberResponse{}
		err = json.Unmarshal(respBody, actualSayNumber)
		assert.NoError(err)

		assert.NotZero(actualSayNumber.SaidCount)
		assert.Equal(expectedSayNumber, actualSayNumber.Say)
	}
}
