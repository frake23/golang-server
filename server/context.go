package server

import (
	"./services"
	"github.com/labstack/echo"
)

type ServerContext struct {
	echo.Context
	UserService *services.IUserService
}
