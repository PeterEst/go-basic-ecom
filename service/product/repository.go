package product

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/peterest/go-basic-ecom/types"
)

type ProductRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (repo *ProductRepository) GetProducts() ([]types.Product, error) {
	rows, err := repo.db.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	}

	products := []types.Product{}
	for rows.Next() {
		product, err := scanRowsIntoProduct(rows)

		if err != nil {
			return nil, err
		}

		products = append(products, *product)
	}

	return products, nil
}

func (repo *ProductRepository) CreateProduct(product types.CreateProductPayload) (*types.Product, error) {
	result, err := repo.db.Exec(
		"INSERT INTO products (name, description, image, price, quantity) VALUES (?, ?, ?, ?, ?)",
		product.Name,
		product.Description,
		product.Image,
		product.Price,
		product.Quantity,
	)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return repo.GetProductByID(int(id))
}

func (repo *ProductRepository) GetProductByID(id int) (*types.Product, error) {
	rows, err := repo.db.Query("SELECT * FROM products WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	product := new(types.Product)
	if rows.Next() {
		product, err = scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
	}

	if product.ID == 0 {
		return nil, sql.ErrNoRows
	}

	return product, nil
}

func (repo *ProductRepository) GetProductsByID(productIds []int) ([]types.Product, error) {
	placeholders := strings.Repeat("?,", len(productIds)-1)
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (%s?)", placeholders)

	stringIds := make([]interface{}, len(productIds))
	for i, id := range productIds {
		stringIds[i] = id
	}

	rows, err := repo.db.Query(query, stringIds...)
	if err != nil {
		return nil, err
	}

	products := []types.Product{}
	for rows.Next() {
		product, err := scanRowsIntoProduct(rows)

		if err != nil {
			return nil, err
		}

		products = append(products, *product)
	}

	return products, nil
}

func (repo *ProductRepository) UpdateProduct(product types.Product) (*types.Product, error) {
	_, err := repo.db.Exec(
		"UPDATE products SET name = ?, description = ?, image = ?, price = ?, quantity = ? WHERE id = ?",
		product.Name,
		product.Description,
		product.Image,
		product.Price,
		product.Quantity,
		product.ID,
	)

	if err != nil {
		return nil, err
	}

	return repo.GetProductByID(product.ID)
}
