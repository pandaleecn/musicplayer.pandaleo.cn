package entity

import (
"database/sql"
"time"
)

type Category struct {
	ID       int64  `db:"id" json:"id"`
	Title    string `db:"title" json:"title"`
	Position uint64 `db:"position" json:"position"`
	ImageURL string `db:"image_url" json:"image_url"`

	// We could use: sql.NullTime or unix time seconds (as int64),
	// note that the dsn parameter "parseTime=true" is required now in order to fill this field correctly.
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

func (c *Category) TableName() string {
	return "categories"
}

func (c *Category) PrimaryKey() string {
	return "id"
}

func (c *Category) SortBy() string {
	return "position"
}

func (c *Category) Scan(rows *sql.Rows) error {
	c.CreatedAt = new(time.Time)
	c.UpdatedAt = new(time.Time)
	return rows.Scan(&c.ID, &c.Title, &c.Position, &c.ImageURL, &c.CreatedAt, &c.UpdatedAt)
}

type Categories []*Category

func (cs *Categories) Scan(rows *sql.Rows) (err error) {
	cp := *cs
	for rows.Next() {
		c := new(Category)
		if err = c.Scan(rows); err != nil {
			return
		}
		cp = append(cp, c)
	}

	if len(cp) == 0 {
		return sql.ErrNoRows
	}

	*cs = cp

	return rows.Err()
}