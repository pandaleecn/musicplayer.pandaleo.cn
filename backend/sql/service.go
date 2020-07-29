package sql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type Service struct {
	db  Database
	rec Record // see `Count`, `List` and `DeleteByID` methods.
}

func NewService(db Database, of Record) *Service {
	return &Service{db: db, rec: of}
}

func (s *Service) DB() Database {
	return s.db
}

func (s *Service) RecordInfo() Record {
	return s.rec
}

var ErrNoRows = sql.ErrNoRows

func (s *Service) GetByID(ctx context.Context, dest interface{}, id int64) error {
	q := fmt.Sprintf("SELECT * FROM %s WHERE %s = ? LIMIT 1", s.rec.TableName(), s.rec.PrimaryKey())
	err := s.db.Get(ctx, dest, q, id)
	return err
}

func (s *Service) Count(ctx context.Context) (total int64, err error) {
	q := fmt.Sprintf("SELECT COUNT(DISTINCT %s) FROM %s", s.rec.PrimaryKey(), s.rec.TableName())
	if err = s.db.Select(ctx, &total, q); err == sql.ErrNoRows {
		err = nil
	}
	return
}

type ListOptions struct {
	Table         string // the table name.
	Offset        uint64 // inclusive.
	Limit         uint64
	OrderByColumn string
	Order         string // "ASC" or "DESC" (could be a bool type instead).
	WhereColumn   string
	WhereValue    interface{}
}

func (opt ListOptions) Where(colName string, colValue interface{}) ListOptions {
	opt.WhereColumn = colName
	opt.WhereValue = colValue
	return opt
}

func (opt ListOptions) BuildQuery() (q string, args []interface{}) {
	q = fmt.Sprintf("SELECT * FROM %s", opt.Table)

	if opt.WhereColumn != "" && opt.WhereValue != nil {
		q += fmt.Sprintf(" WHERE %s = ?", opt.WhereColumn)
		args = append(args, opt.WhereValue)
	}

	if opt.OrderByColumn != "" {
		q += fmt.Sprintf(" ORDER BY %s %s", opt.OrderByColumn, ParseOrder(opt.Order))
	}

	if opt.Limit > 0 {
		q += fmt.Sprintf(" LIMIT %d", opt.Limit) // offset below.
	}

	if opt.Offset > 0 {
		q += fmt.Sprintf(" OFFSET %d", opt.Offset)
	}

	return
}

func ParseListOptions(q url.Values) ListOptions {
	offset, _ := strconv.ParseUint(q.Get("offset"), 10, 64)
	limit, _ := strconv.ParseUint(q.Get("limit"), 10, 64)
	order := q.Get("order") // empty, asc(...) or desc(...).
	orderBy := q.Get("by")  // e.g. price

	return ListOptions{Offset: offset, Limit: limit, Order: order, OrderByColumn: orderBy}
}

func (s *Service) List(ctx context.Context, dest interface{}, opts ListOptions) error {
	// Set table and order by column from record info for `List` by options
	// so it can be more flexible to perform read-only calls of other table's too.
	if opts.Table == "" {
		// If missing then try to set it by record info.
		opts.Table = s.rec.TableName()
	}
	if opts.OrderByColumn == "" {
		if b, ok := s.rec.(Sorted); ok {
			opts.OrderByColumn = b.SortBy()
		}
	}

	q, args := opts.BuildQuery()
	return s.db.Select(ctx, dest, q, args...)
}

func (s *Service) DeleteByID(ctx context.Context, id int64) (int, error) {
	q := fmt.Sprintf("DELETE FROM %s WHERE %s = ? LIMIT 1", s.rec.TableName(), s.rec.PrimaryKey())
	res, err := s.db.Exec(ctx, q, id)
	if err != nil {
		return 0, err
	}

	return GetAffectedRows(res), nil
}

var ErrUnprocessable = errors.New("invalid entity")

func (s *Service) PartialUpdate(ctx context.Context, id int64, schema map[string]reflect.Kind, attrs map[string]interface{}) (int, error) {
	if len(schema) == 0 || len(attrs) == 0 {
		return 0, nil
	}

	var (
		keyLines []string
		values   []interface{}
	)

	for key, kind := range schema {
		v, ok := attrs[key]
		if !ok {
			continue
		}

		switch v.(type) {
		case string:
			if kind != reflect.String {
				return 0, ErrUnprocessable
			}
		case int:
			if kind != reflect.Int {
				return 0, ErrUnprocessable
			}
		case bool:
			if kind != reflect.Bool {
				return 0, ErrUnprocessable
			}
		}

		keyLines = append(keyLines, fmt.Sprintf("%s = ?", key))
		values = append(values, v)
	}

	if len(values) == 0 {
		return 0, nil
	}

	q := fmt.Sprintf("UPDATE %s SET %s WHERE %s = ?;",
		s.rec.TableName(), strings.Join(keyLines, ", "), s.rec.PrimaryKey())

	res, err := s.DB().Exec(ctx, q, append(values, id)...)
	if err != nil {
		return 0, err
	}

	n := GetAffectedRows(res)
	return n, nil
}

func GetAffectedRows(result sql.Result) int {
	if result == nil {
		return 0
	}

	n, _ := result.RowsAffected()
	return int(n)
}

const (
	ascending  = "ASC"
	descending = "DESC"
)

// ParseOrder accept an order string and returns a valid mysql ORDER clause.
// Defaults to "ASC". Two possible outputs: "ASC" and "DESC".
func ParseOrder(order string) string {
	order = strings.TrimSpace(order)
	if len(order) >= 4 {
		if strings.HasPrefix(strings.ToUpper(order), descending) {
			return descending
		}
	}

	return ascending
}