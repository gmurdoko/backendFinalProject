package providerrepository

import (
	"database/sql"
	"finalproject/main/master/models"
	"finalproject/utils"
	"time"
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
	query := "UPDATE m_provider_account SET photo = '' WHERE id=?"
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s providerRepositorytImpl) UpdateDataProvider(id string, Provider *models.Providers) (*models.Providers, error) {
	editedAt := time.Now()
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	query := "UPDATE m_provider_account SET address =?, borndate=?,edited_at=? WHERE id =?"
	_, err = tx.Exec(query, Provider.Address, Provider.Borndate, editedAt, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	providers := new(models.Providers)
	var bornDate, editAt, deletedAt sql.NullString

	err = s.db.QueryRow(utils.SELECT_NEW_PROVIDER, id).Scan(&providers.ID, &providers.Username, &providers.Password,
		&providers.Email, &providers.Fullname, &providers.Photo, &bornDate, &providers.PhoneNumber, &providers.Address,
		&providers.CreatedAt, &editAt, &deletedAt, &providers.Status)
	providers.Borndate = bornDate.String
	providers.EditedAt = editAt.String
	providers.DeletedAt = deletedAt.String
	if err != nil {
		return nil, err
	}
	return providers, nil

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
