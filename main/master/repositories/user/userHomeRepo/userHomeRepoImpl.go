package userHomeRepo

import (
	"database/sql"
	"finalproject/main/master/models"
	"finalproject/utils"
	"log"
	"time"
)

type UserHomeRepoImpl struct {
	db *sql.DB
}

func InitUserHomeRepoImpl(db *sql.DB) UserHome {
	return &UserHomeRepoImpl{db: db}
}
func (ur *UserHomeRepoImpl) GetSaldo(id string) (int, error) {
	row := ur.db.QueryRow(utils.SELECT_USER_SALDO, id)
	var saldo int
	err := row.Scan(&saldo)
	if err != nil {
		return 0, err
	}
	return saldo, nil
}
func (ur *UserHomeRepoImpl) DeleteUserPhoto(id string) error {
	deletedAt := time.Now()
	tx, err := ur.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(utils.DELETE_PHOTO_USER, deletedAt, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
func (ur *UserHomeRepoImpl) UpdateUserData(user *models.UserModel, id string) (*models.UserModel, error) {
	editedAt := time.Now()
	tx, err := ur.db.Begin()
	if err != nil {
		return nil, err
	}
	_, err = tx.Exec(utils.UPDATE_DATA_USER, user.Address, user.BornDate, user.Photo, editedAt, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return user, tx.Commit()
}
func (ur *UserHomeRepoImpl) UpdateUserSaldoTopUp(wallet *models.Wallets, id string) (int, error) {
	editedAt := time.Now()
	tx, err := ur.db.Begin()
	_, err = tx.Exec(utils.UPDATE_USER_SALDO_TOPUP, wallet.Debit, editedAt, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	row := ur.db.QueryRow(utils.SELECT_UPDATED_SALDO_USER, id)

	var saldo int
	err = row.Scan(&saldo)
	if err != nil {
		return 0, err
	}

	return saldo, nil
}
func (ur *UserHomeRepoImpl) GetUserPhoto(id string) (string, error) {
	row := ur.db.QueryRow(utils.SELECT_PHOTO_USER, id)
	var photo string
	err := row.Scan(&photo)
	if err != nil {
		return "", err
	}
	log.Print(photo)
	return photo, nil
}

// func (ur *UserHomeRepoImpl) UpdateUserPhoto(user *models.UserModel, id string) (*models.UserModel, error) {
// 	editedAt := time.Now()
// 	tx, err := ur.db.Begin()
// 	_, err = tx.Exec(utils.UPDATE_USER_SALDO_TOPUP, user.Photo, editedAt, id)
// 	if err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}
// 	return user, tx.Commit()
// }
