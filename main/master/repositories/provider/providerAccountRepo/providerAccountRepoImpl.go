package providerAccountRepo

import (
	"database/sql"
	"errors"
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
	row := pr.db.QueryRow(utils.SELECT_PROVIDER, provider.Username, provider.Email)
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
	if provider.Username == provider.Username && providers.Status == "A" || provider.Email == providers.Email && isPwdValid {
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
	row := pr.db.QueryRow(utils.SELECT_PROVIDER_EXIST, provider.Username, provider.Email)
	var checkproviders = models.CheckProvider{}
	err = row.Scan(&checkproviders.Username, &checkproviders.Email)
	if checkproviders.Username != provider.Username || checkproviders.Email != provider.Email {
		_, err = tx.Exec(utils.INSERT_PROVIDER_ACCOUNT, provider.ID, provider.Username,
			password, provider.Email, provider.Fullname, provider.PhoneNumber,
			provider.CreatedAt)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return nil, err
		}

		providers, _ := pr.GetProviderById(provider.ID)
		tx.Commit()
		return providers, nil
	} else {
		tx.Rollback()
		return nil, errors.New("Username or Email Exist")
	}
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
