package validates

type CreateUpdateSongRequest struct {
	Name		string `json:"name" comment:"歌名"`
	Url			string `json:"url"  comment:"歌曲地址"`
	Cover    	string `json:"cover" comment:"封面"`
	ArtistID    int		`json:"artist_id" comment:"歌手ID"`
	Lrc    		string `json:"lrc" comment:"歌词"`
}