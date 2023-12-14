package midtrans

import (
	"os"

	te "github.com/fazaalexander/montirku-be/modules/entity/transaction"
	ue "github.com/fazaalexander/montirku-be/modules/entity/user"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func CreateMidtransUrl(transaction *te.Transaction, user *ue.UserResponse) (string, error) {
	var itemDetails []midtrans.ItemDetails
	for _, val := range transaction.TransactionDetails {
		item := midtrans.ItemDetails{
			ID:           val.ServiceId,
			Name:         val.ServiceName,
			Qty:          int32(val.Qty),
			Price:        int64(val.SubTotalPrice) / int64(val.Qty),
			MerchantName: val.BengkelName,
		}
		itemDetails = append(itemDetails, item)
	}

	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	var s = snap.Client{}
	s.New(serverKey, midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.TransactionId,
			GrossAmt: int64(transaction.TotalPrice),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.FirstName,
			LName: user.LastName,
			Email: user.Email,
			Phone: user.Phone,
		},
		Items: &itemDetails,
	}

	snapResp, midtransErr := s.CreateTransactionUrl(req)
	if midtransErr != nil {
		return "", echo.NewHTTPError(500, "Gagal membuat transaksi")
	}

	return snapResp, nil
}
