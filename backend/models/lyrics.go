package models

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	"musicplayer.pandaleo.cn/backend/sysinit"
	"musicplayer.pandaleo.cn/backend/validates"
	"time"
)

type Lyric struct {
	gorm.Model
	Name	string	`gorm:"not null VARCHAR(191)"`
	Url		string	`gorm:"VARCHAR(191)"`
	SongID	uint		`gorm:"VARCHAR(191)"`
	CreateUserId uint `gorm:"VARCHAR(191)"`
}

func NewLyric(id uint, name string) *Lyric {
	return &Lyric{
		Model: gorm.Model{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: name,
	}
}

func NewLyricByStruct(vs *validates.CreateUpdateLyricRequest) *Lyric {
	return &Lyric{
		Model: gorm.Model{
			ID:        0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:	vs.Name,
		CreateUserId: vs.CreateUserId,
		Url: vs.Url,
		SongID: vs.SongID,
	}
}

/**
 * 创建
 * @method CreateLyric
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func (a *Lyric) CreateLyric(aul *validates.CreateUpdateLyricRequest) {
	if err := sysinit.Db.Create(a).Error; err != nil {
		color.Red(fmt.Sprintf("CreateLyricErr:%s \n ", err))
	}
	return
}

/**
 * 通过 id 删除歌曲
 * @method DeleteSongById
 */
func (a *Lyric) DeleteLyric() {
	if err := sysinit.Db.Delete(a).Error; err != nil {
		color.Red(fmt.Sprintf("DeleteLyricByIdErr:%s \n ", err))
	}
}

/**
 * 更新
 * @method UpdateSong
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func (a *Lyric) UpdateLyric(uj *validates.CreateUpdateLyricRequest) {
	if err := Update(a, uj); err != nil {
		color.Red(fmt.Sprintf("UpdateLyricErr:%s \n ", err))
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
func GetAllLyric(id uint, orderBy string, offset, limit int) []*Lyric {
	var lyrics []*Lyric
	q := GetAllList(orderBy, offset, limit)
	if err := q.Where("create_user_id = ?", id).Find(&lyrics).Error; err != nil {
		color.Red(fmt.Sprintf("GetAllAlyricErr:%s \n ", err))
		return nil
	}
	return lyrics
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
func (a *Lyric) GetLyricById() {
	IsNotFound(sysinit.Db.Where("id = ?", a.ID).First(a).Error)
}