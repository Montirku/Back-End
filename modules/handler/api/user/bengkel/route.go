package bengkel

import (
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (bh *BengkelHandler) RegisterRoutes(e *echo.Echo) {
	jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	bengkelGroup := e.Group("/user/bengkel")
	bengkelGroup.Use(jwtMiddleware)
	bengkelGroup.GET("", bh.GetAllBengkel())
	bengkelGroup.GET("/:id", bh.GetBengkelById())
	bengkelGroup.GET("/filter", bh.FilterBengkel())
}
