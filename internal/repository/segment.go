package repository

import (
	"context"
	"errors"
	"time"

	"github.com/breeeaaad/dynamic-segmentation/internal/helpers"
)

func (s *Service) SegmentCr(segment helpers.Segment) error {
	var id int
	if err := s.conn.QueryRow(
		context.Background(),
		"insert into segments(segment_name) values($1) returning segment_id",
		segment.Name,
	).Scan(&id); err != nil {
		return err
	}
	if segment.Percent != 0 {
		var count float32
		if err := s.conn.QueryRow(context.Background(),
			"select count(*) from users").Scan(&count); err != nil {
			return err
		}
		count = count / 100 * segment.Percent
		if _, err := s.conn.Exec(
			context.Background(),
			"insert into adding(user_id,segment_id,added_at) select user_id,$1,$2 from users limit $3",
			id, time.Now(), count,
		); err != nil {
			return err
		}
		if _, err := s.conn.Exec(
			context.Background(),
			"insert into history(user_id,segment_name,added_at) select user_id,segment_name,added_at from adding join segments on segments.segment_id=adding.segment_id where adding.segment_id=$1",
			id,
		); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) SegmentDel(segment helpers.Segment) error {
	if tag, err := s.conn.Exec(
		context.Background(),
		"delete from segments where segment_name=$1",
		segment.Name,
	); err != nil {
		return err
	} else if tag.RowsAffected() == 0 {
		return errors.New("Не найден сегмент")
	}
	if _, err := s.conn.Exec(
		context.Background(),
		"update history set deleted_at=$1 where segment_name=$2",
		time.Now(), segment.Name,
	); err != nil {
		return err
	}
	return nil
}
