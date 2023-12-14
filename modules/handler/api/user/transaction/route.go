package transaction

import (
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (th *TransactionHandler) RegisterRoutes(e *echo.Echo) {
	jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	transactionGroup := e.Group("/user/transaction")
	transactionGroup.POST("", th.CreateTransaction(), jwtMiddleware)
	transactionGroup.POST("/midtrans/notifications", th.MidtransNotification(), jwtMiddleware)
}
