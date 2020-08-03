package validates

type CreateUpdateArtistRequest struct {
	CreateUserId	uint	`json:"create_user_id" comment:"用户ID"`
	Name string `json:"name" validate:"required"  comment:"名字"`
	Songs	[]uint	`json:"songs" comment:"用户ID"`
	Poster		string	`gorm:"VARCHAR(191)"`
}