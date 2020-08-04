package routes

import (
	"os"

	"github.com/betacraft/yaag/irisyaag"
	"github.com/kataras/iris/v12"
	"musicplayer.pandaleo.cn/backend/config"
	"musicplayer.pandaleo.cn/backend/controllers"
	"musicplayer.pandaleo.cn/backend/middleware"
	"musicplayer.pandaleo.cn/backend/sysinit"
)

func App(api *iris.Application) {
	//api.Favicon("./static/favicons/favicon.ico")
	app := api.Party("/", middleware.CrsAuth()).AllowMethods(iris.MethodOptions)
	{
		staticPath := config.Root + "resources/app/static"
		if len(os.Getenv("GOPATH")) == 0 {
			staticPath = "resources/app/static"
		}
		app.HandleDir("/static", staticPath)
		app.Get("/", func(ctx iris.Context) { // 首页模块
			_ = ctx.View("app/index.html")
		})

		v1 := app.Party("/v1")
		{
			v1.Post("/admin/login", controllers.UserLogin)
			v1.Use(irisyaag.New())
			v1.PartyFunc("/admin", func(app iris.Party) {
				app.Get("/resetData", controllers.ResetData)
				casbinMiddleware := middleware.New(sysinit.Enforcer)               //casbin for gorm                                                   // <- IMPORTANT, register the middleware.
				app.Use(middleware.JwtHandler().Serve, casbinMiddleware.ServeHTTP) //登录验证
				app.Get("/logout", controllers.UserLogout).Name = "退出"

				app.PartyFunc("/users", func(users iris.Party) {
					users.Get("/", controllers.GetAllUsers).Name = "用户列表"
					users.Get("/{id:uint}", controllers.GetUser).Name = "用户详情"
					users.Post("/", controllers.CreateUser).Name = "创建用户"
					users.Put("/{id:uint}", controllers.UpdateUser).Name = "编辑用户"
					users.Delete("/{id:uint}", controllers.DeleteUser).Name = "删除用户"
					users.Get("/profile", controllers.GetProfile).Name = "个人信息"
				})
				app.PartyFunc("/roles", func(roles iris.Party) {
					roles.Get("/", controllers.GetAllRoles).Name = "角色列表"
					roles.Get("/{id:uint}", controllers.GetRole).Name = "角色详情"
					roles.Post("/", controllers.CreateRole).Name = "创建角色"
					roles.Put("/{id:uint}", controllers.UpdateRole).Name = "编辑角色"
					roles.Delete("/{id:uint}", controllers.DeleteRole).Name = "删除角色"
				})
				app.PartyFunc("/permissions", func(permissions iris.Party) {
					permissions.Get("/", controllers.GetAllPermissions).Name = "权限列表"
					permissions.Get("/{id:uint}", controllers.GetPermission).Name = "权限详情"
					permissions.Post("/import", controllers.ImportPermission).Name = "导入权限"
					permissions.Post("/", controllers.CreatePermission).Name = "创建权限"
					permissions.Put("/{id:uint}", controllers.UpdatePermission).Name = "编辑权限"
					permissions.Delete("/{id:uint}", controllers.DeletePermission).Name = "删除权限"
				})
				app.PartyFunc("/songs", func(songs iris.Party) {
					//songs.Get("/", controllers.GetAllSongs).Name = "歌曲列表"
					songs.Get("/", controllers.GetSongByUser).Name = "用户歌曲列表"
					songs.Get("/{id:uint}", controllers.GetSong).Name = "歌曲详情"
					songs.Post("/", controllers.CreateSong).Name = "新增歌曲"
					songs.Put("/{id:uint}", controllers.UpdateSong).Name = "编辑歌曲"
					songs.Delete("/{id:uint}", controllers.DeleteSong).Name = "删除歌曲"
				})
				app.PartyFunc("/playlists", func(playlists iris.Party) {
					playlists.Get("/", controllers.GetAllPlayList).Name = "歌单列表"
					playlists.Get("/{id:uint}", controllers.GetPlaylistDetail).Name = "歌单详情"
					playlists.Post("/", controllers.CreatePlayList).Name = "新增歌单"
					playlists.Put("/{id:uint}", controllers.UpdatePlaylist).Name = "编辑歌单"
					playlists.Delete("/{id:uint}", controllers.DeletePlaylist).Name = "删除歌单"
				})
				app.PartyFunc("/albums", func(albums iris.Party) {
					albums.Get("/", controllers.GetAllAlbums).Name = "专辑列表"
					albums.Get("/{id:uint}", controllers.GetAlbumDetail).Name = "专辑详情"
					albums.Post("/", controllers.CreateAlbum).Name = "新增专辑"
					albums.Put("/{id:uint}", controllers.UpdateAlbums).Name = "编辑专辑"
					albums.Delete("/{id:uint}", controllers.DeleteAlbums).Name = "删除专辑"
				})
				app.PartyFunc("/artists", func(artists iris.Party) {
					artists.Get("/", controllers.GetAllArtists).Name = "歌手列表"
					artists.Get("/{id:uint}", controllers.GetArtistDetail).Name = "歌手详情"
					artists.Post("/", controllers.CreateArtist).Name = "新增歌手"
					artists.Put("/{id:uint}", controllers.UpdateArtist).Name = "编辑歌手"
					artists.Delete("/{id:uint}", controllers.DeleteArtist).Name = "删除歌手"
				})
				app.PartyFunc("/lyrics", func(lyrics iris.Party) {
					lyrics.Get("/", controllers.GetAllLyrics).Name = "歌词列表"
					lyrics.Get("/{id:uint}", controllers.GetLyricDetail).Name = "歌词详情"
					lyrics.Post("/", controllers.CreateLyric).Name = "新增歌词"
					lyrics.Put("/{id:uint}", controllers.UpdateLyrics).Name = "编辑歌词"
					lyrics.Delete("/{id:uint}", controllers.DeleteLyrics).Name = "删除歌词"
				})
				app.PartyFunc("/qiniutoken", func(qiniu iris.Party) {
					qiniu.Get("/", controllers.GetQiniuToken).Name = "七牛token"
				})
			})
		}
	}

}
