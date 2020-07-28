package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"musicplayer.pandaleo.cn/datasource"
	"musicplayer.pandaleo.cn/repositories"
	"musicplayer.pandaleo.cn/service"
	"musicplayer.pandaleo.cn/web/controllers"
	"musicplayer.pandaleo.cn/web/middleware"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.RegisterView(iris.HTML("./web/views", ".html"))

	mvc.Configure(app.Party("/songs"), songs)

	app.Run(
		iris.Addr(":8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		)
}

func songs(app *mvc.Application)  {
	app.Router.Use(middleware.BasicAuth)

	repo := repositories.NewSongRepository(datasource.Songs)
	songService := service.NewSongService(repo)
	app.Register(songService)

	app.Handle(new(controllers.SongController))
}