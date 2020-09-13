package adminAccount

import "finalproject/main/master/models"

type AdminAccount interface {
	AdminLogin(*models.Admin) (bool, *models.Admin, error)
}
