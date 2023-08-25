package repository

import (
	"context"

	"github.com/breeeaaad/dynamic-segmentation/internal/helpers"
	"github.com/jackc/pgx/v5"
)

func SegmentCr(conn *pgx.Conn, segment helpers.Segment) error {
	_, err := conn.Exec(
		context.Background(),
		"insert into segments(segment_name) values($1)",
		segment.Name,
	)
	return err
}

func SegmentDel(conn *pgx.Conn, segment helpers.Segment) error {
	_, err := conn.Exec(
		context.Background(),
		"delete from segments where segment_name=$1",
		segment.Name,
	)
	return err
}
