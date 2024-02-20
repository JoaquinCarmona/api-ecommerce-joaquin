package infrastructure

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"log"
	"main/services/storage"
	"main/src/products/domain/models"
)

type ProductRepository struct {
}

func (p ProductRepository) CreateOrUpdate(product models.Product) (models.Product, error) {
	db := storage.GetDB()
	var existingID string
	err := db.QueryRow("SELECT id FROM products WHERE id = $1", product.Id).Scan(&existingID)
	switch {
	case err == sql.ErrNoRows:
		return p.createProduct(db, product)
	case err != nil:
		return product, err
	default:
		return p.updateProduct(db, product)
	}
}

func (p ProductRepository) List() ([]models.Product, error) {
	db := storage.GetDB()
	sqlStatement := `SELECT * FROM products`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(rows)

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Sku,
			&product.Stock,
			&product.ImageUrl,
			&product.Price,
			&product.PriceInCents,
			&product.Currency,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.DeletedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (p ProductRepository) GetById(id uuid.UUID) (models.Product, error) {
	db := storage.GetDB()
	var product models.Product
	sqlStatement := `SELECT * FROM products WHERE id = $1`
	err := db.QueryRow(sqlStatement, id).Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Sku,
		&product.Stock,
		&product.ImageUrl,
		&product.Price,
		&product.PriceInCents,
		&product.Currency,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.DeletedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return product, fmt.Errorf("product not found with ID %s", id)
		}
		return product, err
	}
	return product, nil

}

func (p ProductRepository) DeleteById(id uuid.UUID) error {
	db := storage.GetDB()
	sqlStatement := `DELETE FROM products WHERE id = $1`
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}

func (p ProductRepository) createProduct(db *sql.DB, product models.Product) (models.Product, error) {
	sqlStatement := `
		INSERT INTO products (
			name, 
			description,
			sku,
			stock, 
			image_url,
			price, 
			price_in_cents, 
			currency,
			created_at,
			updated_at
		) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
		RETURNING id
	`
	err := db.QueryRow(
		sqlStatement,
		product.Name,
		product.Description,
		product.Sku,
		product.Stock,
		product.ImageUrl,
		product.Price,
		product.PriceInCents,
		product.Currency,
		product.CreatedAt,
		product.UpdatedAt,
	).Scan(&product.Id)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p ProductRepository) updateProduct(db *sql.DB, product models.Product) (models.Product, error) {
	sqlStatement := `
		UPDATE products 
		SET 
			name = $2, 
			description = $3,
			sku = $4,
			stock = $5,
			image_url = $6,
			price = $7,
			price_in_cents = $8,
			currency = $9,
			updated_at = $10
		WHERE id = $1
		RETURNING id
	`
	err := db.QueryRow(
		sqlStatement,
		product.Id,
		product.Name,
		product.Description,
		product.Sku,
		product.Stock,
		product.ImageUrl,
		product.Price,
		product.PriceInCents,
		product.Currency,
		product.UpdatedAt,
	).Scan(&product.Id)
	if err != nil {
		return product, err
	}
	return product, nil
}
