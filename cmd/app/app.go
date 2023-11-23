package app

import (
	"github.com/fazaalexander/montirku-be/cmd/routes"
	"github.com/fazaalexander/montirku-be/common"

	"github.com/fazaalexander/montirku-be/database/mysql"
	authHandler "github.com/fazaalexander/montirku-be/modules/handler/api/auth"
	authRepo "github.com/fazaalexander/montirku-be/modules/repository/auth"
	authUsecase "github.com/fazaalexander/montirku-be/modules/usecase/auth"

	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	mysql.Init()

	authRepo := authRepo.New(mysql.DB)
	authUsecase := authUsecase.New(authRepo)
	authHandler := authHandler.New(authUsecase)

	handler := common.Handler{
		AuthHandler: authHandler,
	}

	router := routes.StartRoute(handler)

	return router
}
