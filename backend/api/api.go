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
		j := jwt.HMAC(15*time.Minute, secret)

		r.Use(requestid.New())
		r.Use(verifyToken(j))

		r.Get("/token", writeToken(j))

		var (
			categoryService = service.NewCategoryService(db)
			productService  = service.NewProductService(db)
		)

		cat := r.Party("/category")
		{
			// TODO: new Use to add middlewares to specific
			// routes per METHOD ( we already have the per path through parties.)
			handler := NewCategoryHandler(categoryService)

			cat.Get("/test", handler.Test)

			cat.Get("/", handler.List)
			cat.Post("/", handler.Create)
			cat.Put("/", handler.Update)

			cat.Get("/{id:int64}", handler.GetByID)
			cat.Patch("/{id:int64}", handler.PartialUpdate)
			cat.Delete("/{id:int64}", handler.Delete)
			cat.Get("/{id:int64}/products", handler.ListProducts)
			cat.Post("/{id:int64}/products", handler.InsertProducts(productService))
		}

		prod := r.Party("/product")
		{
			handler := NewProductHandler(productService)

			prod.Get("/", handler.List)
			prod.Post("/", handler.Create)
			prod.Put("/", handler.Update)

			prod.Get("/{id:int64}", handler.GetByID)
			prod.Patch("/{id:int64}", handler.PartialUpdate)
			prod.Delete("/{id:int64}", handler.Delete)
		}

	}
}

func writeToken(j *jwt.JWT) iris.Handler {
	return func(ctx iris.Context) {
		claims := jwt.Claims{
			Issuer:   "https://iris-go.com",
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