package models

import (
	"github.com/google/uuid"
	"time"
)

type Cart struct {
	Id                uuid.UUID `json:"id"`
	TotalPriceInCents int32     `json:"total_price_in_cents"`
	Currency          string    `json:"currency"`
	Status            string    `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Products          []Product
}

type CartProducts struct {
	CartId       uuid.UUID
	ProductId    uuid.UUID
	Qty          int
	PriceInCents int32
	AddedAt      time.Time
}
