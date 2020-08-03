package controllers

import (
	"github.com/go-playground/validator/v10"
	"musicplayer.pandaleo.cn/backend/models"
	"musicplayer.pandaleo.cn/backend/validates"
	"time"

	gf "github.com/snowlyg/gotransformer"
	"musicplayer.pandaleo.cn/backend/transformer"

	"github.com/kataras/iris/v12"
)

/**
* @api {post} /admin/albums/ 新建专辑
* @apiName 新建专辑
* @apiGroup Albums
* @apiVersion 1.0.0
* @apiDescription 新建专辑
* @apiSampleRequest /admin/albums/
* @apiParam {string} name 专辑名
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func CreateAlbum(ctx iris.Context) {

	aul := new(validates.CreateUpdateAlbumRequest)
	if err := ctx.ReadJSON(aul); err != nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(ApiResource(false, nil, err.Error()))
		return
	}

	err := validates.Validate.Struct(*aul)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validates.ValidateTrans) {
			if len(e) > 0 {
				ctx.StatusCode(iris.StatusOK)
				_, _ = ctx.JSON(ApiResource(false, nil, e))
				return
			}
		}
	}

	userId := ctx.Values().Get("auth_user_id").(uint)
	aul.CreateUserId = userId
	album := models.NewAlbumByStruct(aul)

	album.CreateAlbum(aul)
	ctx.StatusCode(iris.StatusOK)
	if album.ID == 0 {
		_, _ = ctx.JSON(ApiResource(false, album, "操作失败"))
		return
	} else {
		_, _ = ctx.JSON(ApiResource(true, nil, "操作成功"))
		return
	}

}

/**
* @api {get} /admin/albums 获取所有专辑
* @apiName 获取专辑
* @apiGroup Albums
* @apiVersion 1.0.0
* @apiDescription 获取所有歌单
* @apiSampleRequest /admin/albums
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission 登陆用户
 */
func GetAllAlbums(ctx iris.Context) {
	offset := ctx.URLParamIntDefault("offset", 1)
	limit := ctx.URLParamIntDefault("limit", 15)
	orderBy := ctx.URLParam("orderBy")

	Id := ctx.Values().Get("auth_user_id").(uint)
	albums := models.GetAllAlbum(Id, orderBy, offset, limit)

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(ApiResource(true, albumsTransform(albums), "操作成功"))

}

/**
* @api {get} /admin/albums/:id 根据id获取专辑详情
* @apiName 根据id获取歌单详情
* @apiGroup Albums
* @apiVersion 1.0.0
* @apiDescription 根据id获取歌曲信息
* @apiSampleRequest /admin/albums/:id
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission 登陆用户
 */
func GetAlbumDetail(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")
	album := models.NewAlbum(id, "")
	album.GetAlbumById()

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(ApiResource(true, albumTransform(album), "操作成功"))
}

/**
* @api {post} /admin/albums/:id/update 更新账号
* @apiName 更新歌曲
* @apiGroup Albums
* @apiVersion 1.0.0
* @apiDescription 更新歌曲
* @apiSampleRequest /admin/albums/:id/update
* @apiParam {string} name 歌曲名
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func UpdateAlbums(ctx iris.Context) {
	aul := new(validates.CreateUpdateAlbumRequest)

	if err := ctx.ReadJSON(aul); err != nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(ApiResource(false, nil, err.Error()))
	}

	err := validates.Validate.Struct(*aul)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validates.ValidateTrans) {
			if len(e) > 0 {
				ctx.StatusCode(iris.StatusOK)
				_, _ = ctx.JSON(ApiResource(false, nil, e))
				return
			}
		}
	}

	id, _ := ctx.Params().GetUint("id")
	album := models.NewAlbum(id, "")

	album.UpdateAlbum(aul)
	ctx.StatusCode(iris.StatusOK)
	if album.ID == 0 {
		_, _ = ctx.JSON(ApiResource(false, album, "操作失败"))
		return
	} else {
		_, _ = ctx.JSON(ApiResource(true, nil, "操作成功"))
		return
	}

}

/**
* @api {delete} /admin/albums/:id/delete 删除歌单
* @apiName 删除歌单
* @apiGroup Albums
* @apiVersion 1.0.0
* @apiDescription 删除歌曲
* @apiSampleRequest /admin/albums/:id/delete
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func DeleteAlbums(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")

	albums := models.NewAlbum(id, "")
	albums.DeleteAlbum()

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(ApiResource(true, nil, "删除成功"))
}

func albumsTransform(albums []*models.Album) []*transformer.Album {
	var as []*transformer.Album
	for _, album := range albums {
		a := albumTransform(album)
		as = append(as, a)
	}
	return as
}

func albumTransform(album *models.Album) *transformer.Album {
	a := &transformer.Album{}
	g := gf.NewTransform(a, album, time.RFC3339)
	_ = g.Transformer()
	return a
}
