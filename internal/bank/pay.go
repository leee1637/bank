package bank

import (
	"bank/internal/domain"
	"fmt"
)

func Deposit(u *domain.User, num float64) float64 {
	u.Mu.Lock()
	defer u.Mu.Unlock()

	u.Balance = u.Balance + num

	return u.Balance
}

func Withdraw(u *domain.User, num float64) (float64, error) {
	u.Mu.Lock()
	defer u.Mu.Unlock()
	if (u.Balance - num) < 0 {
		return 0, fmt.Errorf("Недостаточно средств!")
	}

	u.Balance = u.Balance - num

	return u.Balance, nil
}
