package domain

import (
	"fmt"
)

func (u *User) Deposit(num float64) float64 {
	u.Mu.Lock()
	defer u.Mu.Unlock()
	u.Balance = u.Balance + num
	return u.Balance
}

func (u *User) Withdraw(num float64) (float64, error) {
	u.Mu.Lock()
	defer u.Mu.Unlock()
	if (u.Balance - num) < 0 {
		return 0, fmt.Errorf("No cash more!")
	}

	u.Balance = u.Balance - num
	return u.Balance, nil
}
