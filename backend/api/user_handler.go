package api

import (
	"github.com/kataras/iris/v12"
	"musicplayer/cache"
	"musicplayer/entity"
	"musicplayer/service"
	"musicplayer/sql"
	"time"
)

type UserHandler struct {
	service *service.UserService
	cache   *cache.Cache
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
		cache:   cache.New(service, "users", time.Minute),
	}
}

func (h *UserHandler) GetByID(ctx iris.Context) {
	id := ctx.Params().GetString("id")

	var user []byte
	err := h.cache.GetByID(ctx.Request().Context(), id, &user)
	if err != nil {
		if err == sql.ErrNoRows {
			writeEntityNotFound(ctx)
			return
		}

		debugf("UserHandler.GetByID(id=%v): %v", id, err)
		writeInternalServerError(ctx)
		return
	}

	ctx.ContentType("application/json")
	ctx.Write(user)

}

func (h *UserHandler) List(ctx iris.Context) {
	key := ctx.Request().URL.RawQuery

	users := []byte("[]")
	err := h.cache.List(ctx.Request().Context(), key, &users)
	if err != nil && err != sql.ErrNoRows {
		debugf("UserHandler.List(DB) (%s): %v",
			key, err)

		writeInternalServerError(ctx)
		return
	}

	ctx.ContentType("application/json")
	ctx.Write(users)
}

func (h *UserHandler) Create(ctx iris.Context) {
	var user entity.User
	if err := ctx.ReadJSON(&user); err != nil {
		return
	}

	id, err := h.service.Insert(ctx.Request().Context(), user)
	if err != nil {
		if err == sql.ErrUnprocessable {
			ctx.StopWithJSON(iris.StatusUnprocessableEntity, newError(iris.StatusUnprocessableEntity, ctx.Request().Method, ctx.Path(), "required fields are missing"))
			return
		}

		debugf("UserHandler.Create(DB): %v", err)
		writeInternalServerError(ctx)
		return
	}

	// Send 201 with body of {"id":$last_inserted_id"}.
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{user.PrimaryKey(): id})
}

func (h *UserHandler) Update(ctx iris.Context) {
	var user entity.User
	if err := ctx.ReadJSON(&user); err != nil {
		return
	}

	affected, err := h.service.Update(ctx.Request().Context(), user)
	if err != nil {
		if err == sql.ErrUnprocessable {
			ctx.StopWithJSON(iris.StatusUnprocessableEntity, newError(iris.StatusUnprocessableEntity, ctx.Request().Method, ctx.Path(), "required fields are missing"))
			return
		}

		debugf("UserHandler.Update(DB): %v", err)
		writeInternalServerError(ctx)
		return
	}

	status := iris.StatusOK
	if affected == 0 {
		status = iris.StatusNotModified
	}

	ctx.StatusCode(status)
}

func (h *UserHandler) PartialUpdate(ctx iris.Context) {
	id := ctx.Params().GetInt64Default("id", 0)

	var attrs map[string]interface{}
	if err := ctx.ReadJSON(&attrs); err != nil {
		return
	}

	affected, err := h.service.PartialUpdate(ctx.Request().Context(), id, attrs)
	if err != nil {
		if err == sql.ErrUnprocessable {
			ctx.StopWithJSON(iris.StatusUnprocessableEntity, newError(iris.StatusUnprocessableEntity, ctx.Request().Method, ctx.Path(), "unsupported value(s)"))
			return
		}

		debugf("ProductHandler.PartialUpdate(DB): %v", err)
		writeInternalServerError(ctx)
		return
	}

	status := iris.StatusOK
	if affected == 0 {
		status = iris.StatusNotModified
	}

	ctx.StatusCode(status)
}

func (h *UserHandler) Delete(ctx iris.Context) {
	id := ctx.Params().GetInt64Default("id", 0)

	affected, err := h.service.DeleteByID(ctx.Request().Context(), id)
	if err != nil {
		debugf("UserHandler.Delete(DB): %v", err)
		writeInternalServerError(ctx)
		return
	}

	status := iris.StatusOK // StatusNoContent
	if affected == 0 {
		status = iris.StatusNotModified
	}

	ctx.StatusCode(status)
}
