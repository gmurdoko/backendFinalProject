package userHomeUsecase

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/user/userHomeRepo"
)

type UserHomeUsecaseImpl struct {
	userRepo userHomeRepo.UserHome
}

func InitUserHomeUsecase(userRepo userHomeRepo.UserHome) UserHome {
	return &UserHomeUsecaseImpl{userRepo: userRepo}
}
func (uc *UserHomeUsecaseImpl) GetSaldo(id string) (int, error) {
	data, err := uc.userRepo.GetSaldo(id)
	if err != nil {
		return 0, err
	}
	return data, nil
}
func (uc *UserHomeUsecaseImpl) DeleteUserPhoto(id string) error {
	err := uc.userRepo.DeleteUserPhoto(id)
	return err
}
func (uc *UserHomeUsecaseImpl) UpdateUserData(user *models.UserModel, id string) (*models.UserModel, error) {
	data, err := uc.userRepo.UpdateUserData(user, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (uc *UserHomeUsecaseImpl) UpdateUserSaldoTopUp(wallet *models.WalletModel, id string) (int, error) {
	data, err := uc.userRepo.UpdateUserSaldoTopUp(wallet, id)
	if err != nil {
		return 0, err
	}
	return data, nil
}
func (uc *UserHomeUsecaseImpl) GetUserPhoto(id string) (string, error) {
	data, err := uc.userRepo.GetUserPhoto(id)
	if err != nil {
		return "", err
	}
	return data, nil
}

// func (uc *UserUsecaseImpl) UpdateUserPhoto(user *models.UserModel, id string) (*models.UserModel, error) {
// 	data, err := uc.userRepo.UpdateUserPhoto(user, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return data, nil
// }
