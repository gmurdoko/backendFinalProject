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

func (ac *AdminAccountUsecaseImpl) AdminLogin(admin *models.Admin) (bool, *models.Admin, error) {

	isValid, admin, err := ac.adminAccRepo.AdminLogin(admin)

	if err != nil {
		return false, nil, err
	}

	return isValid, admin, nil
}
