package entity

type Channelinfo struct {
	Title_chinese    string `json:"title_chinese"`
	Short_description_chinese    string `json:"short_description_chinese"`
	Visible    string `json:"visible"`
}

type Channelinfos []*Channelinfo