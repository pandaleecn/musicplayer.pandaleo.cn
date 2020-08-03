package validates

type CreateUpdateAlbumRequest struct {
	CreateUserId	uint	`json:"user_id" comment:"用户ID"`
	Name string `json:"name" validate:"required"  comment:"名字"`
	Songs	[]uint	`json:"songs" comment:"用户ID"`
	Cover		string `gorm:"json:"cover" comment:"用户ID"`
	ArtistID	uint	 	`gorm:"VARCHAR(191)"`
}