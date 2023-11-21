package app

import (
	"github.com/fazaalexander/montirku-be/cmd/routes"
	"github.com/fazaalexander/montirku-be/common"

	"github.com/fazaalexander/montirku-be/database/mysql"

	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	mysql.Init()

	router := routes.StartRoute(common.Handler{})
	return router
}
