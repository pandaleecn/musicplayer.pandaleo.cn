package models

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	"musicplayer.pandaleo.cn/backend/sysinit"
	"musicplayer.pandaleo.cn/backend/validates"
	"time"
)

type Album struct {
	gorm.Model
	Songs		[]Song	`gorm:"ForeignKey:SongsRefer"`
	Name		string `gorm:"not null VARCHAR(191)"`
	Cover		string `gorm:"VARCHAR(191)"`
	ArtistID	uint	`gorm:"VARCHAR(191)"`
	CreateUserId		uint	`gorm:"VARCHAR(191)"`
}

func NewAlbum(id uint, name string) *Album {
	return &Album{
		Model: gorm.Model{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: name,
	}
}

func NewAlbumByStruct(vs *validates.CreateUpdateAlbumRequest) *Album {
	return &Album{
		Model: gorm.Model{
			ID:        0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:	vs.Name,
		CreateUserId: vs.CreateUserId,
		Cover: vs.Cover,
		ArtistID: vs.ArtistID,
	}
}

/**
 * 创建
 * @method CreateAlbum
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func (a *Album) CreateAlbum(aul *validates.CreateUpdateAlbumRequest) {
	if err := sysinit.Db.Create(a).Error; err != nil {
		color.Red(fmt.Sprintf("CreateAlbumErr:%s \n ", err))
	}
	return
}

/**
 * 通过 id 删除歌曲
 * @method DeleteSongById
 */
func (a *Album) DeleteAlbum() {
	if err := sysinit.Db.Delete(a).Error; err != nil {
		color.Red(fmt.Sprintf("DeleteAlbumByIdErr:%s \n ", err))
	}
}

/**
 * 更新
 * @method UpdateSong
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func (a *Album) UpdateAlbum(uj *validates.CreateUpdateAlbumRequest) {
	if err := Update(a, uj); err != nil {
		color.Red(fmt.Sprintf("UpdateAlbumErr:%s \n ", err))
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
func GetAllAlbum(id uint, orderBy string, offset, limit int) []*Album {
	var albums []*Album
	q := GetAllList(orderBy, offset, limit)
	if err := q.Where("create_user_id = ?", id).Find(&albums).Error; err != nil {
		color.Red(fmt.Sprintf("GetAllAlbumErr:%s \n ", err))
		return nil
	}
	return albums
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
func (a *Album) GetAlbumById() {
	IsNotFound(sysinit.Db.Where("id = ?", a.ID).First(a).Error)
}