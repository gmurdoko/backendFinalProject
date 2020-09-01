package providerRepo

import (
	"database/sql"
	"finalproject/main/master/models"
	"finalproject/utils"
	"log"
	"time"

	"github.com/google/uuid"
)

type ProviderRepoImpl struct {
	db *sql.DB
}

func InitProviderRepoImpl(db *sql.DB) ProviderRepository {
	return &ProviderRepoImpl{db: db}
}
func (pr *ProviderRepoImpl) CreateProvider(provider *models.ProviderModel) (*models.ProviderModel, error) {
	provider.ID = uuid.New().String()
	provider.CreatedAt = time.Now().Format(`2006-01-02 15:04:05`)
	tx, err := pr.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	_, err = tx.Exec(utils.INSERT_PROVIDER_ACCOUNT, provider.ID, provider.Username,
		provider.Password, provider.Email, provider.Fullname, provider.PhoneNumber,
		provider.CreatedAt)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return nil, err
	}
	tx.Commit()
	return provider, nil
}
