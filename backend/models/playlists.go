package models

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	"musicplayer.pandaleo.cn/backend/sysinit"
	"musicplayer.pandaleo.cn/backend/validates"
	"time"
)

type Playlist struct {
	gorm.Model
	UserId		uint	`gorm:"not null VARCHAR(191)"`
	Name		string 	`gorm:"not null VARCHAR(191)"`
	Songs 		[]Song 	`gorm:"ForeignKey:ID;AssociationForeignKey:ID"`
}

func NewPlaylist(id uint, name string) *Playlist {
	return &Playlist{
		Model: gorm.Model{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: name,
	}
}

func NewPlaylistByStruct(vs *validates.CreateUpdatePlaylistRequest) *Playlist {
	return &Playlist{
		Model: gorm.Model{
			ID:        0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:	vs.Name,
		UserId: vs.UserID,
	}
}

/**
 * 创建
 * @method CreateSong
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func (p *Playlist) CreatePlaylist(aul *validates.CreateUpdatePlaylistRequest) {
	if err := sysinit.Db.Create(p).Error; err != nil {
		color.Red(fmt.Sprintf("CreatePlaylistErr:%s \n ", err))
	}
	return
}

/**
 * 通过 id 删除歌曲
 * @method DeleteSongById
 */
func (p *Playlist) DeletePlaylist() {
	if err := sysinit.Db.Delete(p).Error; err != nil {
		color.Red(fmt.Sprintf("DeletePlaylistByIdErr:%s \n ", err))
	}
}

/**
 * 更新
 * @method UpdateSong
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func (p *Playlist) UpdatePlaylist(uj *validates.CreateUpdatePlaylistRequest) {
	if err := Update(p, uj); err != nil {
		color.Red(fmt.Sprintf("UpdatePlaylistErr:%s \n ", err))
	}
}

/**
 * 获取所有的歌单
 * @method GetAllUser
 * @param  {[type]} name string [description]
 * @param  {[type]} artist_id int [description]
 * @param  {[type]} orderBy string [description]
 * @param  {[type]} offset int    [description]
 * @param  {[type]} limit int    [description]
 */
func GetAllPlaylist(id uint, orderBy string, offset, limit int) []*Playlist {
	var playlists []*Playlist
	q := GetAllList(orderBy, offset, limit)
	if err := q.Where("user_id = ?", id).Find(&playlists).Error; err != nil {
		color.Red(fmt.Sprintf("GetAllPlayListErr:%s \n ", err))
		return nil
	}
	return playlists
}

/**
 * 根据ID获取歌单
 * @method GetAllUser
 * @param  {[type]} name string [description]
 * @param  {[type]} artist_id int [description]
 * @param  {[type]} orderBy string [description]
 * @param  {[type]} offset int    [description]
 * @param  {[type]} limit int    [description]
 */
func (p *Playlist) GetPlaylistById() {
	IsNotFound(sysinit.Db.Where("id = ?", p.ID).First(p).Error)
}