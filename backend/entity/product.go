package entity

import (
	"database/sql"
	"log"
	"time"
)

type Product struct {
	ID          int64      `db:"id" json:"id"`
	CategoryID  int64      `db:"category_id" json:"category_id"`
	Title       string     `db:"title" json:"title"`
	ImageURL    string     `db:"image_url" json:"image_url"`
	Price       float32    `db:"price" json:"price"`
	Description string     `db:"description" json:"description"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
}

func (p Product) TableName() string {
	return "products"
}

func (p *Product) PrimaryKey() string {
	return "id"
}

func (p *Product) SortBy() string {
	return "updated_at"
}

func (p *Product) ValidateInsert() bool {
	return p.CategoryID > 0 && p.Title != "" && p.ImageURL != "" && p.Price > 0 /* decimal* */ && p.Description != ""
}

func (p *Product) Scan(rows *sql.Rows) error {
	log.Printf("%s", rows)
	p.CreatedAt = new(time.Time)
	p.UpdatedAt = new(time.Time)
	return rows.Scan(&p.ID, &p.CategoryID, &p.Title, &p.ImageURL, &p.Price, &p.Description, &p.CreatedAt, &p.UpdatedAt)
}

type Products []*Product

func (ps *Products) Scan(rows *sql.Rows) (err error) {
	cp := *ps
	for rows.Next() {
		p := new(Product)
		if err = p.Scan(rows); err != nil {
			return
		}
		cp = append(cp, p)
	}

	if len(cp) == 0 {
		return sql.ErrNoRows
	}

	*ps = cp

	return rows.Err()
}

