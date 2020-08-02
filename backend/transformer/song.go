package transformer

type Song struct {
	Id		int
	Name     string
	Url      string
	Cover    string
	ArtistID int
	Artist   Artist
}