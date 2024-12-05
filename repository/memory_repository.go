package repository

import (
	"errors"
	"sync"
	"tx-parser/domain"
)

var (
	tranDBSync *sync.Map
)

func init() {
	tranDBSync = &sync.Map{}
}

type MemoryRepository struct {
	transactions *sync.Map
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		transactions: tranDBSync,
	}
}

func (r *MemoryRepository) SaveTransaction(address string, tx domain.Transaction) error {
	// Load existing transactions list
	transactionList, ok := r.transactions.Load(address)
	if !ok {
		return errors.New("invalid address")
	}

	// Assert the type safely
	transactions, ok := transactionList.([]domain.Transaction)
	if !ok {
		return errors.New("invalid transaction list type")
	}

	// Append the new transaction
	transactions = append(transactions, tx)

	// Store the updated list back into the map
	r.transactions.Store(address, transactions)
	return nil
}

func (r *MemoryRepository) GetTransactions(address string) ([]domain.Transaction, error) {
	// Load existing transactions list
	transactionList, ok := r.transactions.Load(address)
	if !ok {
		return nil, errors.New("invalid address")
	}

	// Assert the type safely
	transactions, ok := transactionList.([]domain.Transaction)
	if !ok {
		return nil, errors.New("invalid transaction list type")
	}

	return transactions, nil
}

func (r *MemoryRepository) Subscribe(address string) error {
	// Check if the address is already subscribed
	_, ok := r.transactions.Load(address)
	if ok {
		return errors.New("you subscribed to this address: " + address)
	}

	// Initialize an empty transaction list for this address
	transactions := make([]domain.Transaction, 0)
	r.transactions.Store(address, transactions)
	return nil
}
