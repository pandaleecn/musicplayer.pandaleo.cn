package models

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	"musicplayer.pandaleo.cn/backend/sysinit"
	"musicplayer.pandaleo.cn/backend/validates"
	"time"
)

type Artist struct {
	gorm.Model

	Name     string `gorm:"not null VARCHAR(191)"`
	Poster	 string `gorm:"VARCHAR(191)"`
	CreateUserId uint `gorm:"VARCHAR(191)"`
}

func NewArtist(id uint, name string) *Artist {
	return &Artist{
		Model: gorm.Model{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: name,
	}
}

func NewArtistByStruct(vs *validates.CreateUpdateArtistRequest) *Artist {
	return &Artist{
		Model: gorm.Model{
			ID:        0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:	vs.Name,
		CreateUserId: vs.CreateUserId,
		Poster: vs.Poster,
	}
}

/**
 * 创建
 * @method CreateArtist
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func (a *Artist) CreateArtist(aul *validates.CreateUpdateArtistRequest) {
	if err := sysinit.Db.Create(a).Error; err != nil {
		color.Red(fmt.Sprintf("CreateArtistErr:%s \n ", err))
	}
	return
}

/**
 * 通过 id 删除歌曲
 * @method DeleteSongById
 */
func (a *Artist) DeleteArtist() {
	if err := sysinit.Db.Delete(a).Error; err != nil {
		color.Red(fmt.Sprintf("DeleteArtistByIdErr:%s \n ", err))
	}
}

/**
 * 更新
 * @method UpdateSong
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func (a *Artist) UpdateArtist(uj *validates.CreateUpdateArtistRequest) {
	if err := Update(a, uj); err != nil {
		color.Red(fmt.Sprintf("UpdateArtistErr:%s \n ", err))
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
func GetAllArtist(id uint, orderBy string, offset, limit int) []*Artist {
	var artists []*Artist
	q := GetAllList(orderBy, offset, limit)
	if err := q.Where("create_user_id = ?", id).Find(&artists).Error; err != nil {
		color.Red(fmt.Sprintf("GetAllArtistErr:%s \n ", err))
		return nil
	}
	return artists
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
func (a *Artist) GetArtistId() {
	IsNotFound(sysinit.Db.Where("id = ?", a.ID).First(a).Error)
}