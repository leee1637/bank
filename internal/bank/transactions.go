package bank

import (
	"bank/internal/domain"
	"fmt"
	"sync"
)

func (p *PaymentSystem) AddUser(u *domain.User) {
	p.Users[u.ID] = u
}

func (p *PaymentSystem) AddTransaction(t *domain.Transaction) {
	p.TransactionQueue = append(p.TransactionQueue, *t)
}

func (p *PaymentSystem) ProcessingTransactions(t domain.Transaction) error {
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

func (p *PaymentSystem) Worker(wg *sync.WaitGroup, ch <-chan domain.Transaction) {
	defer wg.Done()

	for v := range ch {
		err := p.ProcessingTransactions(v)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Обработал одну транзакцию!")
	}

}
