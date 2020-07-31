package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID          int64      `db:"id" json:"id"`
	RoleID  	int64	   `db:"roles_id" json:"roles_id"`
	UserName	string     `db:"user_name" json:"user_name"`
	PassWord    string     `db:"pass_word" json:"pass_word"`
	Avatar		string     `db:"avatar" json:"avatar"`
	Age			int64     `db:"age" json:"age"`
	Signature	string     `db:"signature" json:"signature"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
}

func (u User) TableName() string {
	return "users"
}

func (u *User) PrimaryKey() string {
	return "id"
}

func (u *User) SortBy() string {
	return "updated_at"
}

func (u *User) ValidateInsert() bool {
	return u.RoleID > 0 && u.UserName != "" && u.PassWord != "" && u.Avatar != "" && u.Signature != ""
}

func (u *User) Scan(rows *sql.Rows) error {
	u.CreatedAt = new(time.Time)
	u.UpdatedAt = new(time.Time)
	return rows.Scan(&u.ID, &u.RoleID, &u.UserName, &u.PassWord, &u.Avatar, &u.Signature, &u.CreatedAt, &u.UpdatedAt)
}

type Users []*User

func (us *Users) Scan(rows *sql.Rows) (err error) {
	cp := *us
	for rows.Next() {
		u := new(User)
		if err = u.Scan(rows); err != nil {
			return
		}
		cp = append(cp, u)
	}

	if len(cp) == 0 {
		return sql.ErrNoRows
	}

	*us = cp

	return rows.Err()
}
