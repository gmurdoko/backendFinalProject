package userAccountUsecase

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/user/userAccountRepo"
	"fmt"
)

type UserAccUsecaseImpl struct {
	userRepo userAccountRepo.UserAccount
}

func InitUseAccUsecase(userRepo userAccountRepo.UserAccount) UserAccount {
	return &UserAccUsecaseImpl{userRepo: userRepo}
}

func (uc *UserAccUsecaseImpl) GetUser(user *models.UserModel) (*models.UserModel, bool, error) {

	data, isValid, err := uc.userRepo.GetUser(user)

	if err != nil {
		return nil, false, err
	}

	return data, isValid, nil
}
func (uc *UserAccUsecaseImpl) CreateUser(user *models.UserModel) (*models.UserModel, error) {
	data, err := uc.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	fmt.Println("usecase", data)
	return data, nil
}
func (uc *UserAccUsecaseImpl) GetUserById(id string) (*models.UserModel, error) {
	data, err := uc.userRepo.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
