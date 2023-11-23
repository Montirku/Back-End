package auth

import (
	"net/http"

	ue "github.com/fazaalexander/montirku-be/modules/entity/user"

	"github.com/labstack/echo/v4"
)

func (ah *AuthHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request ue.RegisterRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"Status":  http.StatusBadRequest,
				"Message": err.Error(),
			})
		}

		err := ah.authUsecase.Register(&request)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Status":  http.StatusInternalServerError,
				"Message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":  http.StatusOK,
			"Message": "Registration Successful",
		})
	}
}

func (ah *AuthHandler) EmailVerification() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request ue.VerifyEmailRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"Status":  http.StatusBadRequest,
				"Message": err.Error(),
			})
		}

		email, err := ah.authUsecase.EmailVerification(&request)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Status":  http.StatusInternalServerError,
				"Message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":  http.StatusOK,
			"Message": "OTP Code has been successfully sent to your email.",
			"Email":   email,
		})
	}
}

func (ah *AuthHandler) EmailOTPVerification() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request ue.VerifOtp
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"Status":  http.StatusBadRequest,
				"Message": err.Error(),
			})
		}

		err := ah.authUsecase.EmailOTPVerification(request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  http.StatusBadRequest,
				"Message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":  http.StatusOK,
			"Message": "Email verified.",
		})
	}
}
