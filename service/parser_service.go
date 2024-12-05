package service

import (
	"fmt"
	"tx-parser/config"
	"tx-parser/domain"
	"tx-parser/utils"
)

type ParserService struct {
	repo domain.Repository
}

func NewParserService(repo domain.Repository) *ParserService {
	return &ParserService{
		repo: repo,
	}
}

func (p *ParserService) GetCurrentBlock() (int, error) {
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
	return p.repo.Subscribe(address)
}

func (p *ParserService) GetTransactions(address string) ([]domain.Transaction, error) {

	// Retrieve transactions from the repository
	return p.repo.GetTransactions(address)
}

func (p *ParserService) SaveTransaction(address string, tx domain.Transaction) error {

	return p.repo.SaveTransaction(address, tx)
}
