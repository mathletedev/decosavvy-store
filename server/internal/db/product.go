package db

import (
	"context"

	"github.com/mathletedev/decosavvy/internal/models"
)

func (d *Database) ReadProducts() ([]*models.Product, error) {
	rows, err := d.Query(
		context.Background(),
		"SELECT * FROM products;",
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []*models.Product{}
	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.Id, &product.Image, &product.Name, &product.Description, &product.Price)
		products = append(products, &product)
	}

	return products, nil
}
