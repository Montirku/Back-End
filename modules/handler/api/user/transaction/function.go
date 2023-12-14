package transaction

import (
	"net/http"

	giu "github.com/fazaalexander/montirku-be/helper/getIdUser"
	mr "github.com/fazaalexander/montirku-be/modules/entity/midtrans"
	te "github.com/fazaalexander/montirku-be/modules/entity/transaction"
	"github.com/labstack/echo/v4"
)

func (th *TransactionHandler) CreateTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := giu.GetIdUser(c)

		transaction := te.Transaction{}
		c.Bind(&transaction)
		transaction.UserId = uint(id)

		snapUrl, transactionId, err := th.transactionUseCase.CreateTransaction(&transaction)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"Status":         201,
			"Message":        "Success Create Transaction",
			"Transaction_Id": transactionId,
			"Payment_url":    snapUrl,
		})
	}
}

func (th *TransactionHandler) MidtransNotification() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := mr.MidtransRequest{}
		c.Bind(request)

		err := th.transactionUseCase.MidtransNotification(&request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  400,
				"Message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":  200,
			"Message": "Payment successfully confirmed",
		})
	}
}
