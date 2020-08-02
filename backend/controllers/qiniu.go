package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

/**
* @api {get} /admin/users/profile 获取登陆用户信息
* @apiName 获取登陆用户信息
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 获取登陆用户信息
* @apiSampleRequest /admin/users/profile
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission 登陆用户
 */
func GetQiniuToken(ctx iris.Context) {

	accessKey := "nGhnoTK0Yabi98wjmrfdRnChtATkptKg-kZ86YoV"
	secretKey := "irqfDl7KhETwxz2KLNvrJY1BWLsZSfInPdFCqmWD"

	bucket:="files"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	_, _ = ctx.JSON(ApiResource(true, upToken, ""))
}