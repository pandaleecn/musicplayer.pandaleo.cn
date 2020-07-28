package controllers

import (
	"errors"
	"github.com/kataras/iris"
	"musicplayer.pandaleo.cn/datamodels"
	"musicplayer.pandaleo.cn/service"
)

/*
type Song struct {
	datamodels.Song
}

func (s Song) IsValid() bool {
	return s.ID > 0
}

func (s Song) Dispatch(ctx context.Context)  {
	if !s.IsValid() {
		ctx.NotFound
		return
	}
	ctx.JSON(s, context.JSON{Indent: " "})
}
 */

type SongController struct {
	Service service.SongService
}

func (c *SongController) GetBy(id int64) (movie datamodels.Song, found bool)  {
	return c.Service.GetByID(id)
}

func (c *SongController) PutBy(ctx iris.Context, id int64) (datamodels.Song, error) {
	file, info, err := ctx.FormFile("poster")
	if err != nil {
		return datamodels.Song{}, errors.New("图片不存在，操作失败！")
	}

	file.Close()

	poster := info.Filename
	link := ctx.FormValue("link")

	return c.Service.UpdatePosterAndLinkByID(id, poster, link)
}

func (c *SongController) DeleteBy(id int64) interface{} {
	wasDel := c.Service.DeleteByID(id)
	if wasDel {
		return iris.Map{"deleted": id}
	}

	return iris.StatusBadRequest
}