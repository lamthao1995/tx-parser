package api

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"tx-parser/domain"
	"tx-parser/mocks"
)

func TestGetCurrentBlock(t *testing.T) {
	// Step 1: Create a mock Parser
	mockParser := new(mocks.Parser)

	// Step 2: Set up mock behavior
	mockParser.On("GetCurrentBlock").Return(12345, nil)

	// Step 3: Create a handler with the mock
	handler := NewHandler(mockParser)

	// Step 4: Simulate a request
	req := httptest.NewRequest(http.MethodGet, "/currentBlock", nil)
	rec := httptest.NewRecorder()

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	mux.ServeHTTP(rec, req)

	// Step 5: Assert the response
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, true, response["success"])
	assert.Equal(t, float64(12345), response["data"].(map[string]interface{})["currentBlock"])

	// Verify expectations on the mock
	mockParser.AssertExpectations(t)
}

func TestSubscribe(t *testing.T) {
	// Step 1: Create a mock Parser
	mockParser := new(mocks.Parser)

	// Step 2: Set up mock behavior
	mockParser.On("Subscribe", "0x123...abc").Return(nil)

	// Step 3: Create a handler with the mock
	handler := NewHandler(mockParser)

	// Step 4: Simulate a request
	req := httptest.NewRequest(http.MethodGet, "/subscribe?address=0x123...abc", nil)
	rec := httptest.NewRecorder()

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	mux.ServeHTTP(rec, req)

	// Step 5: Assert the response
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, true, response["success"])
	assert.Equal(t, "Address subscribed successfully", response["message"])

	// Verify expectations on the mock
	mockParser.AssertExpectations(t)
}

func TestGetTransactions(t *testing.T) {
	// Step 1: Create a mock Parser
	mockParser := new(mocks.Parser)

	// Step 2: Set up mock behavior
	mockParser.On("GetTransactions", "0x123...abc").Return([]domain.Transaction{
		{From: "0xSender", To: "0xReceiver", Value: "100", Hash: "0xabc123"},
	}, nil)

	// Step 3: Create a handler with the mock
	handler := NewHandler(mockParser)

	// Step 4: Simulate a request
	req := httptest.NewRequest(http.MethodGet, "/transactions?address=0x123...abc", nil)
	rec := httptest.NewRecorder()

	handler.RegisterRoutes(http.DefaultServeMux)
	http.DefaultServeMux.ServeHTTP(rec, req)

	// Step 5: Assert the response
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, true, response["success"])
	transactions := response["data"].([]interface{})
	assert.Equal(t, 1, len(transactions))
	txn := transactions[0].(map[string]interface{})
	assert.Equal(t, "0xSender", txn["from"])
	assert.Equal(t, "0xReceiver", txn["to"])
	assert.Equal(t, "100", txn["value"])
	assert.Equal(t, "0xabc123", txn["hash"])

	// Verify expectations on the mock
	mockParser.AssertExpectations(t)
}
