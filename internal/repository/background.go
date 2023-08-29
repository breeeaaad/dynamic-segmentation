package repository

import (
	"context"
	"log"
	"time"

	config "github.com/breeeaaad/dynamic-segmentation/configs"
)

func Bg() {
	conn := config.Config()
	defer conn.Close(context.Background())
	if _, err := conn.Exec(context.Background(),
		"update history set deleted_at=$1 where user_id in (select user_id from adding where added_at+rm_at<$1) and segment_name in (select segment_name from adding join segments on segments.segment_id=adding.segment_id where added_at+rm_at<$1)",
		time.Now()); err != nil {
		log.Fatal(err)
	}
	if _, err := conn.Exec(context.Background(),
		"delete from adding where added_at+rm_at<$1",
		time.Now()); err != nil {
		log.Fatal(err)
	}
	ticker := time.NewTicker(24 * time.Hour)
	for _ = range ticker.C {
		if _, err := conn.Exec(context.Background(),
			"update history set deleted_at=$1 where user_id in (select user_id from adding where added_at+rm_at<$1) and segment_name in (select segment_name from adding join segments on segments.segment_id=adding.segment_id where added_at+rm_at<$1)",
			time.Now()); err != nil {
			log.Fatal(err)
		}
		if _, err := conn.Exec(context.Background(),
			"delete from adding where added_at+rm_at<$1",
			time.Now()); err != nil {
			log.Fatal(err)
		}
	}
}
