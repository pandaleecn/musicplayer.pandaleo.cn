package models

import "github.com/jinzhu/gorm"

type Album struct {
	gorm.Model

	Name		string `gorm:"not null VARCHAR(191)"`
	Cover		string `gorm:"VARCHAR(191)"`
	ArtistID	int	 	`gorm:"VARCHAR(191)"`
}