package models

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	Id           uuid.UUID  `json:"id"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Sku          string     `json:"sku"`
	Stock        int16      `json:"stock"`
	ImageUrl     string     `json:"image_url"`
	Price        float32    `json:"price"`
	PriceInCents int32      `json:"price_in_cents"`
	Currency     string     `json:"currency"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}
