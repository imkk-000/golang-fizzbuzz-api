package atdd_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const HOST = "http://localhost:8080"

func TestATDDSendToUpdateModulationSequence(t *testing.T) {

}

func TestATDDSayNumber(t *testing.T) {
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

		actualSayNumber := struct {
			SaidCount int
			Say       string
		}{}
		err = json.Unmarshal(respBody, &actualSayNumber)
		assert.NoError(err)

		assert.NotZero(actualSayNumber.SaidCount)
		assert.Equal(expectedSayNumber, actualSayNumber.Say)
	}
}
