package domain

type Parser interface {
	GetCurrentBlock() (int, error)
	Subscribe(address string) error
	GetTransactions(address string) ([]Transaction, error)
	SaveTransaction(address string, tx Transaction)
}
