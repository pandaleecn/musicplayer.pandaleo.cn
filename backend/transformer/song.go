package transformer

type Song struct {
	Id				uint
	Name			string
	Url				string
	Cover			string
	UploadUserID	uint
	ArtistID 		uint
	AlbumID 		uint
	Playlists		[]*Playlist
	AlbumName		string
	ArtistName		string
}