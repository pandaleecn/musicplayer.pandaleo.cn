package entity

import (
	"database/sql"
	"time"
)

type Song struct {
	ID          int64      `db:"id" json:"id"`
	SheetID  	int64	   `db:"sheet_id" json:"sheet_id"`
	Name      	string     `db:"name" json:"name"`
	Singer		string     `db:"singer" json:"singer"`
	Link		string		`db:"link" json:"link"`
	ImageURL    string     `db:"image_url" json:"image_url"`
	Description string     `db:"description" json:"description"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
}

func (s Song) TableName() string {
	return "songs"
}

func (s *Song) PrimaryKey() string {
	return "id"
}

func (s *Song) SortBy() string {
	return "updated_at"
}

func (s *Song) ValidateInsert() bool {
	return s.SheetID > 0 && s.Name != "" && s.Singer != "" && s.Link != "" && s.ImageURL != "" && s.Description != ""
}

func (s *Song) Scan(rows *sql.Rows) error {
	s.CreatedAt = new(time.Time)
	s.UpdatedAt = new(time.Time)
	return rows.Scan(&s.ID, &s.SheetID, &s.Name, &s.Singer, &s.Link, &s.ImageURL, &s.Description, &s.CreatedAt, &s.UpdatedAt)
}

type Songs []*Song

func (ss *Songs) Scan(rows *sql.Rows) (err error) {
	cp := *ss
	for rows.Next() {
		s := new(Song)
		if err = s.Scan(rows); err != nil {
			return
		}
		cp = append(cp, s)
	}

	if len(cp) == 0 {
		return sql.ErrNoRows
	}

	*ss = cp

	return rows.Err()
}

