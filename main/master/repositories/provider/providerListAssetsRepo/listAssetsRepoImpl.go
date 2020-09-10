package providerListAssetsRepo

import (
	"database/sql"
	"finalproject/main/master/models"
	constanta "finalproject/utils/constant"
	"log"
)

type ListAssetsRepoImpl struct {
	db *sql.DB
}

func InitListAssetsRepoImpl(mydb *sql.DB) ListAssetsRepo {
	return &ListAssetsRepoImpl{db: mydb}
}

func (s *ListAssetsRepoImpl) GetAllAssets(userId string) ([]*models.Assets, error) {
	query := constanta.GETALLASSETSPERPROVIDER
	rows, err := s.db.Query(query, userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var listAssets []*models.Assets
	for rows.Next() {
		asset := models.Assets{}
		err := rows.Scan(&asset.ID, &asset.AssetName, &asset.AssetArea, &asset.Longitude, &asset.Latitude, &asset.CarCap, &asset.MotorCap, &asset.BicycleCap, &asset.Photo, &asset.Saldo, &asset.Status)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		listAssets = append(listAssets, &asset)
	}
	return listAssets, nil
}
