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
* @api {post} /admin/playlists/ 新建歌单
* @apiName 新建歌曲
* @apiGroup Playlists
* @apiVersion 1.0.0
* @apiDescription 新建歌曲
* @apiSampleRequest /admin/playlists/
* @apiParam {string} name 歌单名
* @apiParam {string} url  歌曲链接
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func CreatePlayList(ctx iris.Context) {

	aul := new(validates.CreateUpdatePlaylistRequest)
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
	aul.UserID = userId
	playlist := models.NewPlaylistByStruct(aul)

	playlist.CreatePlaylist(aul)
	ctx.StatusCode(iris.StatusOK)
	if playlist.ID == 0 {
		_, _ = ctx.JSON(ApiResource(false, playlist, "操作失败"))
		return
	} else {
		_, _ = ctx.JSON(ApiResource(true, nil, "操作成功"))
		return
	}

}

/**
* @api {get} /admin/playlists 获取所有歌单
* @apiName 获取歌单
* @apiGroup Playlists
* @apiVersion 1.0.0
* @apiDescription 获取所有歌单
* @apiSampleRequest /admin/playlists
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission 登陆用户
 */
func GetAllPlayList(ctx iris.Context) {
	offset := ctx.URLParamIntDefault("offset", 1)
	limit := ctx.URLParamIntDefault("limit", 15)
	orderBy := ctx.URLParam("orderBy")

	Id := ctx.Values().Get("auth_user_id").(uint)
	playlists := models.GetAllPlaylist(Id, orderBy, offset, limit)

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(ApiResource(true, playlistsTransform(playlists), "操作成功"))

}

/**
* @api {get} /admin/playlists/:id 根据id获取歌单详情
* @apiName 根据id获取歌单详情
* @apiGroup Playlists
* @apiVersion 1.0.0
* @apiDescription 根据id获取歌曲信息
* @apiSampleRequest /admin/songs/:id
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission 登陆用户
 */
func GetPlaylistDetail(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")
	playlist := models.NewPlaylist(id, "")
	playlist.GetPlaylistById()

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(ApiResource(true, playlistTransform(playlist), "操作成功"))
}

/**
* @api {post} /admin/playlists/:id/update 更新账号
* @apiName 更新歌曲
* @apiGroup Playlists
* @apiVersion 1.0.0
* @apiDescription 更新歌曲
* @apiSampleRequest /admin/playlists/:id/update
* @apiParam {string} name 歌曲名
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func UpdatePlaylist(ctx iris.Context) {
	aul := new(validates.CreateUpdatePlaylistRequest)

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
	playlist := models.NewPlaylist(id, "")

	playlist.UpdatePlaylist(aul)
	ctx.StatusCode(iris.StatusOK)
	if playlist.ID == 0 {
		_, _ = ctx.JSON(ApiResource(false, playlist, "操作失败"))
		return
	} else {
		_, _ = ctx.JSON(ApiResource(true, nil, "操作成功"))
		return
	}

}

/**
* @api {delete} /admin/playlists/:id/delete 删除歌单
* @apiName 删除歌单
* @apiGroup Songs
* @apiVersion 1.0.0
* @apiDescription 删除歌曲
* @apiSampleRequest /admin/playlists/:id/delete
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func DeletePlaylist(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")

	playlist := models.NewPlaylist(id, "")
	playlist.DeletePlaylist()

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(ApiResource(true, nil, "删除成功"))
}

func playlistsTransform(playlists []*models.Playlist) []*transformer.Playlist {
	var ps []*transformer.Playlist
	for _, playlist := range playlists {
		p := playlistTransform(playlist)
		ps = append(ps, p)
	}
	return ps
}

func playlistTransform(playlist *models.Playlist) *transformer.Playlist {
	p := &transformer.Playlist{}
	g := gf.NewTransform(p, playlist, time.RFC3339)
	_ = g.Transformer()
	return p
}
