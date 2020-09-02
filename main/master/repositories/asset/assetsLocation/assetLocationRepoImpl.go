package assetsLocation

import (
	"database/sql"
	"finalproject/main/master/models"
	"finalproject/utils/constant"
	"log"
)

type AssetsLocationRepoImpl struct{
	db *sql.DB
}

func InitAssetsLocationRepoImpl(mydb *sql.DB) AssetsLocationRepo {
	return &AssetsLocationRepoImpl{db: mydb}
}

func (s *AssetsLocationRepoImpl) ReadAssetsLocation() ([]*models.Asset, error) {
	query := constant.READ_LOCATION_ASSET
	rows, err := s.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var listAssetsLocation []*models.Asset
	for rows.Next() {
		asset := models.Asset{}
		err := rows.Scan(&asset.AssetName, &asset.Longitude, &asset.Latitude)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		listAssetsLocation = append(listAssetsLocation, &asset)
	}
	return listAssetsLocation, nil
}



