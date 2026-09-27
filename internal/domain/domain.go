package domain

import "sync"

type User struct {
	ID      string
	Name    string
	Balance float64

	Mu *sync.Mutex
}

type Transaction struct {
	FromID string
	ToID   string
	Amount float64
}
