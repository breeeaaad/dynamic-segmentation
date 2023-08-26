package repository

import (
	"context"
	"time"

	"github.com/breeeaaad/dynamic-segmentation/internal/helpers"
	"github.com/jackc/pgx/v5"
)

func Addsegments(conn *pgx.Conn, add helpers.Add) error {
	_, err := conn.Exec(
		context.Background(),
		"insert into adding(user_id,added_at,segment_id) select $1 , $2, segment_id from segments where segment_name=any($3)",
		add.Id, time.Now(), add.Addsegments,
	)
	return err
}

func Delsegments(conn *pgx.Conn, add helpers.Add) error {
	_, err := conn.Exec(
		context.Background(),
		"delete from adding where user_id=$1 and segment_id in (select segment_id from segments where segment_name=any($2))",
		add.Id, add.Delsegments,
	)
	return err
}
