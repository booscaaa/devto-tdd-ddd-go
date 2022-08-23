package e2etests_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Transaction struct {
	ID                 string        `json:"id"`
	Type               string        `json:"type"`
	Value              float64       `json:"value"`
	NumberInstallments int64         `json:"numberInstallments"`
	Installments       []Installment `json:"installments"`
}

type Installment struct {
	ID            string  `json:"id"`
	TransactionID string  `json:"transactionID"`
	Value         float64 `json:"value"`
}

const (
	BASE_URL = "http://localhost:3000"
)

func doRequest(method, url string, payload []byte) (*http.Response, error) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		return nil, err
	}

	client := http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func createTransaction(t *testing.T, value float64, numberInstallments int64) (*string, error) {
	payload := map[string]interface{}{
		"type":               "CREDIT_CARD",
		"value":              value,
		"numberInstallments": numberInstallments,
	}
	responsePayload := map[string]interface{}{}
	payloadJson, _ := json.Marshal(payload)

	response, err := doRequest(http.MethodPost, BASE_URL+"/", payloadJson)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Status code is not 200, its %v", response.StatusCode)
	}

	json.NewDecoder(response.Body).Decode(&responsePayload)
	transactionID := responsePayload["transactionID"].(string)
	return &transactionID, nil
}

func getTransactionByID(t *testing.T, transactionID string) (*Transaction, error) {
	response, err := doRequest(http.MethodGet, BASE_URL+"/transaction/"+transactionID, nil)
	responsePayload := Transaction{}

	if err != nil {
		return nil, err
	}

	json.NewDecoder(response.Body).Decode(&responsePayload)
	return &responsePayload, nil
}

func TestE2E(t *testing.T) {
	t.Run("Should create a transaction returning ID", func(t *testing.T) {
		transactionID, err := createTransaction(t, 630.21, 12)
		assert.Nil(t, err, fmt.Sprintf("Should error return null but return %s", err.Error()))
		assert.NotNil(t, transactionID, fmt.Sprintf("Should return transaction ID but return null"))
	})
}
