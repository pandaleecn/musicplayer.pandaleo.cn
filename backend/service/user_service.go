package service

import (
	"context"
	"fmt"
	"reflect"

	"musicplayer/entity"
	"musicplayer/sql"
	"strings"
)

type UserService struct {
	*sql.Service
	rec sql.Record
}

func NewUserService(db sql.Database) *UserService {
	return &UserService{Service: sql.NewService(db, new(entity.User))}
}

func (s *UserService) Insert(ctx context.Context, e entity.User) (int64, error) {
	if !e.ValidateInsert() {
		return 0, sql.ErrUnprocessable
	}

	q := fmt.Sprintf(`INSERT INTO %s (role_id, user_name, pass_word, avatar, age, signature)
	VALUES (?,?,?,?,?,?);`, e.TableName())

	res, err := s.DB().Exec(ctx, q, e.RoleID, e.UserName, e.PassWord, e.Avatar, e.Age, e.Signature)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (s *UserService) BatchInsert(ctx context.Context, users []entity.User) (int, error) {
	if len(users) == 0 {
		return 0, nil
	}

	var (
		valuesLines []string
		args        []interface{}
	)

	for _, u := range users {
		if !u.ValidateInsert() {
			// all products should be "valid", we don't skip, we cancel.
			return 0, sql.ErrUnprocessable
		}

		valuesLines = append(valuesLines, "(?,?,?,?,?)")
		args = append(args, []interface{}{u.RoleID, u.UserName, u.PassWord, u.Avatar, u.Age, u.Signature}...)
	}

	q := fmt.Sprintf("INSERT INTO %s (role_id, user_name, pass_word, avatar, age, signature) VALUES %s;",
		s.RecordInfo().TableName(),
		strings.Join(valuesLines, ", "))

	res, err := s.DB().Exec(ctx, q, args...)
	if err != nil {
		return 0, err
	}

	n := sql.GetAffectedRows(res)
	return n, nil
}

func (s *UserService) Update(ctx context.Context, e entity.User) (int, error) {
	q := fmt.Sprintf(`UPDATE %s
    SET
	    role_id = ?,
	    user_name = ?,
	    pass_word = ?,
	    avatar = ?,
	    age = ?,
		signature = ?
	WHERE %s = ?;`, e.TableName(), e.PrimaryKey())

	res, err := s.DB().Exec(ctx, q, e.RoleID, e.UserName, e.PassWord, e.Avatar, e.Age, e.Signature, e.ID)
	if err != nil {
		return 0, err
	}

	n := sql.GetAffectedRows(res)
	return n, nil
}

var userUpdateSchema = map[string]reflect.Kind{
	"roles_id": reflect.Int,
	"user_name":       reflect.String,
	"pass_word":   reflect.String,
	"avatar":   reflect.String,
	"age":       reflect.Int,
	"signature": reflect.String,
}

// PartialUpdate accepts a key-value map to
// update the record based on the given "id".
func (s *UserService) PartialUpdate(ctx context.Context, id int64, attrs map[string]interface{}) (int, error) {
	return s.Service.PartialUpdate(ctx, id, userUpdateSchema, attrs)
}