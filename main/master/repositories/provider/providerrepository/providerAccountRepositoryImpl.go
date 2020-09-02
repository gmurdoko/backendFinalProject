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

func (s providerRepositorytImpl) UpdateDataProvider(Provider *models.Providers) error {
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
func (s providerRepositorytImpl) UpdatePhotoProvider(photo, id string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	query := "UPDATE m_provider_account SET photo=? WHERE id=?;"
	_, err = tx.Exec(query, photo, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s providerRepositorytImpl) SelectPhotoByID(id string) (*string, error) {
	var photo = new(string)
	err := s.db.QueryRow("SELECT photo FROM m_provider_account WHERE id=?;", id).Scan(photo)
	if err != nil {
		return nil, err
	}
	return photo, nil

}

//InitProviderRepositoryImpl is init gate for repository
func InitProviderRepositoryImpl(db *sql.DB) ProviderRepository {
	return &providerRepositorytImpl{db}
}
