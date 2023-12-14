package app

import (
	"github.com/fazaalexander/montirku-be/cmd/routes"
	"github.com/fazaalexander/montirku-be/common"

	"github.com/fazaalexander/montirku-be/database/mysql"
	authHandler "github.com/fazaalexander/montirku-be/modules/handler/api/auth"
	bengkelHandler "github.com/fazaalexander/montirku-be/modules/handler/api/user/bengkel"
	transactionHandler "github.com/fazaalexander/montirku-be/modules/handler/api/user/transaction"
	authRepo "github.com/fazaalexander/montirku-be/modules/repository/auth"
	bengkelRepo "github.com/fazaalexander/montirku-be/modules/repository/user/bengkel"
	transactionRepo "github.com/fazaalexander/montirku-be/modules/repository/user/transaction"
	authUsecase "github.com/fazaalexander/montirku-be/modules/usecase/auth"
	bengkelUseCase "github.com/fazaalexander/montirku-be/modules/usecase/user/bengkel"
	transactionUseCase "github.com/fazaalexander/montirku-be/modules/usecase/user/transaction"

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
	transactionRepo := transactionRepo.New(mysql.DB)
	transactionUseCase := transactionUseCase.New(transactionRepo)
	transactionHandler := transactionHandler.New(transactionUseCase)

	handler := common.Handler{
		AuthHandler:        authHandler,
		BengkelHandler:     bengkelHandler,
		TransactionHandler: transactionHandler,
	}

	router := routes.StartRoute(handler)

	return router
}
