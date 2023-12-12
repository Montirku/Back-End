package bengkel

import (
	"github.com/labstack/echo/v4"
)

func (bh *BengkelHandler) RegisterRoutes(e *echo.Echo) {
	bengkelGroup := e.Group("/user/bengkel")
	bengkelGroup.GET("", bh.GetAllBengkel())
	bengkelGroup.GET("/:id", bh.GetBengkelById())
	bengkelGroup.GET("/filter", bh.FilterBengkel())
}
