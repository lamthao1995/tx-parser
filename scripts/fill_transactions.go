package scripts

import (
	"fmt"
	"log"
	"tx-parser/domain"
	"tx-parser/repository"
	"tx-parser/service"
)

func FillData() {
	// Initialize the repository and parser service
	repo := repository.NewMemoryRepository()
	parserService := service.NewParserService(repo)

	// Example subscriber addresses
	addresses := []string{"0x123abc", "0x456def", "0x789ghi"}

	// Subscribe addresses through parserService
	for _, addr := range addresses {
		err := parserService.Subscribe(addr)
		if err != nil {
			log.Fatalf("Failed to subscribe address %s: %v", addr, err)
		}
		log.Printf("Subscribed address: %s\n", addr)
	}

	// Mock transactions to fill for each address
	mockTransactions := []domain.Transaction{
		{From: "0xaaa", To: "0xbbb", Value: "100", Hash: "0xhash1"},
		{From: "0xccc", To: "0xddd", Value: "200", Hash: "0xhash2"},
		{From: "0xeee", To: "0xfff", Value: "300", Hash: "0xhash3"},
	}

	// Add transactions for each address through parserService
	for _, addr := range addresses {
		for _, txn := range mockTransactions {
			err := repo.SaveTransaction(addr, txn)
			if err != nil {
				fmt.Println("error = ", err)
			}
		}
	}

	log.Println("Finished populating transactions!")
}
