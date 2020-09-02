package repositories

import (
	"database/sql"
	"finalproject/main/master/models"
	"finalproject/utils/constant"
	"log"
)

type AssetCapacityRepoImpl struct {
	db *sql.DB
}

func InitAssetCapacityRepoImpl(mydb *sql.DB) AssetCapacityRepo {
	return &AssetCapacityRepoImpl{db: mydb}
}

func (s *AssetCapacityRepoImpl) ReadCurrentCapacity(ac *models.AssetCapacity, assetId string) (*models.AssetCapacity, error) {
	//panic("implement me")
	query := constant.READ_CURRENT_CAPACITY
	row, err := s.db.Query(query, assetId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = row.Scan(&ac.CarCapMax, &ac.CarCapAvailable,
		&ac.MotorCapMax, &ac.MotorCapAvailable,
		&ac.BicycleCapMax, &ac.BicycleCapAvailable)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return ac, err
}

