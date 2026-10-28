package controller

import "github.com/labstack/echo/v4"

type IUserController interface {

	SignUp(C echo.Context) error
	LogIn(C echo.Context) error
	LogOut(C echo.Context) error
}