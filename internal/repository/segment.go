package repository

import (
	"context"
	"time"

	"github.com/breeeaaad/dynamic-segmentation/internal/helpers"
	"github.com/jackc/pgx/v5"
)

func SegmentCr(conn *pgx.Conn, segment helpers.Segment) error {
	var id int
	if err := conn.QueryRow(
		context.Background(),
		"insert into segments(segment_name) values($1) returning segment_id",
		segment.Name,
	).Scan(&id); err != nil {
		return err
	}
	if segment.Percent != 0 {
		var count float32
		if err := conn.QueryRow(context.Background(),
			"select count(*) from users").Scan(&count); err != nil {
			return err
		}
		count = count / 100 * segment.Percent
		if _, err := conn.Exec(
			context.Background(),
			"insert into adding(user_id,segment_id,added_at) select user_id,$1,$2 from users limit $3",
			id, time.Now(), count,
		); err != nil {
			return err
		}
		if _, err := conn.Exec(
			context.Background(),
			"insert into history(user_id,segment_name,added_at) select user_id,segment_name,added_at from adding join segments on segments.segment_id=adding.segment_id where adding.segment_id=$1",
			id,
		); err != nil {
			return err
		}
	}
	return nil
}

func SegmentDel(conn *pgx.Conn, segment helpers.Segment) error {
	if _, err := conn.Exec(
		context.Background(),
		"delete from segments where segment_name=$1",
		segment.Name,
	); err != nil {
		return err
	}
	if _, err := conn.Exec(
		context.Background(),
		"update history set deleted_at=$1 where segment_name=$2",
		time.Now, segment.Name,
	); err != nil {
		return err
	}
	return nil
}
