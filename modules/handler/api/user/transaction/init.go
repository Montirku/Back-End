package transaction

import (
	tc "github.com/fazaalexander/montirku-be/modules/usecase/user/transaction"
)

type TransactionHandler struct {
	transactionUseCase tc.TransactionUseCase
}

func New(transactionUseCase tc.TransactionUseCase) *TransactionHandler {
	return &TransactionHandler{
		transactionUseCase,
	}
}
