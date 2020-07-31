package api

import (
	"github.com/kataras/iris/v12"
	"musicplayer/entity"
	"musicplayer/service"
	"musicplayer/sql"
)

type RoleHandler struct {
	// [...options]

	service *service.RoleService
}

func NewRoleHandler(service *service.RoleService) *RoleHandler {
	return &RoleHandler{service}
}

func (h *RoleHandler) GetByID(ctx iris.Context) {
	id := ctx.Params().GetInt64Default("id", 0)

	var rol entity.Role
	err := h.service.GetByID(ctx.Request().Context(), &rol, id)
	if err != nil {
		if err == sql.ErrNoRows {
			writeEntityNotFound(ctx)
			return
		}

		debugf("RoleHandler.GetByID(id=%d): %v", id, err)
		writeInternalServerError(ctx)
		return
	}

	ctx.JSON(rol)
}

func (h *RoleHandler) List(ctx iris.Context) {
	q := ctx.Request().URL.Query()
	opts := sql.ParseListOptions(q)

	// initialize here in order to return an empty json array `[]` instead of `null`.
	roles := entity.Roles{}
	err := h.service.List(ctx.Request().Context(), &roles, opts)
	if err != nil && err != sql.ErrNoRows {
		debugf("RoleHandler.List(DB) (limit=%d offset=%d where=%s=%v): %v",
			opts.Limit, opts.Offset, opts.WhereColumn, opts.WhereValue, err)

		writeInternalServerError(ctx)
		return
	}

	ctx.JSON(roles)
}

func (h *RoleHandler) Create(ctx iris.Context) {
	var rol entity.Role
	if err := ctx.ReadJSON(&rol); err != nil {
		return
	}
	id, err := h.service.Insert(ctx.Request().Context(), rol)
	if err != nil {
		if err == sql.ErrUnprocessable {
			ctx.StopWithJSON(iris.StatusUnprocessableEntity, newError(iris.StatusUnprocessableEntity, ctx.Request().Method, ctx.Path(), "required fields are missing"))
			return
		}

		debugf("RoleHandler.Create(DB): %v", err)
		writeInternalServerError(ctx)
		return
	}

	// Send 201 with body of {"id":$last_inserted_id"}.
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{rol.PrimaryKey(): id})
}

func (h *RoleHandler) Update(ctx iris.Context) {
	var rol entity.Role
	if err := ctx.ReadJSON(&rol); err != nil {
		return
	}

	affected, err := h.service.Update(ctx.Request().Context(), rol)
	if err != nil {
		if err == sql.ErrUnprocessable {
			ctx.StopWithJSON(iris.StatusUnprocessableEntity, newError(iris.StatusUnprocessableEntity, ctx.Request().Method, ctx.Path(), "required fields are missing"))
			return
		}

		debugf("RoleHandler.Update(DB): %v", err)
		writeInternalServerError(ctx)
		return
	}

	status := iris.StatusOK
	if affected == 0 {
		status = iris.StatusNotModified
	}

	ctx.StatusCode(status)
}

func (h *RoleHandler) PartialUpdate(ctx iris.Context) {
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

		debugf("RoleHandler.PartialUpdate(DB): %v", err)
		writeInternalServerError(ctx)
		return
	}

	status := iris.StatusOK
	if affected == 0 {
		status = iris.StatusNotModified
	}

	ctx.StatusCode(status)
}

func (h *RoleHandler) Delete(ctx iris.Context) {
	id := ctx.Params().GetInt64Default("id", 0)

	affected, err := h.service.DeleteByID(ctx.Request().Context(), id)
	if err != nil {
		debugf("RoleHandler.Delete(DB): %v", err)
		writeInternalServerError(ctx)
		return
	}

	status := iris.StatusOK // StatusNoContent
	if affected == 0 {
		status = iris.StatusNotModified
	}

	ctx.StatusCode(status)
}

func (h *RoleHandler) ListUsers(ctx iris.Context) {
	id := ctx.Params().GetInt64Default("id", 0)

	// NOTE: could add cache here too.

	q := ctx.Request().URL.Query()
	opts := sql.ParseListOptions(q).Where("role_id", id)
	opts.Table = "users"
	if opts.OrderByColumn == "" {
		opts.OrderByColumn = "updated_at"
	}

	var users entity.Users
	err := h.service.List(ctx.Request().Context(), &users, opts)
	if err != nil {
		debugf("RoleHandler.ListUsers(DB) (table=%s where=%s=%v limit=%d offset=%d): %v",
			opts.Table, opts.WhereColumn, opts.WhereValue, opts.Limit, opts.Offset, err)

		writeInternalServerError(ctx)
		return
	}

	ctx.JSON(users)
}

func (h *RoleHandler) InsertUsers(userService *service.UserService) iris.Handler {
	return func(ctx iris.Context) {
		roleID := ctx.Params().GetInt64Default("id", 0)

		var users []entity.User
		if err := ctx.ReadJSON(&users); err != nil {
			return
		}

		for i := range users {
			users[i].RoleID = roleID
		}

		inserted, err := userService.BatchInsert(ctx.Request().Context(), users)
		if err != nil {
			if err == sql.ErrUnprocessable {
				ctx.StopWithJSON(iris.StatusUnprocessableEntity, newError(iris.StatusUnprocessableEntity, ctx.Request().Method, ctx.Path(), "required fields are missing"))
				return
			}

			debugf("RoleHandler.InsertProducts(DB): %v", err)
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
