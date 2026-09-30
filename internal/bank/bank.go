package bank

import "bank/internal/domain"

type PaymentSystem struct {
	Users            map[string]*domain.User
	TransactionQueue []domain.Transaction
}
