package domain

type Repository interface {
	SaveTransaction(address string, tx Transaction)
	GetTransactions(address string) []Transaction
}
