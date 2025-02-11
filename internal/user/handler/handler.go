package handler

import (
	"paypath/internal/user/model"
	"paypath/internal/user/processor"

	"github.com/labstack/echo/v4"
)

type User interface {
	Create(ctx echo.Context) error
	Confirm(ctx echo.Context) error
	Login(ctx echo.Context) error
}

func NewUser(processor processor.User) User {
	return &handler{processor: processor}
}

type handler struct {
	processor processor.User
}

func (u *handler) Login(ctx echo.Context) error {
	userBody := model.User{}
	ctx.Bind(&userBody)
	token, err := u.processor.Login(ctx.Request().Context(), &userBody)
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]string{"token": token})

}

func (u *handler) Create(ctx echo.Context) error {
	userBody := model.User{}
	ctx.Bind(&userBody)
	err := u.processor.CreateUser(&userBody)
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]string{"message": "user created"})
}

func (u *handler) Confirm(ctx echo.Context) error {
	userID := ctx.QueryParam("code")
	token, err := u.processor.ConfirmUser(userID)
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]string{"token": token})
}

// func (u *handler) delete(ctx echo.Context) error {

// }
