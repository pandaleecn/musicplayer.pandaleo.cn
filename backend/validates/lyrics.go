package validates

type CreateUpdateLyricRequest struct {
	Name string `json:"name" validate:"required"  comment:"名字"`
	SongID	uint	`json:"song_id" comment:"用户ID"`
	Url	string	`json:"url"  comment:"名字"`
	CreateUserId uint `json:"create_user_id" comment:"用户ID"`
}