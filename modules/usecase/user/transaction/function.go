package transaction

import (
	"errors"
	"strconv"
	"time"

	"github.com/fazaalexander/montirku-be/helper/hash"
	midtransHelper "github.com/fazaalexander/montirku-be/helper/midtrans"
	vld "github.com/fazaalexander/montirku-be/helper/validator"
	mr "github.com/fazaalexander/montirku-be/modules/entity/midtrans"
	te "github.com/fazaalexander/montirku-be/modules/entity/transaction"
	ue "github.com/fazaalexander/montirku-be/modules/entity/user"
)

func (tc *transactionUseCase) CreateTransaction(transaction *te.Transaction) (string, string, error) {
	user, err := tc.transactionRepo.GetUserById(transaction.UserId)
	if err != nil {
		return "", "", nil
	}

	userResponse := ue.UserResponse{
		FirstName: user.UserDetail.FirstName,
		LastName:  user.UserDetail.LastName,
		Email:     user.Email,
		Phone:     user.UserDetail.Phone,
	}

	if user.RoleId != 2 {
		return "", "", errors.New("admin dan mitra tidak boleh melakukan transaksi")
	}

	var totalPrice float64
	for _, price := range transaction.TransactionDetails {
		totalPrice += price.SubTotalPrice
	}

	transactionId := "mon" + strconv.FormatUint(uint64(transaction.UserId), 10) + time.Now().UTC().Format("2006010215040105")
	transaction.TransactionId = transactionId
	transaction.StatusTransaction = "Belum Bayar"
	transaction.TotalPrice = totalPrice

	redirectUrl, err := midtransHelper.CreateMidtransUrl(transaction, &userResponse)
	if err != nil {
		return "", "", err
	}
	transaction.PaymentUrl = redirectUrl

	if err := vld.Validation(transaction); err != nil {
		return "", "", err
	}

	if err := tc.transactionRepo.CreateTransaction(transaction); err != nil {
		return "", "", err
	}

	return redirectUrl, transactionId, nil
}

func (tc *transactionUseCase) MidtransNotification(request *mr.MidtransRequest) error {

	if request != nil {
		signatureKey := hash.Hash(request.TransactionId, request.StatusCode, request.GrossAmount)
		if signatureKey != request.SignatureKey {
			return errors.New("invalid transaction")
		}

		transaction := te.Transaction{
			TransactionId: request.TransactionId,
			PaymentStatus: request.TransactionStatus,
			PaymentMethod: request.PaymentType,
		}

		if request != nil {
			if request.TransactionStatus == "capture" {
				if request.FraudStatus == "challenge" {
					transaction.PaymentStatus = "challenge"
				} else if request.FraudStatus == "accept" {
					transaction.PaymentStatus = "success"
				}
			} else if request.TransactionStatus == "settlement" {
				transaction.PaymentStatus = "success"
			} else if request.TransactionStatus == "deny" {
				transaction.PaymentStatus = "denied"
			} else if request.TransactionStatus == "cancel" || request.TransactionStatus == "expire" {
				transaction.PaymentStatus = "failure"
			} else if request.TransactionStatus == "pending" {
				transaction.PaymentStatus = "pending"
			}
		}

		if err := tc.transactionRepo.UpdateTransaction(&transaction).Error; err != nil {
			return errors.New("invalid transaction")
		}
	}

	return nil
}
