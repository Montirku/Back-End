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
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
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

		return c.JSON(http.StatusOK, echo.Map{
			"Status":  http.StatusOK,
			"Message": "Registration Successful",
		})
	}
}

func (ah *AuthHandler) UserLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request ue.LoginRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
				"Status":  http.StatusBadRequest,
				"Message": err.Error(),
			})
		}

		data, role, err := ah.authUsecase.Login(&request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  http.StatusBadRequest,
				"Message": err.Error(),
			})
		}

		if role != 2 {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"Status":  http.StatusUnauthorized,
				"Message": "Invalid email or password.",
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"Status":  http.StatusOK,
			"Message": "Login successful.",
			"Data":    data,
		})
	}
}

func (ah *AuthHandler) EmailVerification() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request ue.VerifyEmailRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
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

		return c.JSON(http.StatusOK, echo.Map{
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
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
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

		return c.JSON(http.StatusOK, echo.Map{
			"Status":  http.StatusOK,
			"Message": "Email verified.",
		})
	}
}

func (ah *AuthHandler) ForgotPassword() echo.HandlerFunc {
	return func(c echo.Context) error {

		var request ue.ForgotPasswordRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
				"Status":  http.StatusBadRequest,
				"Message": err.Error(),
			})
		}

		email, err := ah.authUsecase.ForgotPassword(&request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  http.StatusBadRequest,
				"Message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"Status":  http.StatusOK,
			"Message": "OTP Code has been successfully sent to your email.",
			"Email":   email,
		})
	}
}
func (ah *AuthHandler) PasswordOTPVerification() echo.HandlerFunc {
	return func(c echo.Context) error {

		var request ue.VerifOtp
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
				"Status":  http.StatusBadRequest,
				"Message": err.Error(),
			})
		}

		err := ah.authUsecase.PasswordOTPVerification(request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  http.StatusBadRequest,
				"Message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"Status":  http.StatusOK,
			"Message": "Email verified",
		})
	}
}
func (ah *AuthHandler) ChangePassword() echo.HandlerFunc {
	return func(c echo.Context) error {

		var request ue.RecoveryRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
				"Status":  http.StatusBadRequest,
				"Message": err.Error(),
			})
		}

		err := ah.authUsecase.ChangePassword(request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  http.StatusBadRequest,
				"Message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"Status":  http.StatusOK,
			"Message": "Successfully change password.",
		})
	}
}
