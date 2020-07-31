package entity

import (
	"database/sql"
	"time"
)

type Sheet struct {
	ID       int64  `db:"id" json:"id"`
	Title    string `db:"title" json:"title"`
	Position uint64 `db:"position" json:"position"`
	ImageURL string `db:"image_url" json:"image_url"`

	// We could use: sql.NullTime or unix time seconds (as int64),
	// note that the dsn parameter "parseTime=true" is required now in order to fill this field correctly.
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

func (s *Sheet) TableName() string {
	return "sheets"
}

func (s *Sheet) PrimaryKey() string {
	return "id"
}

func (s *Sheet) SortBy() string {
	return "position"
}

func (s *Sheet) Scan(rows *sql.Rows) error {
	s.CreatedAt = new(time.Time)
	s.UpdatedAt = new(time.Time)
	return rows.Scan(&s.ID, &s.Title, &s.Position, &s.ImageURL, &s.CreatedAt, &s.UpdatedAt)
}

type Sheets []*Sheet

func (ss *Sheets) Scan(rows *sql.Rows) (err error) {
	cp := *ss
	for rows.Next() {
		s := new(Sheet)
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