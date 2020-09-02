package providerrepository

import (
	"database/sql"
	"finalproject/main/master/model"
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

func (s providerRepositorytImpl) UpdateDataProvider(Provider *model.Providers) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	query := "UPDATE m_provider_account SET address =?, borndate=? WHERE id =?"
	_, err = tx.Exec(query, Provider.Address, Provider.Borndate, Provider.ID)
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
