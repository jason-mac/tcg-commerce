package models

import (
	"database/sql"
	//"errors"
	"math"
	"time"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
	Category    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ArchivedAt  time.Time
}

type ProductModel struct {
	DB *sql.DB
}

type ProductFilter struct {
	Name     string
	Category string
	MinPrice float64
	MaxPrice float64
	Sort     string
	Order    string
	Limit    int
	Offset   int
}

func (m *ProductModel) Insert(
	name string,
	description string,
	price float64,
	quantity int,
	category string,
) (int, error) {
	// pass
	return 0, nil
}

func (m *ProductModel) Get(id int) (Product, error) {
	query := `
		SELECT id, name, description, price, quantity, category, created, updated, archived_at
		FROM products
		WHERE id = $1 AND archived_at IS NULL
	`

	var p Product
	err := m.DB.QueryRow(query, id).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.Quantity,
		&p.Category,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.ArchivedAt,
	)

	if err != nil {
		return Product{}, err
	}

	return p, nil
}

func (m *ProductModel) Update(p *Product) error {
	query :=
		`
		UPDATE products
		SET name = $1, description = $2, price = $3, quantity = $4, category = $5, updated = NOW()
		WHERE id = $6 && 
	  `

	return nil
}

func (m *ProductModel) Filter(filter *ProductFilter) ([]*Product, error) {
	// query := `
	//         SELECT *
	//         FROM snippets
	//         WHERE id = $1
	//         `
	if filter.Category != "" {
		// do stuff
	}

	if filter.Name != "" {

	}

	if filter.MaxPrice != math.Inf(1) {

	}

	if filter.MinPrice != math.Inf(-1) {

	}
	return nil, nil
}

func (m *ProductModel) Delete(id int) error {
	return nil
}

func (m *ProductModel) Archive(id int, setArchived bool) error {
	var archive any

	if setArchived {
		now := time.Now()
		archive = now
	} else {
		archive = nil
	}

	query :=
		`
		UPDATE products
		SET archived_at = $1
		WHERE id = $2
	  `
	_, err := m.DB.Exec(query, archive, id)
	return err
}

func (m *ProductModel) Exists(id int) (bool, error) {
	var exists bool

	stmt := `SELECT EXISTS(SELECT true FROM users WHERE id = $1);`
	err := m.DB.QueryRow(stmt, id).Scan(&exists)
	return exists, err
}
