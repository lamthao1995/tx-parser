package service

import (
	"fmt"
	"sync"
	"tx-parser/config"
	"tx-parser/domain"
	"tx-parser/utils"
)

type ParserService struct {
	currentBlock  int
	subscriptions map[string]bool
	repo          domain.Repository
	mu            sync.RWMutex
}

func NewParserService(repo domain.Repository) *ParserService {
	return &ParserService{
		subscriptions: make(map[string]bool),
		repo:          repo,
	}
}

func (p *ParserService) GetCurrentBlock() (int, error) {
	// Lock the read access to ensure thread safety when reading the block number
	p.mu.RLock()
	defer p.mu.RUnlock()

	url := config.AppConfig.EthRPCURL
	resp, err := utils.JsonRPCRequest(url, map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_blockNumber",
		"params":  []interface{}{},
		"id":      utils.RandomID(),
	})
	if err != nil {
		return 0, err
	}

	blockHex, ok := resp["result"].(string)
	if !ok {
		return 0, fmt.Errorf("invalid result: %v", resp["result"])
	}

	block, err := utils.HexToInt(blockHex)
	if err != nil {
		return 0, fmt.Errorf("error converting block: %v", err)
	}

	return block, nil
}

func (p *ParserService) Subscribe(address string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.subscriptions[address] {
		return nil
	}
	p.subscriptions[address] = true
	return nil
}

func (p *ParserService) GetTransactions(address string) ([]domain.Transaction, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	// Retrieve transactions from the repository
	return p.repo.GetTransactions(address), nil
}

func (p *ParserService) SaveTransaction(address string, tx domain.Transaction) {
	// Save transaction using the repository
	p.repo.SaveTransaction(address, tx)
}
