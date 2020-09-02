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
func (pr *ProviderRepoImpl) GetProvider(provider *models.ProviderModel) (bool, error) {
	row := pr.db.QueryRow(utils.SELECT_PROVIDER, provider.Username, provider.Password)
	var users = models.UserModel{}
	err := row.Scan(&users.Username, &users.Password)
	if err != nil {
		return false, err
	}
	if provider.Username == provider.Username && provider.Password == provider.Password {
		return true, nil
	} else {
		return false, err
	}
}
func (pr *ProviderRepoImpl) CreateProvider(provider *models.ProviderModel) (*models.ProviderModel, error) {
	provider.ID = uuid.New().String()
	provider.CreatedAt = time.Now().Format(utils.DATE_FORMAT)
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
func (pr *ProviderRepoImpl) GetProviderSaldo(id string) (int, error) {
	row := pr.db.QueryRow(utils.SELECT_PROVIDER_SALDO, id)
	var saldo int
	err := row.Scan(&saldo)
	if err != nil {
		return 0, err
	}
	return saldo, nil
}
func (pr *ProviderRepoImpl) CreateAssetProvider(asset *models.AssetModel) (*models.AssetModel, error) {
	var wallet models.WalletModel
	asset.ID = uuid.New().String()
	wallet.ID = uuid.New().String()
	asset.CreatedAt = time.Now().Format(utils.DATE_FORMAT)
	tx, err := pr.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	_, err = tx.Exec(utils.INSERT_WALLET, wallet.ID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	_, err = tx.Exec(utils.INSERT_ASSET, asset.ID, asset.IdWallet.ID, asset.ProviderId.ID,
		asset.AssetName, asset.AssetArea, asset.Longitude, asset.Latitude,
		asset.CarCapacity, asset.MotorcycleCapacity, asset.BicycleCapacity,
		asset.Photo, asset.CreatedAt)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return nil, err
	}
	tx.Commit()
	return asset, nil
}
