package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tx-parser/domain"
	"tx-parser/repository"
)

func TestSubscribe(t *testing.T) {
	repo := repository.NewMemoryRepository()
	parserService := NewParserService(repo)

	err := parserService.Subscribe("0x123abc")
	assert.NoError(t, err, "Expected no error while subscribing")

	err = parserService.Subscribe("0x123abc")
	assert.NoError(t, err, "Expected no error when subscribing to an already subscribed address")
}

func TestGetTransactions(t *testing.T) {
	repo := repository.NewMemoryRepository()
	parserService := NewParserService(repo)

	err := parserService.Subscribe("0x123abc")
	assert.NoError(t, err, "Expected no error while subscribing")

	tx := domain.Transaction{
		From:  "0xSender",
		To:    "0xReceiver",
		Value: "100",
		Hash:  "0x123abc",
	}

	// Save a transaction
	err = parserService.SaveTransaction("0x123abc", tx)
	assert.NoError(t, err, "Expected no error while saving transactions")

	transactions, err := parserService.GetTransactions("0x123abc")
	assert.NoError(t, err, "Expected no error while fetching transactions")
	assert.Equal(t, len(transactions), 1, "Expected 1 transaction")
}
