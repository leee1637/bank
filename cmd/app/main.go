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

	p.AddTransaction(&domain.Transaction{
		FromID: "1",
		ToID:   "2",
		Amount: 500.0,
	})

	p.AddTransaction(&domain.Transaction{
		FromID: "2",
		ToID:   "1",
		Amount: 500.0,
	})

	p.AddTransaction(&domain.Transaction{
		FromID: "1",
		ToID:   "2",
		Amount: 500.0,
	})

	p.AddTransaction(&domain.Transaction{
		FromID: "1",
		ToID:   "2",
		Amount: 500.0,
	})

	ch := make(chan domain.Transaction, len(p.TransactionQueue))

	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go p.Worker(&wg, ch)
	}

	for _, v := range p.TransactionQueue {
		ch <- v
	}

	close(ch)
	wg.Wait()

	fmt.Println(user1.Balance)
	fmt.Println(user2.Balance)

	// p.AddTransaction(&t)
	// p.AddTransaction(&t2)

}
