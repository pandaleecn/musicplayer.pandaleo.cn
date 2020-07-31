package service

import (
	"context"
	"fmt"
	"reflect"

	"musicplayer/entity"
	"musicplayer/sql"
	"strings"
)

type SongService struct {
	*sql.Service
	rec sql.Record
}

func NewSongService(db sql.Database) *SongService {
	return &SongService{Service: sql.NewService(db, new(entity.Song))}
}

func (s *SongService) Insert(ctx context.Context, e entity.Song) (int64, error) {
	if !e.ValidateInsert() {
		return 0, sql.ErrUnprocessable
	}

	q := fmt.Sprintf(`INSERT INTO %s (sheet_id, name, singer, link, image_url, description)
	VALUES (?,?,?,?,?,?);`, e.TableName())

	res, err := s.DB().Exec(ctx, q, e.SheetID, e.Name, e.Singer,e.Link, e.ImageURL, e.Description)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (s *SongService) BatchInsert(ctx context.Context, songs []entity.Song) (int, error) {
	if len(songs) == 0 {
		return 0, nil
	}

	var (
		valuesLines []string
		args        []interface{}
	)

	for _, s := range songs {
		if !s.ValidateInsert() {
			// all products should be "valid", we don't skip, we cancel.
			return 0, sql.ErrUnprocessable
		}

		valuesLines = append(valuesLines, "(?,?,?,?,?,?)")
		args = append(args, []interface{}{s.SheetID, s.Name, s.Singer, s.ImageURL, s.Description}...)
	}

	q := fmt.Sprintf("INSERT INTO %s (sheet_id, title, singer, link, image_url, description) VALUES %s;",
		s.RecordInfo().TableName(),
		strings.Join(valuesLines, ", "))

	res, err := s.DB().Exec(ctx, q, args...)
	if err != nil {
		return 0, err
	}

	n := sql.GetAffectedRows(res)
	return n, nil
}

func (s *SongService) Update(ctx context.Context, e entity.Song) (int, error) {
	q := fmt.Sprintf(`UPDATE %s
    SET
	    sheet_id = ?,
	    name = ?,
	    singer = ?,
		link = ?,
	    image_url = ?,
	    description = ?
	WHERE %s = ?;`, e.TableName(), e.PrimaryKey())

	res, err := s.DB().Exec(ctx, q, e.SheetID, e.Name, e.Singer, e.Link, e.Link, e.ImageURL, e.Description, e.ID)
	if err != nil {
		return 0, err
	}

	n := sql.GetAffectedRows(res)
	return n, nil
}

var songUpdateSchema = map[string]reflect.Kind{
	"sheet_id": reflect.Int,
	"name":       reflect.String,
	"singer":   reflect.String,
	"link":		reflect.String,
	"image_url":       reflect.String,
	"description": reflect.String,
}

// PartialUpdate accepts a key-value map to
// update the record based on the given "id".
func (s *SongService) PartialUpdate(ctx context.Context, id int64, attrs map[string]interface{}) (int, error) {
	return s.Service.PartialUpdate(ctx, id, songUpdateSchema, attrs)
}