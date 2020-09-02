package providerAccountRepo

import (
	"database/sql"
	"finalproject/main/master/models"
	"finalproject/utils"
	"finalproject/utils/pwd"
	"log"
	"time"

	"github.com/google/uuid"
)

type ProviderRepoAccountImpl struct {
	db *sql.DB
}

func InitProviderRepoAccImpl(db *sql.DB) ProviderAccount {
	return &ProviderRepoAccountImpl{db: db}
}
func (pr *ProviderRepoAccountImpl) GetProvider(provider *models.ProviderModel) (bool, error) {
	row := pr.db.QueryRow(utils.SELECT_PROVIDER, provider.Username)
	var providers = models.UserModel{}
	err := row.Scan(&providers.Username, &providers.Password)
	if err != nil {
		return false, err
	}
	isPwdValid := pwd.CheckPasswordHash(provider.Password, providers.Password)
	if provider.Username == provider.Username && isPwdValid {
		return true, nil
	} else {
		return false, err
	}
}
func (pr *ProviderRepoAccountImpl) CreateProvider(provider *models.ProviderModel) (*models.ProviderModel, error) {
	provider.ID = uuid.New().String()
	provider.CreatedAt = time.Now().Format(utils.DATE_FORMAT)
	tx, err := pr.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	password, _ := pwd.HashPassword(provider.Password)
	_, err = tx.Exec(utils.INSERT_PROVIDER_ACCOUNT, provider.ID, provider.Username,
		password, provider.Email, provider.Fullname, provider.PhoneNumber,
		provider.CreatedAt)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return nil, err
	}
	tx.Commit()
	return provider, nil
}
