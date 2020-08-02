package models

import "github.com/jinzhu/gorm"

type Playlist struct {
	gorm.Model

	Name		string `gorm:"not null VARCHAR(191)"`
}