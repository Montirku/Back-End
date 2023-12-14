package transaction

import (
	"errors"

	te "github.com/fazaalexander/montirku-be/modules/entity/transaction"
	ue "github.com/fazaalexander/montirku-be/modules/entity/user"
)

func (tr *transactionRepo) GetUserById(id uint) (*ue.User, error) {
	user := &ue.User{}
	err := tr.db.Preload("UserDetail").Model(user).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	return user, nil
}

func (tr *transactionRepo) CreateTransaction(transaction *te.Transaction) error {
	err := tr.db.Create(transaction).Error
	if err != nil {
		return err
	}

	return nil
}

func (tr *transactionRepo) UpdateTransaction(transactionData *te.Transaction) error {
	if err := tr.db.Model(te.Transaction{}).Where("transaction_id = ?", transactionData.ID).Updates(transactionData).Error; err != nil {
		return err
	}

	return nil
}
