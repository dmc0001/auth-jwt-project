package store

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/dmc0001/auth-jwt-project/internal/types"
	"github.com/dmc0001/auth-jwt-project/internal/validation"
)

type ProductModel struct {
	Db *sql.DB
}

func NewProductModel(db *sql.DB) (productModel *ProductModel) {
	return &ProductModel{
		Db: db,
	}
}

func (u *ProductModel) GetProductById(id int) (*types.Product, error) {
	product := &types.Product{}
	const q = `
		SELECT id, name, description, image, price, quantity,created_at
		FROM products
		WHERE id = ?
		LIMIT 1;
	`
	row := u.Db.QueryRow(q, id)
	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return product, nil
}

func (u *ProductModel) GetProductByName(name string) ([]types.Product, error) {
	const q = `
		SELECT id, name, description, image, price, quantity, created_at
		FROM products
		WHERE name = ?
		LIMIT 20;
	`
	rows, err := u.Db.Query(q, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanProductRows(rows)
}

func (u *ProductModel) GetProducts() ([]types.Product, error) {
	const q = `
		SELECT id, name, description, image, price, quantity, created_at
		FROM products
		LIMIT 20;
	`
	rows, err := u.Db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanProductRows(rows)
}

func (u *ProductModel) CreateProduct(payload types.CreateProductRequest) error {

	validProduct, err := validation.ValidateCreateProduct(&payload)
	if err != nil {
		return err
	}

	q := "INSERT INTO products (name, description, image, price, quantity,created_at) VALUES (?, ?, ?, ?, ?, ?)"

	_, err = u.Db.Exec(q,
		validProduct.Name,
		validProduct.Description,
		validProduct.Image,
		validProduct.Price,
		validProduct.Quantity,
		validProduct.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to insert product: %v", err)
	}
	return nil
}

func ScanProductRows(rows *sql.Rows) ([]types.Product, error) {
	var products []types.Product
	for rows.Next() {
		var product types.Product
		if err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Image,
			&product.Price,
			&product.Quantity,
			&product.CreatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}
