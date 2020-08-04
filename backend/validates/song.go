package validates

type CreateUpdateSongRequest struct {
	Name			string `json:"name" comment:"歌名"`
	Url				string `json:"url"  comment:"歌曲地址"`
	Cover    		string `json:"cover" comment:"封面"`
	ArtistID    	uint	`json:"artist_id" comment:"歌手ID"`
	UploadUserID	uint	`json:"upload_user_id" comment:"用户ID"`
	Lrc    			string `json:"lrc" comment:"歌词"`
	AlbumID			uint `json:"album_id" comment:"歌词"`
}