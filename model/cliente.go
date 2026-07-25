package model

import "time"

// Definição da entidade cliente
type Cliente struct {
	ID           int
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}
