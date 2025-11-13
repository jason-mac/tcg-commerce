package models

import (
	"database/sql"
	//"errors"
	"time"
)

type ProductImage struct {
	ID        int
	ProductID int
	Created   time.Time
	Url       string
	SortOrder int
}

type ProductImageModel struct {
	DB *sql.DB
}

func (pi *ProductImageModel) Insert(productID int, url string) (int, error) {
	query := `
    INSERT INTO product_images (product_id, url, sort_order)
    VALUES (
      $1,
      $2,
      COALESCE((SELECT MAX(sort_order) + 1 FROM product_images WHERE product_id = $1), 1)
    )
    RETURNING id;
    `
	var id int
	err := pi.DB.QueryRow(query, productID, url).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (pi *ProductImageModel) GetAllUrls(productID int) ([]string, error) {
	query := `
		SELECT url
		FROM product_images
		WHERE product_id = $1
		ORDER BY sort_order ASC
	`

	rows, err := pi.DB.Query(query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []string
	for rows.Next() {
		var url string
		if err := rows.Scan(&url); err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

func (pi *ProductImageModel) GetFirst(productID int) (string, error) {
	query :=
		`
    SELECT url
    FROM product_images
    WHERE product_id = $1
    ORDER BY sort_order ASC
    LIMIT 1
  `

	var url string
	err := pi.DB.QueryRow(query, productID).Scan(&url)

	if err != nil {
		// error not found later
		return "", err
	}

	return url, nil
}
