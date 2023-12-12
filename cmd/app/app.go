package app

import (
	"github.com/fazaalexander/montirku-be/cmd/routes"
	"github.com/fazaalexander/montirku-be/common"

	"github.com/fazaalexander/montirku-be/database/mysql"
	authHandler "github.com/fazaalexander/montirku-be/modules/handler/api/auth"
	bengkelHandler "github.com/fazaalexander/montirku-be/modules/handler/api/user/bengkel"
	authRepo "github.com/fazaalexander/montirku-be/modules/repository/auth"
	bengkelRepo "github.com/fazaalexander/montirku-be/modules/repository/user/bengkel"
	authUsecase "github.com/fazaalexander/montirku-be/modules/usecase/auth"
	bengkelUseCase "github.com/fazaalexander/montirku-be/modules/usecase/user/bengkel"

	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	mysql.Init()

	authRepo := authRepo.New(mysql.DB)
	authUsecase := authUsecase.New(authRepo)
	authHandler := authHandler.New(authUsecase)
	bengkelRepo := bengkelRepo.New(mysql.DB)
	bengkelUsecase := bengkelUseCase.New(bengkelRepo)
	bengkelHandler := bengkelHandler.New(bengkelUsecase)

	handler := common.Handler{
		AuthHandler:    authHandler,
		BengkelHandler: bengkelHandler,
	}

	router := routes.StartRoute(handler)

	return router
}
