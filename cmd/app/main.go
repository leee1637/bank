package main

import (
	"bank/internal/domain"
	"bank/internal/pay"
	"fmt"
)

func main() {
	user1 := &pay.User{User: &domain.User{
		ID:      "1",
		Name:    "Вася",
		Balance: 1000.0,
	},
	}

	user2 := &pay.User{User: &domain.User{
		ID:      "2",
		Name:    "Жора",
		Balance: 0.0,
	},
	}

	user1.Deposit(22.0)

	fmt.Println(user1.Balance)

	_, err := user2.Withdraw(233.0)
	if err != nil {
		fmt.Println(err)
	}

	user2.Deposit(34223.0)

	fmt.Println(user2.Balance)

}
