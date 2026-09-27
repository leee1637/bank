package main

import (
	"bank/internal/bank"
	"bank/internal/domain"
	"fmt"
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

	t := domain.Transaction{
		FromID: "1",
		ToID:   "2",
		Amount: 500.0,
	}

	t2 := domain.Transaction{
		FromID: "2",
		ToID:   "1",
		Amount: 500.0,
	}

	p.AddTransaction(&t)
	p.AddTransaction(&t2)

	for _, v := range p.Transactions {
		err := p.ProcessingTransactions(&v)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(user1.Balance)
		fmt.Println(user2.Balance)
	}

}
