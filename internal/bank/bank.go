package bank

import "bank/internal/domain"

type PaymentSystem struct {
	Users        map[string]*domain.User
	Transactions []domain.Transaction
}
