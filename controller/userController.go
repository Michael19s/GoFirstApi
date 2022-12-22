package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Michael19s/go_first_api/ent"
	"github.com/Michael19s/go_first_api/service"
	"github.com/labstack/echo/v4"
)

func UserGetAll(ctx echo.Context) error {
	users, err := service.NewUserService(ctx.Request().Context()).UserGetAll()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"Error Message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, users)
}

func UserGetById(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"Error Message": err.Error()})
	}

	user, err := service.NewUserService(ctx.Request().Context()).UserGetById(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"Error Message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, user)
}

func UserCreate(ctx echo.Context) error {
	var newUser ent.User
	err := json.NewDecoder(ctx.Request().Body).Decode(&newUser)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"Error Message": err.Error()})
	}

	user, err := service.NewUserService(ctx.Request().Context()).UserCreate(newUser)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"Error Message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, user)
}

func UserUpdate(ctx echo.Context) error {
	var newUserData ent.User
	err := json.NewDecoder(ctx.Request().Body).Decode(&newUserData)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"Error Message": err.Error()})
	}

	user, err := service.NewUserService(ctx.Request().Context()).UserUpdate(newUserData)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"Error Message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, user)
}

func UserDelete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]string{"Error Message": err.Error()})
	}

	user, err := service.NewUserService(ctx.Request().Context()).UserDelete(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"Error Message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, user)
}
