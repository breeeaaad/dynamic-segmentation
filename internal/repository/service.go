package repository

import "github.com/jackc/pgx/v5"

type Service struct {
	conn *pgx.Conn
}

func New(conn *pgx.Conn) *Service {
	return &Service{
		conn: conn,
	}
}
