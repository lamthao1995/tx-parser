package domain

type Repository interface {
	SaveTransaction(address string, tx Transaction) error
	GetTransactions(address string) ([]Transaction, error)
	Subscribe(address string) error
}
