package entity

import (
	"database/sql"
	"time"
)

type Role struct {
	ID          int64      `db:"id" json:"id"`
	Name      	string     `db:"name" json:"name"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
}

func (r Role) TableName() string {
	return "roles"
}

func (r *Role) PrimaryKey() string {
	return "id"
}

func (r *Role) SortBy() string {
	return "updated_at"
}

func (r *Role) Scan(rows *sql.Rows) error {
	r.CreatedAt = new(time.Time)
	r.UpdatedAt = new(time.Time)
	return rows.Scan(&r.ID, &r.Name, &r.CreatedAt, &r.UpdatedAt)
}

type Roles []*Role

func (rs *Roles) Scan(rows *sql.Rows) (err error) {
	cp := *rs
	for rows.Next() {
		r := new(Role)
		if err = r.Scan(rows); err != nil {
			return
		}
		cp = append(cp, r)
	}

	if len(cp) == 0 {
		return sql.ErrNoRows
	}

	*rs = cp

	return rows.Err()
}
