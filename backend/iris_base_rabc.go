package main

import (
	"fmt"
	"musicplayer.pandaleo.cn/backend/libs"
	"os"
	"time"

	"github.com/betacraft/yaag/yaag"
	"github.com/fatih/color"
	"github.com/kataras/iris/v12"
	"musicplayer.pandaleo.cn/backend/config"
	"musicplayer.pandaleo.cn/backend/models"
	"musicplayer.pandaleo.cn/backend/routes"
	"musicplayer.pandaleo.cn/backend/sysinit"
)

func NewLogFile() *os.File {
	path := "./logs/"
	_ = libs.CreateFile(path)
	filename := path + time.Now().Format("2006-01-02") + ".log"
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		color.Red(fmt.Sprintf("日志记录出错: %v", err))
	}

	return f
}

func NewApp() *iris.Application {
	api := iris.New()
	api.Logger().SetLevel("debug")

	htmlPath := config.Root + "resources"
	if len(os.Getenv("GOPATH")) == 0 {
		htmlPath = "resources"
	}
	api.RegisterView(iris.HTML(htmlPath, ".html"))

	db := sysinit.Db
	db.AutoMigrate(
		&models.User{},
		&models.OauthToken{},
		&models.Role{},
		&models.Permission{},
		&models.Stream{},
	)

	iris.RegisterOnInterrupt(func() {
		_ = db.Close()
	})

	docPath := config.Root + "resources/apiDoc/index.html"
	if len(os.Getenv("GOPATH")) == 0 {
		docPath = "resources/apiDoc/index.html"
	}
	yaag.Init(&yaag.Config{ // <- IMPORTANT, init the middleware. //api 文档配置
		On:       true,
		DocTitle: "irisadminapi",
		DocPath:  docPath, //设置绝对路径
		BaseUrls: map[string]string{
			"Production": config.Config.Host,
			"Staging":    "",
		},
	})

	routes.App(api) //注册 app 路由

	return api
}
