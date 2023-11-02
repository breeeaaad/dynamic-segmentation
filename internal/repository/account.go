package repository

import (
	"context"

	"github.com/breeeaaad/dynamic-segmentation/internal/helpers"
)

func (s *Service) AddUser() (int, error) {
	var id int
	err := s.conn.QueryRow(
		context.Background(),
		"insert into users(user_id) values(default) returning user_id",
	).Scan(&id)
	return id, err
}

func (s *Service) View(id helpers.User, segment *[]string) error {
	rows, err := s.conn.Query(
		context.Background(),
		"select segment_name from adding join segments on adding.segment_id=segments.segment_id where user_id=$1",
		id.Id,
	)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var b string
		if err := rows.Scan(&b); err != nil {
			return err
		}
		*segment = append(*segment, b)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}
