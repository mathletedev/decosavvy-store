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

func (d *Database) RemoveFromCart(userId string, productId string) error {
	_, err := d.Exec(
		context.Background(),
		`
UPDATE users 
SET cart = (
    SELECT array_agg(product_id)
    FROM (
        SELECT product_id
        FROM unnest(cart) WITH ORDINALITY AS t(product_id, ord)
        WHERE NOT (product_id = $2 AND ord = (
            SELECT MIN(ord) FROM unnest(cart) WITH ORDINALITY AS t2(product_id, ord)
            WHERE t2.product_id = $2
        ))
    ) AS filtered
)
WHERE id = $1;
		`,
		userId,
		productId,
	)

	fmt.Println(err)

	return err
}

func (d *Database) ReadCart(userId string) ([]*models.Product, error) {
	rows, err := d.Query(
		context.Background(),
		"SELECT p.id, p.image, p.name, p.description, p.price FROM users u JOIN LATERAL unnest(u.cart) AS product_id ON true JOIN products p ON p.id = product_id WHERE u.id = $1;",
		userId,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []*models.Product{}
	for rows.Next() {
		var product models.Product
		_ = rows.Scan(&product.Id, &product.Image, &product.Name, &product.Description, &product.Price)
		products = append(products, &product)
	}

	return products, nil
}
