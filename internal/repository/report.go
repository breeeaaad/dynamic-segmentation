package repository

import (
	"context"
	"time"
)

func (s *Service) GetAdd(date string, data *[][]string) error {
	rows, err := s.conn.Query(
		context.Background(),
		"select user_id,segment_name,added_at from history WHERE added_at::text LIKE $1",
		date+"-%",
	)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			id, name string
			time     time.Time
		)
		if err := rows.Scan(&id, &name, &time); err != nil {
			return err
		}
		*data = append(*data, []string{id, name, "add", time.String()})
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetDel(date string, datad *[][]string) error {
	rows, err := s.conn.Query(
		context.Background(),
		"select user_id,segment_name,deleted_at from history WHERE deleted_at::text LIKE $1",
		date+"-%",
	)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			id, name string
			time     time.Time
		)
		if err := rows.Scan(&id, &name, &time); err != nil {
			return err
		}
		*datad = append(*datad, []string{id, name, "delete", time.String()})
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}
