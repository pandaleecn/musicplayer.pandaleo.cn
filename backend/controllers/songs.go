package controllers

import (
	"time"

	"github.com/go-playground/validator/v10"
	gf "github.com/snowlyg/gotransformer"
	"musicplayer.pandaleo.cn/backend/models"
	"musicplayer.pandaleo.cn/backend/transformer"
	"musicplayer.pandaleo.cn/backend/validates"

	"github.com/kataras/iris/v12"
)

/**
* @api {get} /admin/songs/:user_id 根据id获取歌曲
* @apiName 根据id获取歌曲
* @apiGroup Songs
* @apiVersion 1.0.0
* @apiDescription 根据用户id获取歌曲信息
* @apiSampleRequest /admin/songs/:user_id
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission 登陆用户
 */
func GetSongByUser(ctx iris.Context) {
	offset := ctx.URLParamIntDefault("offset", 1)
	limit := ctx.URLParamIntDefault("limit", 15)
	orderBy := ctx.URLParam("orderBy")

	Id := ctx.Values().Get("auth_user_id").(uint)
	songs := models.GetAllSongsByUserId(Id, orderBy, offset, limit)

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(ApiResource(true, songsTransform(songs), "操作成功"))

}

/**
* @api {get} /admin/songs/:id 根据id获取歌曲
* @apiName 根据id获取歌曲
* @apiGroup Songs
* @apiVersion 1.0.0
* @apiDescription 根据id获取歌曲信息
* @apiSampleRequest /admin/songs/:id
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission 登陆用户
 */
func GetSong(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")
	song := models.NewSong(id, "")
	song.GetSongById()

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(ApiResource(true, songTransform(song), "操作成功"))
}

/**
* @api {post} /admin/songs/ 新建歌曲
* @apiName 新建歌曲
* @apiGroup Songs
* @apiVersion 1.0.0
* @apiDescription 新建歌曲
* @apiSampleRequest /admin/songs/
* @apiParam {string} name 歌曲名
* @apiParam {string} url  歌曲链接
* @apiParam {string} cover  封面
* @apiParam {string} artist_id  歌手ID
* @apiParam {string} lrc  歌词
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func CreateSong(ctx iris.Context) {

	aul := new(validates.CreateUpdateSongRequest)
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

	userId := ctx.Values().Get("auth_user_id").(int)
	aul.UploadUserID = userId
	song := models.NewSongByStruct(aul)

	song.CreateSong(aul)
	ctx.StatusCode(iris.StatusOK)
	if song.ID == 0 {
		_, _ = ctx.JSON(ApiResource(false, song, "操作失败"))
		return
	} else {
		_, _ = ctx.JSON(ApiResource(true, nil, "操作成功"))
		return
	}

}

/**
* @api {post} /admin/songs/:id/update 更新账号
* @apiName 更新歌曲
* @apiGroup Songs
* @apiVersion 1.0.0
* @apiDescription 更新歌曲
* @apiSampleRequest /admin/songs/:id/update
* @apiParam {string} name 歌曲名
* @apiParam {string} url  歌曲链接
* @apiParam {string} cover  封面
* @apiParam {string} artist_id  歌手ID
* @apiParam {string} lrc  歌词
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func UpdateSong(ctx iris.Context) {
	aul := new(validates.CreateUpdateSongRequest)

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
	song := models.NewSong(id, "")

	song.UpdateSong(aul)
	ctx.StatusCode(iris.StatusOK)
	if song.ID == 0 {
		_, _ = ctx.JSON(ApiResource(false, song, "操作失败"))
		return
	} else {
		_, _ = ctx.JSON(ApiResource(true, nil, "操作成功"))
		return
	}

}

/**
* @api {delete} /admin/songs/:id/delete 删除歌曲
* @apiName 删除歌曲
* @apiGroup Songs
* @apiVersion 1.0.0
* @apiDescription 删除歌曲
* @apiSampleRequest /admin/songs/:id/delete
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func DeleteSong(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")

	song := models.NewSong(id, "")
	song.DeleteSong()

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(ApiResource(true, nil, "删除成功"))
}

/**
* @api {get} /songs 获取所有的歌曲
* @apiName 获取所有的歌曲
* @apiGroup Songs
* @apiVersion 1.0.0
* @apiDescription 获取所有的歌曲
* @apiSampleRequest /songs
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func GetAllSongs(ctx iris.Context) {
	offset := ctx.URLParamIntDefault("offset", 1)
	limit := ctx.URLParamIntDefault("limit", 15)
	name := ctx.URLParam("name")
	orderBy := ctx.URLParam("orderBy")

	songs := models.GetAllSongs(name, orderBy, offset, limit)

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(ApiResource(true, songsTransform(songs), "操作成功"))
}

func songsTransform(songs []*models.Song) []*transformer.Song {
	var ss []*transformer.Song
	for _, song := range songs {
		s := songTransform(song)
		ss = append(ss, s)
	}
	return ss
}

func songTransform(song *models.Song) *transformer.Song {
	s := &transformer.Song{}
	g := gf.NewTransform(s, song, time.RFC3339)
	_ = g.Transformer()
	return s
}
