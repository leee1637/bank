package domain

import (
	"fmt"
)

func (u *User) Deposit(num float64) float64 {
	u.Balance = u.Balance + num
	return u.Balance
}

func (u *User) Withdraw(num float64) (float64, error) {
	if (u.Balance - num) < 0 {
		return 0, fmt.Errorf("Недостаточно средств!")
	}

	u.Balance = u.Balance - num
	return u.Balance, nil
}
