package controllers

import (
	"github.com/go-playground/validator/v10"
	"log"
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
* @apiDescription 新建歌词
* @apiSampleRequest /admin/lyrics/
* @apiParam {string} name 歌词名
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func CreateLyric(ctx iris.Context) {

	aul := new(validates.CreateUpdateLyricRequest)
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
	log.Printf("%s", userId)
	aul.CreateUserId = userId
	lyric := models.NewLyricByStruct(aul)
	lyric.CreateLyric(aul)
	ctx.StatusCode(iris.StatusOK)
	if lyric.ID == 0 {
		_, _ = ctx.JSON(ApiResource(false, lyric, "操作失败"))
		return
	} else {
		_, _ = ctx.JSON(ApiResource(true, nil, "操作成功"))
		return
	}

}

/**
* @api {get} /admin/lyrics 获取所有歌词
* @apiName 获取歌词
* @apiGroup Lyrics
* @apiVersion 1.0.0
* @apiDescription 获取所有歌单
* @apiSampleRequest /admin/lyrics
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission 登陆用户
 */
func GetAllLyrics(ctx iris.Context) {
	offset := ctx.URLParamIntDefault("offset", 1)
	limit := ctx.URLParamIntDefault("limit", 15)
	orderBy := ctx.URLParam("orderBy")

	Id := ctx.Values().Get("auth_user_id").(uint)
	lyrics := models.GetAllLyric(Id, orderBy, offset, limit)

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(ApiResource(true, lyricsTransform(lyrics), "操作成功"))

}

/**
* @api {get} /admin/lyrics/:id 根据id获取歌词详情
* @apiName 根据id获取茖葱详情
* @apiGroup Lyrics
* @apiVersion 1.0.0
* @apiDescription 根据id获取歌词信息
* @apiSampleRequest /admin/lyrics/:id
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission 登陆用户
 */
func GetLyricDetail(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")
	lyric := models.NewLyric(id, "")
	lyric.GetLyricById()

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(ApiResource(true, lyricTransform(lyric), "操作成功"))
}

/**
* @api {post} /admin/lyrics/:id/update 更新歌词
* @apiName 更新歌词
* @apiGroup /Lyrics
* @apiVersion 1.0.0
* @apiDescription 更新歌词
* @apiSampleRequest /admin/lyrics/:id/update
* @apiParam {string} name 歌曲名
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func UpdateLyrics(ctx iris.Context) {
	aul := new(validates.CreateUpdateLyricRequest)

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
	lyric := models.NewLyric(id, "")

	lyric.UpdateLyric(aul)
	ctx.StatusCode(iris.StatusOK)
	if lyric.ID == 0 {
		_, _ = ctx.JSON(ApiResource(false, lyric, "操作失败"))
		return
	} else {
		_, _ = ctx.JSON(ApiResource(true, nil, "操作成功"))
		return
	}

}

/**
* @api {delete} /admin/lyrics/:id/delete 删除歌词
* @apiName 删除歌词
* @apiGroup Lyrics
* @apiVersion 1.0.0
* @apiDescription 删除歌曲
* @apiSampleRequest /admin/lyrics/:id/delete
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func DeleteLyrics(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")

	lyrics := models.NewLyric(id, "")
	lyrics.DeleteLyric()

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(ApiResource(true, nil, "删除成功"))
}

func lyricsTransform(lyrics []*models.Lyric) []*transformer.Lyric {
	var as []*transformer.Lyric
	for _, lyric := range lyrics {
		a := lyricTransform(lyric)
		as = append(as, a)
	}
	return as
}

func lyricTransform(lyric *models.Lyric) *transformer.Lyric {
	a := &transformer.Lyric{}
	g := gf.NewTransform(a, lyric, time.RFC3339)
	_ = g.Transformer()
	return a
}
