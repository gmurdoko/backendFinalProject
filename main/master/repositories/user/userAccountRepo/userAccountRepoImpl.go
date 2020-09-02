package userAccountRepo

import (
	"database/sql"
	"finalproject/main/master/models"
	"finalproject/utils"
	"finalproject/utils/pwd"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type UserAccRepoImpl struct {
	db *sql.DB
}

func InitUserAccRepoImpl(db *sql.DB) UserAccount {
	return &UserAccRepoImpl{db: db}
}
func (ur *UserAccRepoImpl) GetUser(user *models.UserModel) (bool, error) {
	row := ur.db.QueryRow(utils.SELECT_USER, user.Username)
	var users = models.UserModel{}
	err := row.Scan(&users.Username, &users.Password)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	isPwdValid := pwd.CheckPasswordHash(user.Password, users.Password)
	if user.Username == users.Username && isPwdValid {
		return true, nil
	} else {
		return false, err
	}
}
func (ur *UserAccRepoImpl) CreateUser(user *models.UserModel) (*models.UserModel, error) {
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
	user.IdWallet = wallet.ID
	password, _ := pwd.HashPassword(user.Password)
	_, err = tx.Exec(utils.INSERT_USER_ACCOUNT, user.ID, user.IdWallet, user.Username,
		password, user.Email, user.Fullname, user.PhoneNumber,
		user.CreatedAt)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return nil, err
	}
	var users *models.UserModel
	err = tx.Commit()
	if err == nil {
		fmt.Println("id", user.ID)
		row := ur.db.QueryRow(utils.SELECT_NEW_USER, user.ID)
		fmt.Println("row", row)

		err = row.Scan(&users.ID, &users.IdWallet, &users.Username, &users.Password,
			&users.Email, &users.Fullname, &users.Photo, &users.BornDate, &users.Address, &users.PhoneNumber,
			&users.CreatedAt, &users.EditedAt, &users.DeletedAt, &users.Status)
	}
	return users, nil

}
