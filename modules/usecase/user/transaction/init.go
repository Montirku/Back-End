package transaction

import (
	mr "github.com/fazaalexander/montirku-be/modules/entity/midtrans"
	te "github.com/fazaalexander/montirku-be/modules/entity/transaction"
	tr "github.com/fazaalexander/montirku-be/modules/repository/user/transaction"
)

type TransactionUseCase interface {
	CreateTransaction(transaction *te.Transaction) (string, string, error)
	MidtransNotification(request *mr.MidtransRequest) error
}

type transactionUseCase struct {
	transactionRepo tr.TransactionRepo
}

func New(transactionRepo tr.TransactionRepo) *transactionUseCase {
	return &transactionUseCase{
		transactionRepo,
	}
}
