package providerrepository

import (
	"database/sql"
	"finalproject/main/master/models"
)

//providerRepositortImpl is for init Repository
type providerRepositorytImpl struct {
	db *sql.DB
}

func (s providerRepositorytImpl) DeletePhotoProvider(id string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	query := "UPDATE m_provider_account SET photo = '' WHERE id=?;"
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s providerRepositorytImpl) UpdateDataProvider(Provider *models.ProviderModel) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	query := "UPDATE m_provider_account SET address =?, borndate=? WHERE id =?"
	_, err = tx.Exec(query, Provider.Address, Provider.BornDate, Provider.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil

}

//InitProviderRepositoryImpl is init gate for repository
func InitProviderRepositoryImpl(db *sql.DB) ProviderRepository {
	return &providerRepositorytImpl{db}
}
