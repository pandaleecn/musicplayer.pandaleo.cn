package service

import (
	"context"
	"fmt"
	"reflect"

	"musicplayer/entity"
	"musicplayer/sql"
)

type RoleService struct {
	*sql.Service
}

func NewRoleService(db sql.Database) *RoleService {
	return &RoleService{Service: sql.NewService(db, new(entity.Role))}
}

func (s *RoleService) Insert(ctx context.Context, e entity.Role) (int64, error) {
	if e.Name == "" {
		return 0, sql.ErrUnprocessable
	}
	q := fmt.Sprintf(`INSERT INTO %s (name)
	VALUES (?);`, e.TableName())

	res, err := s.DB().Exec(ctx, q, e.Name)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (s *RoleService) Update(ctx context.Context, e entity.Role) (int, error) {
	if e.ID == 0 || e.Name == ""  {
		return 0, sql.ErrUnprocessable
	}

	q := fmt.Sprintf(`UPDATE %s
    SET
	    name = ?,
	WHERE %s = ?;`, e.TableName(), e.PrimaryKey())

	res, err := s.DB().Exec(ctx, q, e.Name, e.ID)
	if err != nil {
		return 0, err
	}

	n := sql.GetAffectedRows(res)
	return n, nil
}

var roleUpdateSchema = map[string]reflect.Kind{
	"name":     reflect.String,
}

func (s *RoleService) PartialUpdate(ctx context.Context, id int64, attrs map[string]interface{}) (int, error) {
	return s.Service.PartialUpdate(ctx, id, roleUpdateSchema, attrs)
}
