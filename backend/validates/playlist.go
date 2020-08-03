package validates

type CreateUpdatePlaylistRequest struct {
	UserID	uint	`json:"user_id" comment:"用户ID"`
	Name string `json:"name" validate:"required"  comment:"名字"`
	Songs	[]uint	`json:"songs" comment:"用户ID"`
}