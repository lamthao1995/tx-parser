package repository

import (
	"testing"
	"tx-parser/domain"

	"github.com/stretchr/testify/assert"
)

func TestSaveTransaction(t *testing.T) {
	repo := NewMemoryRepository()
	err := repo.Subscribe("0x123abc")
	assert.NoError(t, err, "Expected no error while subscribe")

	tx := domain.Transaction{
		From:  "0xSender",
		To:    "0xReceiver",
		Value: "100",
		Hash:  "0x123abc",
	}

	err = repo.SaveTransaction("0x123abc", tx)
	assert.NoError(t, err, "Expected no error while SaveTransaction")

	transactions, _ := repo.GetTransactions("0x123abc")
	assert.Equal(t, len(transactions), 1, "Expected 1 transaction")
	assert.Equal(t, transactions[0].Hash, "0x123abc", "Expected matching transaction hash")
}
