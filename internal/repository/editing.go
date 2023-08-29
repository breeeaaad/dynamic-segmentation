package repository

import (
	"context"
	"time"

	"github.com/breeeaaad/dynamic-segmentation/internal/helpers"
	"github.com/jackc/pgx/v5"
)

func Addsegments(conn *pgx.Conn, add helpers.Add) error {
	if _, err := conn.Exec(
		context.Background(),
		"insert into adding(user_id,added_at,segment_id) select $1 , $2, segment_id from segments where segment_name=any($3)",
		add.Id, time.Now(), add.Addsegments,
	); err != nil {
		return err
	}
	if _, err := conn.Exec(
		context.Background(),
		"insert into history(user_id,added_at,segment_name) select $1,$2,segment_name from segments where segment_name=any($3)",
		add.Id, time.Now(), add.Addsegments,
	); err != nil {
		return err
	}
	return nil
}

func Delsegments(conn *pgx.Conn, add helpers.Add) error {
	if _, err := conn.Exec(context.Background(),
		"update history set deleted_at=$1 where user_id=$2 and segment_name=any($3)",
		time.Now(), add.Id, add.Delsegments,
	); err != nil {
		return err
	}
	if _, err := conn.Exec(
		context.Background(),
		"delete from adding where user_id=$1 and segment_id in (select segment_id from segments where segment_name=any($2))",
		add.Id, add.Delsegments,
	); err != nil {
		return err
	}
	return nil
}
