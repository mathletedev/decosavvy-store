package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/mathletedev/decosavvy/internal/models"
)

func (d *Database) CreateUser(email string, avatar string) (string, error) {
	id := uuid.NewString()

	rows, err := d.Query(
		context.Background(),
		"INSERT INTO users (id, email, avatar) VALUES ($1, $2, $3);",
		id,
		email,
		avatar,
	)
	if err != nil {
		return "", err
	}

	defer rows.Close()

	return id, nil
}

func (d *Database) ReadUser(id string) (*models.User, error) {
	rows, err := d.Query(
		context.Background(),
		"SELECT email, avatar FROM users WHERE id=$1;",
		id,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, err
	}

	var user models.User
	err = rows.Scan(&user.Email, &user.Avatar)

	return &user, nil
}
