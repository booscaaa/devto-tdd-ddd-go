package main_test

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

func createTransaction(t *testing.T, value float64, numberInstallments int64) *string {
	payload := map[string]interface{}{
		"type":               "CREDIT_CARD",
		"value":              value,
		"numberInstallments": numberInstallments,
	}
	responsePayload := map[string]interface{}{}

	payloadJson, _ := json.Marshal(payload)

	response, err := doRequest(http.MethodPost, BASE_URL+"/", payloadJson)

	if err != nil {
		t.Error(err)
		return nil
	}

	fmt.Println(response.StatusCode)

	assert.NotEqual(t, 200, response.StatusCode, "Erro")

	json.NewDecoder(response.Body).Decode(&responsePayload)
	transactionID := responsePayload["transactionID"].(string)
	return &transactionID
}

func getTransactionByID(t *testing.T, transactionID string) *Transaction {
	response, err := doRequest(http.MethodGet, BASE_URL+"/transaction/"+transactionID, nil)
	responsePayload := Transaction{}

	if err != nil {
		t.Error(err)
		return nil
	}

	json.NewDecoder(response.Body).Decode(&responsePayload)
	return &responsePayload
}

func TestMain(t *testing.T) {
	transactionID := createTransaction(t, 630.21, 12)
	transactionCreated := getTransactionByID(t, *transactionID)

	assert.NotNil(t, transactionCreated, "Ta nulo meu patrão!")
	assert.NotZero(t, transactionCreated.Installments, "Não tem parcela nenhuma aqui!")
	assert.Len(t, transactionCreated.Installments, 12, "Deveria ter 12 parcelas!")
	assert.Equal(t, 52.51, transactionCreated.Installments[0].Value, "O valor da primeira parcela ta errado!")
	assert.Equal(t, 52.6, transactionCreated.Installments[len(transactionCreated.Installments)-1].Value, "O valor da última parcela ta errado!")

	sumInstallments := 0.00
	for _, installment := range transactionCreated.Installments {
		sumInstallments += installment.Value
	}

	assert.Equal(t, 630.21, sumInstallments, "A soma das parelas não fechou meu patrão!")
}
