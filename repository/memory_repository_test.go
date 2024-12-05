package repository

import (
	"testing"
	"tx-parser/domain"

	"github.com/stretchr/testify/assert"
)

func TestSaveTransaction(t *testing.T) {
	repo := NewMemoryRepository()

	tx := domain.Transaction{
		From:  "0xSender",
		To:    "0xReceiver",
		Value: "100",
		Hash:  "0x123abc",
	}

	repo.SaveTransaction("0x123abc", tx)

	transactions := repo.GetTransactions("0x123abc")
	assert.Equal(t, len(transactions), 1, "Expected 1 transaction")
	assert.Equal(t, transactions[0].Hash, "0x123abc", "Expected matching transaction hash")
}
