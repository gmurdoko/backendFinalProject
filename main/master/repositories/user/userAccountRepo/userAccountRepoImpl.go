package userAccountRepo

import (
	"database/sql"
	"finalproject/main/master/models"
	"finalproject/utils"
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
