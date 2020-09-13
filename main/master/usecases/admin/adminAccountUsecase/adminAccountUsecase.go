package adminAccountUsecase

import "finalproject/main/master/models"

type AdminAccountUsecase interface {
	AdminLogin(*models.Admin) (bool, *models.Admin, error)
}
