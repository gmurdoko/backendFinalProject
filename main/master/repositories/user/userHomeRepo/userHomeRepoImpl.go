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
func (ur *UserHomeRepoImpl) GetSaldo(id string) (*models.UserWallet, error) {
	var userSaldo = new(models.UserWallet)
	row := ur.db.QueryRow(utils.SELECT_USER_SALDO, id)
	err := row.Scan(&userSaldo.ID, &userSaldo.Saldo)
	if err != nil {
		return nil, err
	}
	return userSaldo, nil
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
	_, err = tx.Exec(utils.UPDATE_DATA_USER, user.Address, user.BornDate, editedAt, id)
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
func (ur *UserHomeRepoImpl) GetUserPhoto(id string) (*string, error) {
	var photo = new(string)
	err := ur.db.QueryRow(utils.SELECT_PHOTO_USER, id).Scan(&photo)
	if err != nil {
		return nil, err
	}
	log.Print(photo)
	return photo, nil
}

func (ur *UserHomeRepoImpl) UpdateUserPhoto(photo string, id string) error {
	var users models.UserModel
	users.EditedAt = time.Now().Format(utils.DATE_FORMAT)
	tx, err := ur.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(utils.UPDATE_PHOTO_USER, photo, users.EditedAt, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
