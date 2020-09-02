package providerAccountRepo

import (
	"database/sql"
	"finalproject/main/master/models"
	"finalproject/utils"
	"finalproject/utils/pwd"
	"fmt"
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
func (pr *ProviderRepoAccountImpl) GetProvider(provider *models.Providers) (*models.Providers, bool, error) {
	row := pr.db.QueryRow(utils.SELECT_PROVIDER, provider.Username)
	var providers = models.Providers{}
	var bornDate, editedAt, deletedAt sql.NullString
	err := row.Scan(&providers.ID, &providers.Username, &providers.Password,
		&providers.Email, &providers.Fullname, &providers.Photo, &bornDate, &providers.Address, &providers.PhoneNumber,
		&providers.CreatedAt, &editedAt, &deletedAt, &providers.Status)
	providers.Borndate = bornDate.String
	providers.EditedAt = editedAt.String
	providers.DeletedAt = deletedAt.String
	if err != nil {
		return nil, false, err
	}
	isPwdValid := pwd.CheckPasswordHash(provider.Password, providers.Password)
	if provider.Username == provider.Username && isPwdValid {
		data, _ := pr.GetProviderById(providers.ID)
		return data, true, nil
	} else {
		return nil, false, err
	}
}
func (pr *ProviderRepoAccountImpl) CreateProvider(provider *models.Providers) (*models.Providers, error) {
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
	providers, _ := pr.GetProviderById(provider.ID)
	return providers, nil
}
func (pr *ProviderRepoAccountImpl) GetProviderById(id string) (*models.Providers, error) {
	providers := new(models.Providers)
	var bornDate, editedAt, deletedAt sql.NullString

	err := pr.db.QueryRow(utils.SELECT_NEW_USER, id).Scan(&providers.ID, &providers.Username, &providers.Password,
		&providers.Email, &providers.Fullname, &providers.Photo, &bornDate, &providers.Address, &providers.PhoneNumber,
		&providers.CreatedAt, &editedAt, &deletedAt, &providers.Status)
	providers.Borndate = bornDate.String
	providers.EditedAt = editedAt.String
	providers.DeletedAt = deletedAt.String
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return providers, nil
}
