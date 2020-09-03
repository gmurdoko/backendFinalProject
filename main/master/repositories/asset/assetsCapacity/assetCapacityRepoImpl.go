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

func (s *AssetCapacityRepoImpl) ReadCurrentCapacity(assetId string) (*models.AssetCapacity, error) {
	query := constant.READ_CURRENT_CAPACITY
	row, err := s.db.Query(query, assetId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var ac models.AssetCapacity
	err = row.Scan(&ac.CarCapMax, &ac.CarCapAvailable,
		&ac.MotorCapMax, &ac.MotorCapAvailable,
		&ac.BicycleCapMax, &ac.BicycleCapAvailable)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &ac, err
}

