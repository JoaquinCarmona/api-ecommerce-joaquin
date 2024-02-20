package infrastructure

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"log"
	"main/services/storage"
	"main/src/products/domain/models"
	"time"
)

type CartRepository struct {
	ProductRepository ProductRepository
}

func (c *CartRepository) AddProduct(cart models.Cart, product_id uuid.UUID, existentRelation models.CartProducts, typeSum string) (models.Cart, error) {
	db := storage.GetDB()

	if typeSum == "sum" {
		sqlStatement := `
			UPDATE cart_product 
			SET 
				qty = $1, 
				price_in_cents = $2,
				added_at = $3
			WHERE product_id = $4 and cart_id = $5
			RETURNING product_id
		`
		sum := existentRelation.Qty + 1

		err := db.QueryRow(
			sqlStatement,
			sum,
			existentRelation.PriceInCents,
			time.Now(),
			existentRelation.CartId,
			existentRelation.ProductId,
		).Scan(&existentRelation.ProductId)
		if err != nil {
			log.Println(err.Error())
			return cart, err
		}

		cart, err = c.GetInfo(cart.Id)
		if err != nil {
			return cart, err
		}

		return cart, nil
	} else {
		log.Println("Searching existent Product")

		var product models.Product

		err := db.QueryRow("SELECT id, price_in_cents FROM products WHERE id = $1", product_id).Scan(&product.Id, &product.PriceInCents)

		if err != nil {
			log.Println("Product not exists")
			return cart, err
		}

		sqlStatement := `
		INSERT INTO cart_product (cart_id, product_id, price_in_cents, qty, added_at) 
		VALUES ($1, $2, $3, $4, $5)
	`

		_, err = db.Exec(sqlStatement, cart.Id, product_id, product.PriceInCents, 1, time.Now())

		if err != nil {
			return cart, err
		}

		cart, err = c.GetInfo(cart.Id)
		if err != nil {
			return cart, err
		}

		return cart, nil
	}
}

func (c *CartRepository) GetById(id string) (models.Cart, error) {
	db := storage.GetDB()
	var cart models.Cart
	sqlStatement := `SELECT * FROM carts WHERE id = $1`
	err := db.QueryRow(sqlStatement, id).Scan(
		&cart.Id,
		&cart.TotalPriceInCents,
		&cart.Currency,
		&cart.Status,
		&cart.CreatedAt,
		&cart.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return cart, fmt.Errorf("cart not found with ID %s", id)
		}
		return cart, err
	}
	return cart, nil
}

func (c *CartRepository) GetInfo(id uuid.UUID) (models.Cart, error) {
	db := storage.GetDB()

	sqlStatement := `
        SELECT cp.product_id
        FROM cart_product cp
        WHERE cp.cart_id = $1
    `
	rows, err := db.Query(sqlStatement, id)
	if err != nil {
		return models.Cart{}, err
	}
	defer rows.Close()

	// Crear una lista de IDs de productos asociados al carrito
	var productIDs []uuid.UUID
	for rows.Next() {
		var productID uuid.UUID
		if err := rows.Scan(&productID); err != nil {
			return models.Cart{}, err
		}
		productIDs = append(productIDs, productID)
	}

	// Construir el modelo Cart completo con los productos asociados
	cart := models.Cart{Id: id}
	err = db.QueryRow("SELECT * FROM carts WHERE id = $1", cart.Id).Scan(
		&cart.Id,
		&cart.TotalPriceInCents,
		&cart.Currency,
		&cart.Status,
		&cart.CreatedAt,
		&cart.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return cart, fmt.Errorf("cart not found with ID %s", id)
		}
		return cart, err
	}
	cart.Products = make([]models.Product, len(productIDs))

	for i, productID := range productIDs {
		product, err := c.ProductRepository.GetById(productID)
		if err != nil {
			return models.Cart{}, err
		}
		cart.Products[i] = product
	}

	return cart, nil
}

func (c *CartRepository) RemoveProduct(cart models.Cart, productId uuid.UUID, existentRelation models.CartProducts) (models.Cart, error) {
	db := storage.GetDB()

	total := existentRelation.Qty - 1

	if total == 0 {
		sqlStatement := `DELETE FROM cart_product WHERE product_id = $1 AND cart_id = $2`
		result, err := db.Exec(sqlStatement, existentRelation.CartId, existentRelation.ProductId)

		if err != nil {
			log.Println("Error executing query:", err)
			return cart, err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Println("Error getting rowsaffected:", err)
			return cart, err
		}

		log.Printf(" %d rows deleted\n", rowsAffected)
	} else {
		sqlStatement := `
			UPDATE cart_product 
			SET 
				qty = $1, 
				price_in_cents = $2,
				added_at = $3
			WHERE product_id = $4 and cart_id = $5
			RETURNING product_id
		`

		err := db.QueryRow(
			sqlStatement,
			total,
			existentRelation.PriceInCents,
			time.Now(),
			existentRelation.CartId,
			existentRelation.ProductId,
		).Scan(&existentRelation.ProductId)
		if err != nil {
			log.Println(err.Error())
			return cart, err
		}
	}

	cart, err := c.GetInfo(cart.Id)
	if err != nil {
		log.Println("Error getting info:", err)
		return cart, err
	}

	return cart, nil

}

func (c *CartRepository) CreateCart(cart models.Cart) (models.Cart, error) {
	db := storage.GetDB()
	var existingCartID string
	err := db.QueryRow("SELECT id FROM carts WHERE id = $1", cart.Id).Scan(&existingCartID)

	sqlStatement := `
		INSERT INTO carts (
			total_price_in_cents,
			currency,
			status,
			created_at,
			updated_at
		) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id
	`
	err = db.QueryRow(
		sqlStatement,
		cart.TotalPriceInCents,
		cart.Currency,
		cart.Status,
		cart.CreatedAt,
		cart.UpdatedAt,
	).Scan(&cart.Id)
	if err != nil {
		log.Println("could not create the cart into database")
		return cart, err
	}
	return cart, nil
}

func (c *CartRepository) GetExistentRelation(cart models.Cart, productId uuid.UUID) (models.CartProducts, bool, error) {
	db := storage.GetDB()
	var relation models.CartProducts
	err := db.QueryRow("SELECT * FROM cart_product WHERE product_id = $1 AND cart_id =$2", productId, cart.Id).Scan(
		&relation.ProductId,
		&relation.CartId,
		&relation.Qty,
		&relation.PriceInCents,
		&relation.AddedAt,
	)

	if err != nil {
		log.Println("could not get the relation into database")

		return relation, false, err
	}

	return relation, true, nil
}
