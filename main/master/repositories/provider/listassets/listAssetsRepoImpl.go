package listassetsrepo

import (
	"database/sql"
	"finalproject/main/master/models"
	"log"
)

type ListAssetsRepoImpl struct {
	db *sql.DB
}

func InitListAssetsRepoImpl(db *sql.DB) ListAssetsRepo {
	return &ListAssetsRepoImpl{db: db}
}

func (s *ListAssetsRepoImpl) GetAllAssets(userId string) ([]*models.Assets, error) {
	query := constanta
	rows, err := s.db.Query(query, userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var listAssets []*models.Assets
	for rows.Next() {
		asset := models.Assets{}
		err := rows.Scan(&asset.AssetName, &asset.AssetArea, &asset.Longitude, &asset.Latitude, &asset.CarCap, &asset.MotorCap, &asset.BicycleCap, &asset.Photo, &asset.Saldo, &asset.Status)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		listAssets = append(listAssets, &asset)
	}
	return listAssets, nil
}
