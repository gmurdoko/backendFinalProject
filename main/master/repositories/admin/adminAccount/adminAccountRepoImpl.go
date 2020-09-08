package adminAccount

import (
	"database/sql"
	"finalproject/main/master/models"
	"finalproject/utils"
	"finalproject/utils/pwd"
	"fmt"
)

type AdminAccountRepoImpl struct {
	db *sql.DB
}

func InitAdminAccountImpl(mydb *sql.DB) AdminAccount {
	return &AdminAccountRepoImpl{db: mydb}
}
func (ac *AdminAccountRepoImpl) AdminLogin(admin *models.Admin) (bool, error) {
	row := ac.db.QueryRow(utils.SELECT_ADMIN_LOGIN, admin.Username)
	var admins = models.Admin{}
	err := row.Scan(&admins.ID, &admins.Username, &admins.Password)

	if err != nil {
		return false, err
	}
	isPwdValid := pwd.CheckPasswordHash(admin.Password, admins.Password)
	fmt.Println(isPwdValid)

	if admin.Username == admins.Username && isPwdValid {
		return true, nil
	} else {
		return false, err
	}
}
