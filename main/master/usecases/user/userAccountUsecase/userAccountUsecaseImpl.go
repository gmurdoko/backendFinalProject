package userAccountUsecase

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/user/userAccountRepo"
)

type UserAccUsecaseImpl struct {
	userRepo userAccountRepo.UserAccount
}

func InitUseAccUsecase(userRepo userAccountRepo.UserAccount) UserAccount {
	return &UserAccUsecaseImpl{userRepo: userRepo}
}

func (uc *UserAccUsecaseImpl) GetUser(user *models.UserModel) (bool, error) {
	isValid, err := uc.userRepo.GetUser(user)
	if err != nil {
		return false, err
	}
	return isValid, nil
}
func (uc *UserAccUsecaseImpl) CreateUser(user *models.UserModel) (*models.UserModel, error) {
	data, err := uc.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return data, nil
}
