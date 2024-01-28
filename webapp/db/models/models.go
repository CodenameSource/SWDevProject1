package models

import "time"

type Product struct {
	ID      uint
	Url     string
	Price   float64
	AddedAt time.Time
}
