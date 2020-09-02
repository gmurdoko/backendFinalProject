package providerHomeRepo

import (
	"database/sql"
	"finalproject/main/master/models"
	"finalproject/utils"
	"log"
	"time"

	"github.com/google/uuid"
)

type ProviderHomeRepoImpl struct {
	db *sql.DB
}

func InitProviderHomeRepoImpl(db *sql.DB) ProviderHome {
	return &ProviderHomeRepoImpl{db: db}
}
func (pr *ProviderHomeRepoImpl) GetProviderSaldo(id string) (int, error) {
	row := pr.db.QueryRow(utils.SELECT_PROVIDER_SALDO, id)
	var saldo int
	err := row.Scan(&saldo)
	if err != nil {
		return 0, err
	}
	return saldo, nil
}
func (pr *ProviderHomeRepoImpl) CreateAssetProvider(asset *models.AssetModel) (*models.AssetModel, error) {
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
	asset.IdWallet = wallet.ID
	_, err = tx.Exec(utils.INSERT_ASSET, asset.ID, asset.IdWallet, asset.ProviderId,
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
