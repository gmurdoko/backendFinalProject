package adminAccountUsecase

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/admin/adminAccount"
)

type AdminAccountUsecaseImpl struct {
	adminAccRepo adminAccount.AdminAccount
}

func InitAdminAccountUsecaseImpl(adminAccRepo adminAccount.AdminAccount) AdminAccountUsecase {
	// return &ListAssetsUsecaseImpl{listAssetsRepo: listAssetsRepo}
	return &AdminAccountUsecaseImpl{adminAccRepo: adminAccRepo}
}

func (ac *AdminAccountUsecaseImpl) AdminLogin(admin *models.Admin) (bool, error) {

	isValid, err := ac.adminAccRepo.AdminLogin(admin)

	if err != nil {
		return false, err
	}

	return isValid, nil
}
