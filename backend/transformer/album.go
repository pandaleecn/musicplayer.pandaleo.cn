package transformer

type Album struct {
	Id          uint
	Name        string
	CreateUserId		uint
	Songs		[]Song
	Cover		string
	ArtistID	uint
}
