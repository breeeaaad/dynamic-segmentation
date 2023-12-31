package repository

import (
	"context"
	"log"
	"time"
)

func (s *Service) Bg() {
	if _, err := s.conn.Exec(context.Background(),
		"update history set deleted_at=$1 where user_id in (select user_id from adding where added_at+rm_at<$1) and segment_name in (select segment_name from adding join segments on segments.segment_id=adding.segment_id where added_at+rm_at<$1)",
		time.Now()); err != nil {
		log.Print(err)
	}
	if _, err := s.conn.Exec(context.Background(),
		"delete from adding where added_at+rm_at<$1",
		time.Now()); err != nil {
		log.Print(err)
	}
	ticker := time.NewTicker(24 * time.Hour)
	for _ = range ticker.C {
		if _, err := s.conn.Exec(context.Background(),
			"update history set deleted_at=$1 where user_id in (select user_id from adding where added_at+rm_at<$1) and segment_name in (select segment_name from adding join segments on segments.segment_id=adding.segment_id where added_at+rm_at<$1)",
			time.Now()); err != nil {
			log.Print(err)
		}
		if _, err := s.conn.Exec(context.Background(),
			"delete from adding where added_at+rm_at<$1",
			time.Now()); err != nil {
			log.Print(err)
		}
	}
}
