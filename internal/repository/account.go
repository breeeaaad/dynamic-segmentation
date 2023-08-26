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

func View(conn *pgx.Conn, id helpers.User, segment *[]string) error {
	rows, err := conn.Query(
		context.Background(),
		"select segment_name from adding join segments on adding.segment_id=segments.segment_id where user_id=$1",
		id.Id,
	)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			return err
		}
		*segment = append(*segment, s)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}
