package api

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/kataras/iris/v12/middleware/requestid"
	"musicplayer/service"
	"musicplayer/sql"
	"time"
)

func Router(db sql.Database, secret string) func(iris.Party) {
	return func(r iris.Party) {
		j := jwt.HMAC(1500*time.Minute, secret)

		r.Use(requestid.New())
		r.Use(verifyToken(j))

		r.Get("/token", writeToken(j))

		var (
			sheetService = service.NewSheetService(db)
			songService = service.NewSongService(db)
			userService = service.NewUserService(db)
			roleService = service.NewRoleService(db)
		)

		sht := r.Party("/api/sheet")
		{

			handler := NewSheetHandler(sheetService)

			sht.Get("/", handler.List)
			sht.Post("/", handler.Create)
			sht.Put("/", handler.Update)

			sht.Get("/{id:int64}", handler.GetByID)
			sht.Patch("/{id:int64}", handler.PartialUpdate)
			sht.Delete("/{id:int64}", handler.Delete)
			sht.Get("/{id:int64}/songs", handler.ListSongs)
			sht.Post("/{id:int64}/songs", handler.InsertSongs(songService))
		}

		sng := r.Party("/api/song")
		{
			handler := NewSongHandler(songService)

			sng.Get("/", handler.List)
			sng.Post("/", handler.Create)
			sng.Put("/", handler.Update)

			sng.Get("/{id:int64}", handler.GetByID)
			sng.Patch("/{id:int64}", handler.PartialUpdate)
			sng.Delete("/{id:int64}", handler.Delete)
		}

		usr := r.Party("/api/user")
		{
			handler := NewUserHandler(userService)

			usr.Get("/", handler.List)
			usr.Post("/", handler.Create)
			usr.Put("/", handler.Update)

			usr.Get("/{id:int64}", handler.GetByID)
			usr.Patch("/{id:int64}", handler.PartialUpdate)
			usr.Delete("/{id:int64}", handler.Delete)
		}

		rol := r.Party("/api/role")
		{
			handler := NewRoleHandler(roleService)

			rol.Get("/", handler.List)
			rol.Post("/", handler.Create)
			rol.Put("/", handler.Update)

			rol.Get("/{id:int64}", handler.GetByID)
			rol.Patch("/{id:int64}", handler.PartialUpdate)
			rol.Delete("/{id:int64}", handler.Delete)
		}

	}
}

func writeToken(j *jwt.JWT) iris.Handler {
	return func(ctx iris.Context) {
		claims := jwt.Claims{
			//Issuer:   "https://iris-go.com",
			Issuer:   "http://localhost:8080",
			Audience: jwt.Audience{requestid.Get(ctx)},
		}

		j.WriteToken(ctx, claims)
	}
}


func verifyToken(j *jwt.JWT) iris.Handler {
	return func(ctx iris.Context) {
		// Allow all GET.
		if ctx.Method() == iris.MethodGet {
			ctx.Next()
			return
		}

		j.Verify(ctx)
	}
}
