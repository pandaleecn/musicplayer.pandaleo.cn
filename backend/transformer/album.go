package transformer

type Album struct {
	Id          int
	Name        string
	CreateUserId		uint
	Songs		[]Song
	Cover		string
	ArtistID	uint
}
