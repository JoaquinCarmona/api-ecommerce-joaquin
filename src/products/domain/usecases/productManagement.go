package usecases

import (
	"github.com/google/uuid"
	"main/src/products/domain/models"
)

type Product interface {
	CreateOrUpdate(product models.Product) (models.Product, error)
	List() ([]models.Product, error)
	DeleteById(id uuid.UUID) error
	GetById(id uuid.UUID) (models.Product, error)
}
