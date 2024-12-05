package repository

import (
	"tx-parser/domain"
)

var (
	memDB map[string][]domain.Transaction
)

func init() {
	memDB = make(map[string][]domain.Transaction)
}

type MemoryRepository struct {
	transactions map[string][]domain.Transaction
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		transactions: memDB,
	}
}

func (r *MemoryRepository) SaveTransaction(address string, tx domain.Transaction) error {
	r.transactions[address] = append(r.transactions[address], tx)
	return nil
}

func (r *MemoryRepository) GetTransactions(address string) ([]domain.Transaction, error) {
	data, _ := r.transactions[address]
	return data, nil
}
