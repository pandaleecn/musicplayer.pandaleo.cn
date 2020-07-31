package service

import (
	"context"
	"fmt"
	"musicplayer/entity"
	"musicplayer/sql"
	"reflect"
)

type SheetService struct {
	*sql.Service
}

func NewSheetService(db sql.Database) *SheetService {
	return &SheetService{Service: sql.NewService(db, new(entity.Sheet))}
}

func (s *SheetService) Insert(ctx context.Context, e entity.Sheet) (int64, error) {
	if e.Title == "" || e.ImageURL == "" {
		return 0, sql.ErrUnprocessable
	}

	q := fmt.Sprintf(`INSERT INTO %s (title, position, image_url)
	VALUES (?,?,?);`, e.TableName())

	res, err := s.DB().Exec(ctx, q, e.Title, e.Position, e.ImageURL)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}


func (s *SheetService) Update(ctx context.Context, e entity.Sheet) (int, error) {
	if e.ID == 0 || e.Title == "" || e.ImageURL == "" {
		return 0, sql.ErrUnprocessable
	}

	q := fmt.Sprintf(`UPDATE %s
    SET
	    title = ?,
	    position = ?,
	    image_url = ?
	WHERE %s = ?;`, e.TableName(), e.PrimaryKey())

	res, err := s.DB().Exec(ctx, q, e.Title, e.Position, e.ImageURL, e.ID)
	if err != nil {
		return 0, err
	}

	n := sql.GetAffectedRows(res)
	return n, nil
}

var sheetUpdateSchema = map[string]reflect.Kind{
	"title":     reflect.String,
	"image_url": reflect.String,
	"position":  reflect.Int,
}

func (s *SheetService) PartialUpdate(ctx context.Context, id int64, attrs map[string]interface{}) (int, error) {
	return s.Service.PartialUpdate(ctx, id, sheetUpdateSchema, attrs)
}
