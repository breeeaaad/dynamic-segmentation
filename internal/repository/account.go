package repository

import (
	"context"

	"github.com/breeeaaad/dynamic-segmentation/internal/helpers"
	"github.com/jackc/pgx/v5"
)

func AddUser(conn *pgx.Conn, id helpers.User) error {
	_, err := conn.Exec(
		context.Background(),
		"insert into users(user_id) values($1)",
		id.Id,
	)
	return err
}
