package userUsecase

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/userRepo"
)

type UserUsecaseImpl struct {
	userRepo userRepo.UserRepository
}

func InitUserUsecase(userRepo userRepo.UserRepository) UserUsecase {
	return &UserUsecaseImpl{userRepo: userRepo}
}

func (uc *UserUsecaseImpl) CreateUser(user *models.UserModel) (*models.UserModel, error) {
	data, err := uc.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (uc *UserUsecaseImpl) GetSaldo(id string) (int, error) {
	data, err := uc.userRepo.GetSaldo(id)
	if err != nil {
		return 0, err
	}
	return data, nil
}
func (uc *UserUsecaseImpl) DeleteUserPhoto(id string) error {
	err := uc.userRepo.DeleteUserPhoto(id)
	return err
}
func (uc *UserUsecaseImpl) UpdateUserData(user *models.UserModel, id string) (*models.UserModel, error) {
	data, err := uc.userRepo.UpdateUserData(user, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (uc *UserUsecaseImpl) UpdateUserSaldoTopUp(wallet *models.WalletModel, id string) (int, error) {
	data, err := uc.userRepo.UpdateUserSaldoTopUp(wallet, id)
	if err != nil {
		return 0, err
	}
	return data, nil
}
func (uc *UserUsecaseImpl) GetUserPhoto(id string) (string, error) {
	data, err := uc.userRepo.GetUserPhoto(id)
	if err != nil {
		return "", err
	}
	return data, nil
}
func (uc *UserUsecaseImpl) UpdateUserPhoto(user *models.UserModel, id string) (string, error) {
	data, err := uc.userRepo.UpdateUserPhoto(user, id)
	if err != nil {
		return "", err
	}
	return data, nil
}
