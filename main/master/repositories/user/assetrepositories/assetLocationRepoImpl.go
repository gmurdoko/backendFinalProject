package assetrepositories

import (
	"database/sql"
	"finalproject/main/master/models"
	constanta "finalproject/utils/constant"
	"log"
)

type AssetsLocationRepoImpl struct {
	db *sql.DB
}

func InitAssetsLocationRepoImpl(mydb *sql.DB) AssetsLocationRepo {
	return &AssetsLocationRepoImpl{db: mydb}
}

func (s *AssetsLocationRepoImpl) ReadAssetsLocation() ([]*models.AssetLocation, error) {
	query := constanta.READ_LOCATION_ASSET
	rows, err := s.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	println("masuk")
	var listAssetsLocation []*models.AssetLocation
	for rows.Next() {
		asset := models.AssetLocation{}
		err := rows.Scan(&asset.ID, &asset.AssetName, &asset.Longitude, &asset.Latitude)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		log.Println(asset.AssetName)
		listAssetsLocation = append(listAssetsLocation, &asset)
	}
	return listAssetsLocation, nil
}
