package bank

import (
	"bank/internal/domain"
	"fmt"
)

func (p *PaymentSystem) AddUser(u *domain.User) {
	p.Users[u.ID] = u
}

func (p *PaymentSystem) AddTransaction(t *domain.Transaction) {
	p.Transactions = append(p.Transactions, *t)
}

func (p *PaymentSystem) ProcessingTransactions(t *domain.Transaction) error {
	fromUser, ok := p.Users[t.FromID]
	if !ok {
		return fmt.Errorf("User not found")
	}

	toUser, ok := p.Users[t.ToID]
	if !ok {
		return fmt.Errorf("To pay User not found")
	}

	_, err := fromUser.Withdraw(t.Amount)
	if err != nil {
		return err
	}

	toUser.Deposit(t.Amount)

	return nil
}
