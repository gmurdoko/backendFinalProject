package userRepo

import (
	"database/sql"
	"finalproject/main/master/models"
	"finalproject/utils"
	"log"
	"time"

	"github.com/google/uuid"
)

type UserRepoImpl struct {
	db *sql.DB
}

func InitUserRepoImpl(db *sql.DB) UserRepository {
	return &UserRepoImpl{db: db}
}
func (ur *UserRepoImpl) GetUser(user *models.UserModel) (bool, error) {
	row := ur.db.QueryRow(utils.SELECT_USER, user.Username, user.Password)
	var users = models.UserModel{}
	err := row.Scan(&users.Username, &users.Password)
	if err != nil {
		return false, err
	}
	if user.Username == user.Username && user.Password == user.Password {
		return true, nil
	} else {
		return false, err
	}
}
func (ur *UserRepoImpl) CreateUser(user *models.UserModel) (*models.UserModel, error) {
	var wallet models.WalletModel
	user.ID = uuid.New().String()
	wallet.ID = uuid.New().String()
	user.CreatedAt = time.Now().Format(utils.DATE_FORMAT)
	tx, err := ur.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	_, err = tx.Exec(utils.INSERT_WALLET, wallet.ID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	_, err = tx.Exec(utils.INSERT_USER_ACCOUNT, user.ID, user.IdWallet.ID, user.Username,
		user.Password, user.Email, user.Fullname, user.PhoneNumber,
		user.CreatedAt)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return nil, err
	}
	tx.Commit()
	return user, nil
}
func (ur *UserRepoImpl) GetSaldo(id string) (int, error) {
	row := ur.db.QueryRow(utils.SELECT_USER_SALDO, id)
	var saldo int
	err := row.Scan(&saldo)
	if err != nil {
		return 0, err
	}
	return saldo, nil
}
func (ur *UserRepoImpl) DeleteUserPhoto(id string) error {
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
func (ur *UserRepoImpl) UpdateUserData(user *models.UserModel, id string) (*models.UserModel, error) {
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
func (ur *UserRepoImpl) UpdateUserSaldoTopUp(wallet *models.WalletModel, id string) (int, error) {
	editedAt := time.Now()
	tx, err := ur.db.Begin()
	_, err = tx.Exec(utils.UPDATE_USER_SALDO_TOPUP, wallet.Debit, editedAt, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	row := ur.db.QueryRow(utils.SELECT_UPDATED_SALDO_USER, id)

	var saldo int
	err = row.Scan(&saldo)
	if err != nil {
		return 0, err
	}

	return saldo, tx.Commit()
}
func (ur *UserRepoImpl) GetUserPhoto(id string) (string, error) {
	row := ur.db.QueryRow(utils.SELECT_PHOTO_USER, id)
	var photo string
	err := row.Scan(&photo)
	if err != nil {
		return "", err
	}
	log.Print(photo)
	return photo, nil
}

// func (ur *UserRepoImpl) UpdateUserPhoto(user *models.UserModel, id string) (*models.UserModel, error) {
// 	editedAt := time.Now()
// 	tx, err := ur.db.Begin()
// 	_, err = tx.Exec(utils.UPDATE_USER_SALDO_TOPUP, user.Photo, editedAt, id)
// 	if err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}
// 	return user, tx.Commit()
// }
