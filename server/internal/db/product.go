package db

import (
	"context"
	"fmt"

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

func (d *Database) AddToCart(userId string, productId string) error {
	_, err := d.Exec(
		context.Background(),
		"UPDATE users SET cart = array_append(cart, $1) where id = $2;",
		productId,
		userId,
	)

	return err
}

func (d *Database) ReadCart(cart []string) ([]string, error) {
	rows, err := d.Query(
		context.Background(),
		"SELECT name, price FROM products WHERE id = ANY($1);",
		cart,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []string{}
	for rows.Next() {
		var name string
		var price float64
		err = rows.Scan(&name, &price)
		products = append(products, fmt.Sprintf("%s: %.2f", name, price))
	}

	return products, nil
}
