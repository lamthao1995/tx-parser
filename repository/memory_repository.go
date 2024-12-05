
package repository

import "tx-parser/domain"

type MemoryRepository struct {
    transactions map[string][]domain.Transaction
}

func NewMemoryRepository() *MemoryRepository {
    return &MemoryRepository{
        transactions: make(map[string][]domain.Transaction),
    }
}

func (r *MemoryRepository) SaveTransaction(address string, tx domain.Transaction) {
    r.transactions[address] = append(r.transactions[address], tx)
}

func (r *MemoryRepository) GetTransactions(address string) []domain.Transaction {
    return r.transactions[address]
}
