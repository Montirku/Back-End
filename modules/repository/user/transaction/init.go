package transaction

import (
	te "github.com/fazaalexander/montirku-be/modules/entity/transaction"
	ue "github.com/fazaalexander/montirku-be/modules/entity/user"
	"gorm.io/gorm"
)

type TransactionRepo interface {
	GetUserById(id uint) (*ue.User, error)
	CreateTransaction(transaction *te.Transaction) error
	UpdateTransaction(transactionData *te.Transaction) error
}

type transactionRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) TransactionRepo {
	return &transactionRepo{
		db,
	}
}
