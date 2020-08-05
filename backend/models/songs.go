package models

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	"musicplayer.pandaleo.cn/backend/sysinit"
)

type Song struct {
	gorm.Model

	Name         string      `gorm:"not null VARCHAR(191)"`
	Url          string      `gorm:"VARCHAR(191)"`
	Cover        string      `gorm:"VARCHAR(191)"`
	AlbumID      uint        `gorm:"VARCHAR(191)"`
	ArtistID     uint        `gorm:"VARCHAR(191)"`
	UploadUserID uint        `gorm:"VARCHAR(191)"`
	Lrc          string      `gorm:"VARCHAR(191)"`
	Playlists    []*Playlist `gorm:"many2many:playlists_songs;"`
}

func NewSong(id uint, name string) *Song {
	return &Song{
		Model: gorm.Model{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: name,
	}
}

func NewSongByStruct(vs *CreateUpdateSongRequest) *Song {
	return &Song{
		Model: gorm.Model{
			ID:        0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:         vs.Name,
		Url:          vs.Url,
		Cover:        vs.Cover,
		AlbumID:      vs.AlbumID,
		ArtistID:     vs.ArtistID,
		UploadUserID: vs.UploadUserID,
	}
}

func (s *Song) GetSongBySongName() {
	IsNotFound(sysinit.Db.Where("name = ?", s.Name).First(s).Error)
}

func (s *Song) GetSongById() {
	IsNotFound(sysinit.Db.Where("id = ?", s.ID).First(s).Error)
}

func (s *Song) GetSongByUser() {
	IsNotFound(sysinit.Db.Where("upload_user_id = ?", s.UploadUserID).First(s).Error)
}

func (s *Song) GetSongByArtistId() {
	IsNotFound(sysinit.Db.Where("artist_id = ?", s.ID).First(s).Error)
}

/**
 * 通过 id 删除歌曲
 * @method DeleteSongById
 */
func (s *Song) DeleteSong() {
	if err := sysinit.Db.Delete(s).Error; err != nil {
		color.Red(fmt.Sprintf("DeleteSongByIdErr:%s \n ", err))
	}
}

/**
 * 获取所有的歌曲
 * @method GetAllUser
 * @param  {[type]} name string [description]
 * @param  {[type]} artist_id int [description]
 * @param  {[type]} orderBy string [description]
 * @param  {[type]} offset int    [description]
 * @param  {[type]} limit int    [description]
 */
func GetAllSongs(name, orderBy string, offset, limit int) []*Song {
	var songs []*Song
	q := GetAll(name, orderBy, offset, limit)
	if err := q.Find(&songs).Error; err != nil {
		color.Red(fmt.Sprintf("GetAllSongErr:%s \n ", err))
		return nil
	}
	return songs
}

/**
 * 获取所有的歌曲歌单
 * @method GetAllUser
 * @param  {[type]} name string [description]
 * @param  {[type]} artist_id int [description]
 * @param  {[type]} orderBy string [description]
 * @param  {[type]} offset int    [description]
 * @param  {[type]} limit int    [description]
 */
func GetAllPlaylistsBySong(id uint, orderBy string, offset, limit int) []*Playlist {
	song := NewSong(id, "")
	//q := GetAllList(orderBy, offset, limit)

	var playlists []*Playlist
	if err := sysinit.Db.Model(&song).Related(&playlists, "Playlists").Error; err != nil {
		color.Red(fmt.Sprintf("GetAllSongErr:%s \n ", err))
		return nil
	}
	return playlists
}

/**
 * 获取所有的歌曲
 * @method GetAllUser
 * @param  {[type]} name string [description]
 * @param  {[type]} artist_id int [description]
 * @param  {[type]} orderBy string [description]
 * @param  {[type]} offset int    [description]
 * @param  {[type]} limit int    [description]
 */
func GetAllSongsByUserId(id uint, orderBy string, offset, limit int) []*Song {
	var songs []*Song
	q := GetAllList(orderBy, offset, limit)
	if err := q.Where("upload_user_id = ?", id).Find(&songs).Error; err != nil {
		color.Red(fmt.Sprintf("GetAllSongErr:%s \n ", err))
		return nil
	}
	return songs
}

/**
 * 创建
 * @method CreateSong
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func (s *Song) CreateSong(aul *CreateUpdateSongRequest) {
	if err := sysinit.Db.Create(s).Error; err != nil {
		color.Red(fmt.Sprintf("CreateUserErr:%s \n ", err))
	}
	return
}

/**
 * 更新
 * @method UpdateSong
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func (s *Song) UpdateSong(uj *CreateUpdateSongRequest) {
	if err := Update(s, uj); err != nil {
		color.Red(fmt.Sprintf("UpdateSongErr:%s \n ", err))
	}
}

type CreateUpdateSongRequest struct {
	Name         string `json:"name" comment:"歌名"`
	Url          string `json:"url"  comment:"歌曲地址"`
	Cover        string `json:"cover" comment:"封面"`
	ArtistID     uint   `json:"artist_id" comment:"歌手ID"`
	UploadUserID uint   `json:"upload_user_id" comment:"用户ID"`
	Lrc          string `json:"lrc" comment:"歌词"`
	AlbumID      uint   `json:"album_id" comment:"歌词"`
	PlaylistsIds []uint `json:"playlist_ids" comment:"权限"`
	Playlists    []*Playlist
}
