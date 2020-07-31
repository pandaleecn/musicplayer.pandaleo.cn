package api

import (
	"github.com/kataras/iris/v12"
	"musicplayer/cache"
	"musicplayer/entity"
	"musicplayer/service"
	"musicplayer/sql"
	"time"
)

type SongHandler struct {
	service *service.SongService
	cache   *cache.Cache
}

func NewSongHandler(service *service.SongService) *SongHandler {
	return &SongHandler{
		service: service,
		cache:   cache.New(service, "songs", time.Minute),
	}
}

func (h *SongHandler) GetByID(ctx iris.Context) {
	id := ctx.Params().GetString("id")

	var song []byte
	err := h.cache.GetByID(ctx.Request().Context(), id, &song)
	if err != nil {
		if err == sql.ErrNoRows {
			writeEntityNotFound(ctx)
			return
		}

		debugf("SongHandler.GetByID(id=%v): %v", id, err)
		writeInternalServerError(ctx)
		return
	}

	ctx.ContentType("application/json")
	ctx.Write(song)

}

func (h *SongHandler) List(ctx iris.Context) {

	key := ctx.Request().URL.RawQuery

	songs := []byte("[]")
	err := h.cache.List(ctx.Request().Context(), key, &songs)
	if err != nil && err != sql.ErrNoRows {
		debugf("SongHandler.List(DB) (%s): %v",
			key, err)

		writeInternalServerError(ctx)
		return
	}

	ctx.ContentType("application/json")
	ctx.Write(songs)
}

func (h *SongHandler) Create(ctx iris.Context) {
	var song entity.Song
	if err := ctx.ReadJSON(&song); err != nil {
		return
	}

	id, err := h.service.Insert(ctx.Request().Context(), song)
	if err != nil {
		if err == sql.ErrUnprocessable {
			ctx.StopWithJSON(iris.StatusUnprocessableEntity, newError(iris.StatusUnprocessableEntity, ctx.Request().Method, ctx.Path(), "required fields are missing"))
			return
		}

		debugf("SongHandler.Create(DB): %v", err)
		writeInternalServerError(ctx)
		return
	}

	// Send 201 with body of {"id":$last_inserted_id"}.
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{song.PrimaryKey(): id})
}

func (h *SongHandler) Update(ctx iris.Context) {
	var song entity.Song
	if err := ctx.ReadJSON(&song); err != nil {
		return
	}

	affected, err := h.service.Update(ctx.Request().Context(), song)
	if err != nil {
		if err == sql.ErrUnprocessable {
			ctx.StopWithJSON(iris.StatusUnprocessableEntity, newError(iris.StatusUnprocessableEntity, ctx.Request().Method, ctx.Path(), "required fields are missing"))
			return
		}

		debugf("SongHandler.Update(DB): %v", err)
		writeInternalServerError(ctx)
		return
	}

	status := iris.StatusOK
	if affected == 0 {
		status = iris.StatusNotModified
	}

	ctx.StatusCode(status)
}

func (h *SongHandler) PartialUpdate(ctx iris.Context) {
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

		debugf("SongHandler.PartialUpdate(DB): %v", err)
		writeInternalServerError(ctx)
		return
	}

	status := iris.StatusOK
	if affected == 0 {
		status = iris.StatusNotModified
	}

	ctx.StatusCode(status)
}

func (h *SongHandler) Delete(ctx iris.Context) {
	id := ctx.Params().GetInt64Default("id", 0)

	affected, err := h.service.DeleteByID(ctx.Request().Context(), id)
	if err != nil {
		debugf("SongHandler.Delete(DB): %v", err)
		writeInternalServerError(ctx)
		return
	}

	status := iris.StatusOK // StatusNoContent
	if affected == 0 {
		status = iris.StatusNotModified
	}

	ctx.StatusCode(status)
}
