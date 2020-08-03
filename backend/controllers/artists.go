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
* @api {post} /admin/artist/ 新建歌手
* @apiName 新建歌手
* @apiGroup Artist
* @apiVersion 1.0.0
* @apiDescription 新建歌手
* @apiSampleRequest /admin/artist/
* @apiParam {string} name 歌手名
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func CreateArtist(ctx iris.Context) {

	aul := new(validates.CreateUpdateArtistRequest)
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
	artist := models.NewArtistByStruct(aul)

	artist.CreateArtist(aul)
	ctx.StatusCode(iris.StatusOK)
	if artist.ID == 0 {
		_, _ = ctx.JSON(ApiResource(false, artist, "操作失败"))
		return
	} else {
		_, _ = ctx.JSON(ApiResource(true, nil, "操作成功"))
		return
	}

}

/**
* @api {get} /admin/artists 获取所有歌手
* @apiName 获取歌手
* @apiGroup Artists
* @apiVersion 1.0.0
* @apiDescription 获取所有歌手
* @apiSampleRequest /admin/artists
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission 登陆用户
 */
func GetAllArtists(ctx iris.Context) {
	offset := ctx.URLParamIntDefault("offset", 1)
	limit := ctx.URLParamIntDefault("limit", 15)
	orderBy := ctx.URLParam("orderBy")

	Id := ctx.Values().Get("auth_user_id").(uint)
	artists := models.GetAllArtist(Id, orderBy, offset, limit)

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(ApiResource(true, artistsTransform(artists), "操作成功"))

}

/**
* @api {get} /admin/artists/:id 根据id获歌手详情
* @apiName 根据id获取歌手详情
* @apiGroup Artists
* @apiVersion 1.0.0
* @apiDescription 根据id获取歌手信息
* @apiSampleRequest /admin/artists/:id
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission 登陆用户
 */
func GetArtistDetail(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")
	artist := models.NewArtist(id, "")
	artist.GetArtistId()

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(ApiResource(true, artistTransform(artist), "操作成功"))
}

/**
* @api {post} /admin/artists/:id/update 更新歌手
* @apiName 更新歌手
* @apiGroup Artists
* @apiVersion 1.0.0
* @apiDescription 更新歌手
* @apiSampleRequest /admin/artists/:id/update
* @apiParam {string} name 歌手名
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func UpdateArtist(ctx iris.Context) {
	aul := new(validates.CreateUpdateArtistRequest)

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
	artist := models.NewArtist(id, "")

	artist.UpdateArtist(aul)
	ctx.StatusCode(iris.StatusOK)
	if artist.ID == 0 {
		_, _ = ctx.JSON(ApiResource(false, artist, "操作失败"))
		return
	} else {
		_, _ = ctx.JSON(ApiResource(true, nil, "操作成功"))
		return
	}

}

/**
* @api {delete} /admin/artists/:id/delete 删除歌手
* @apiName 删除歌手
* @apiGroup Artists
* @apiVersion 1.0.0
* @apiDescription 删除歌手
* @apiSampleRequest /admin/artists/:id/delete
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func DeleteArtist(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")

	artist := models.NewArtist(id, "")
	artist.DeleteArtist()

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(ApiResource(true, nil, "删除成功"))
}

func artistsTransform(artists []*models.Artist) []*transformer.Artist {
	var ps []*transformer.Artist
	for _, artist := range artists {
		p := artistTransform(artist)
		ps = append(ps, p)
	}
	return ps
}

func artistTransform(artist *models.Artist) *transformer.Artist {
	p := &transformer.Artist{}
	g := gf.NewTransform(p, artist, time.RFC3339)
	_ = g.Transformer()
	return p
}
