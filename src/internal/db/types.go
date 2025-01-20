package db

type TransactionType string

const (
	Income  TransactionType = "income"
	Expense TransactionType = "expense"
)
