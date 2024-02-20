package usecases

import (
	"github.com/google/uuid"
	"main/src/products/domain/models"
)

type Cart interface {
	AddProduct(cart models.Cart, productId uuid.UUID, existentRelation models.CartProducts, typeSum string) (models.Cart, error)
	GetInfo(id uuid.UUID) (models.Cart, error)
	GetById(id string) (models.Cart, error)
	RemoveProduct(cart models.Cart, productId uuid.UUID, existentRelation models.CartProducts) (models.Cart, error)
	CreateCart(cart models.Cart) (models.Cart, error)
	GetExistentRelation(cart models.Cart, productId uuid.UUID) (models.CartProducts, bool, error)
}
