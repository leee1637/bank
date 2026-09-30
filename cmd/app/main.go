package main

import (
	"bank/internal/bank"
	"bank/internal/domain"
	"fmt"
	"sync"
)

func main() {
	user1 := domain.User{
		ID:      "1",
		Name:    "Вася",
		Balance: 1000.0,
	}

	user2 := domain.User{
		ID:      "2",
		Name:    "Жора",
		Balance: 0.0,
	}

	p := bank.PaymentSystem{
		Users: make(map[string]*domain.User),
	}

	p.AddUser(&user1)
	p.AddUser(&user2)

	ch := make(chan *domain.Transaction, 4)

	ch <- &domain.Transaction{
		FromID: "1",
		ToID:   "2",
		Amount: 500.0,
	}

	ch <- &domain.Transaction{
		FromID: "2",
		ToID:   "1",
		Amount: 500.0,
	}

	ch <- &domain.Transaction{
		FromID: "1",
		ToID:   "2",
		Amount: 500.0,
	}

	ch <- &domain.Transaction{
		FromID: "1",
		ToID:   "2",
		Amount: 5000.0,
	}

	wg := sync.WaitGroup{}

	close(ch)
	for i := 0; i < 3; i++ {
		wg.Go(func() {
			for t := range ch {
				worker(t, &p)
			}
		})

	}

	wg.Wait()

	fmt.Println(user1.Balance)
	fmt.Println(user2.Balance)
	// p.AddTransaction(&t)
	// p.AddTransaction(&t2)

}

func worker(ch *domain.Transaction, p *bank.PaymentSystem) {
	err := p.ProcessingTransactions(ch)
	if err != nil {
		fmt.Println(err)
	}
}
