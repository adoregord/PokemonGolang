package domain

import "time"

type Pokemon struct {
	ID             int
	Name           string
	Type           string
	CatchRate      float32
	IsRare         bool
	RegisteredDate time.Time
}
