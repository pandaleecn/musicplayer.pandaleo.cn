package api

import (
	"github.com/kataras/iris/v12"
	"musicplayer/entity"
	"musicplayer/service"
	"musicplayer/sql"
)

type SheetHandler struct {
	// [...options]

	service *service.SheetService
}

func NewSheetHandler(service *service.SheetService) *SheetHandler {
	return &SheetHandler{service}
}

func (h *SheetHandler) GetByID(ctx iris.Context) {
	id := ctx.Params().GetInt64Default("id", 0)

	var sht entity.Sheet
	err := h.service.GetByID(ctx.Request().Context(), &sht, id)
	if err != nil {
		if err == sql.ErrNoRows {
			writeEntityNotFound(ctx)
			return
		}

		debugf("SheetHandler.GetByID(id=%d): %v", id, err)
		writeInternalServerError(ctx)
		return
	}

	ctx.JSON(sht)
}

func (h *SheetHandler) List(ctx iris.Context) {
	q := ctx.Request().URL.Query()
	opts := sql.ParseListOptions(q)

	// initialize here in order to return an empty json array `[]` instead of `null`.
	sheets := entity.Sheets{}
	err := h.service.List(ctx.Request().Context(), &sheets, opts)
	if err != nil && err != sql.ErrNoRows {
		debugf("SheetHandler.List(DB) (limit=%d offset=%d where=%s=%v): %v %s",
			opts.Limit, opts.Offset, opts.WhereColumn, opts.WhereValue, err)

		writeInternalServerError(ctx)
		return
	}

	ctx.JSON(sheets)
}

func (h *SheetHandler) Create(ctx iris.Context) {
	var cat entity.Sheet
	if err := ctx.ReadJSON(&cat); err != nil {
		return
	}

	id, err := h.service.Insert(ctx.Request().Context(), cat)
	if err != nil {
		if err == sql.ErrUnprocessable {
			ctx.StopWithJSON(iris.StatusUnprocessableEntity, newError(iris.StatusUnprocessableEntity, ctx.Request().Method, ctx.Path(), "required fields are missing"))
			return
		}

		debugf("SheetHandler.Create(DB): %v", err)
		writeInternalServerError(ctx)
		return
	}

	// Send 201 with body of {"id":$last_inserted_id"}.
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{cat.PrimaryKey(): id})
}

func (h *SheetHandler) Update(ctx iris.Context) {
	var cat entity.Sheet
	if err := ctx.ReadJSON(&cat); err != nil {
		return
	}

	affected, err := h.service.Update(ctx.Request().Context(), cat)
	if err != nil {
		if err == sql.ErrUnprocessable {
			ctx.StopWithJSON(iris.StatusUnprocessableEntity, newError(iris.StatusUnprocessableEntity, ctx.Request().Method, ctx.Path(), "required fields are missing"))
			return
		}

		debugf("SheetHandler.Update(DB): %v", err)
		writeInternalServerError(ctx)
		return
	}

	status := iris.StatusOK
	if affected == 0 {
		status = iris.StatusNotModified
	}

	ctx.StatusCode(status)
}

func (h *SheetHandler) PartialUpdate(ctx iris.Context) {
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

		debugf("SheetHandler.PartialUpdate(DB): %v", err)
		writeInternalServerError(ctx)
		return
	}

	status := iris.StatusOK
	if affected == 0 {
		status = iris.StatusNotModified
	}

	ctx.StatusCode(status)
}

func (h *SheetHandler) Delete(ctx iris.Context) {
	id := ctx.Params().GetInt64Default("id", 0)

	affected, err := h.service.DeleteByID(ctx.Request().Context(), id)
	if err != nil {
		debugf("SheetHandler.Delete(DB): %v", err)
		writeInternalServerError(ctx)
		return
	}

	status := iris.StatusOK // StatusNoContent
	if affected == 0 {
		status = iris.StatusNotModified
	}

	ctx.StatusCode(status)
}



func (h *SheetHandler) ListSongs(ctx iris.Context) {
	id := ctx.Params().GetInt64Default("id", 0)

	// NOTE: could add cache here too.

	q := ctx.Request().URL.Query()
	opts := sql.ParseListOptions(q).Where("sheet_id", id)
	opts.Table = "songs"
	if opts.OrderByColumn == "" {
		opts.OrderByColumn = "updated_at"
	}

	var songs entity.Songs
	err := h.service.List(ctx.Request().Context(), &songs, opts)
	if err != nil {
		debugf("SheetHandler.ListSongs(DB) (table=%s where=%s=%v limit=%d offset=%d): %v",
			opts.Table, opts.WhereColumn, opts.WhereValue, opts.Limit, opts.Offset, err)

		writeInternalServerError(ctx)
		return
	}

	ctx.JSON(songs)
}

func (h *SheetHandler) InsertSongs(songService *service.SongService) iris.Handler {
	return func(ctx iris.Context) {
		sheetID := ctx.Params().GetInt64Default("id", 0)

		var songs []entity.Song
		if err := ctx.ReadJSON(&songs); err != nil {
			return
		}

		for i := range songs {
			songs[i].ID = sheetID
		}

		inserted, err := songService.BatchInsert(ctx.Request().Context(), songs)
		if err != nil {
			if err == sql.ErrUnprocessable {
				ctx.StopWithJSON(iris.StatusUnprocessableEntity, newError(iris.StatusUnprocessableEntity, ctx.Request().Method, ctx.Path(), "required fields are missing"))
				return
			}

			debugf("SheetHandler.InsertProducts(DB): %v", err)
			writeInternalServerError(ctx)
			return
		}

		if inserted == 0 {
			ctx.StatusCode(iris.StatusNotModified)
			return
		}

		// Send 201 with body of {"inserted":$inserted"}.
		ctx.StatusCode(iris.StatusCreated)
		ctx.JSON(iris.Map{"inserted": inserted})
	}
}
