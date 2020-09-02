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
func (ur *UserAccRepoImpl) GetUser(user *models.UserModel) (*models.UserModel, bool, error) {
	row := ur.db.QueryRow(utils.SELECT_USER, user.Username)
	var users = models.UserModel{}
	var bornDate, editedAt, deletedAt sql.NullString
	err := row.Scan(&users.ID, &users.IdWallet, &users.Username, &users.Password,
		&users.Email, &users.Fullname, &users.Photo, &bornDate, &users.Address, &users.PhoneNumber,
		&users.CreatedAt, &editedAt, &deletedAt, &users.Status)
	users.BornDate = bornDate.String
	users.EditedAt = editedAt.String
	users.DeletedAt = deletedAt.String
	if err != nil {
		fmt.Println(err)
		return nil, false, err
	}
	isPwdValid := pwd.CheckPasswordHash(user.Password, users.Password)

	if user.Username == users.Username && isPwdValid {
		data, _ := ur.GetUserById(users.ID)
		fmt.Println("repo", data)
		return data, true, nil
	} else {
		return nil, false, err
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
	tx.Commit()
	users, _ := ur.GetUserById(user.ID)
	return users, nil
}

func (ur *UserAccRepoImpl) GetUserById(id string) (*models.UserModel, error) {
	users := new(models.UserModel)
	var bornDate, editedAt, deletedAt sql.NullString

	err := ur.db.QueryRow(utils.SELECT_NEW_USER, id).Scan(&users.ID, &users.IdWallet, &users.Username, &users.Password,
		&users.Email, &users.Fullname, &users.Photo, &bornDate, &users.Address, &users.PhoneNumber,
		&users.CreatedAt, &editedAt, &deletedAt, &users.Status)
	users.BornDate = bornDate.String
	users.EditedAt = editedAt.String
	users.DeletedAt = deletedAt.String
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return users, nil
}
