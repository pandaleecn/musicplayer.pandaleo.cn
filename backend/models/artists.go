package models

import "github.com/jinzhu/gorm"

type Artist struct {
	gorm.Model

	Name     string `gorm:"not null VARCHAR(191)"`
	Poster	 string `gorm:"VARCHAR(191)"`
}