package auth

import "github.com/labstack/echo/v4"

func (ah *AuthHandler) RegisterRoutes(e *echo.Echo) {
	userGroup := e.Group("/user")
	userGroup.POST("/register", ah.Register())
	userGroup.POST("/email-verify", ah.EmailVerification())
	userGroup.PUT("/email-otp-verify", ah.EmailOTPVerification())
}
