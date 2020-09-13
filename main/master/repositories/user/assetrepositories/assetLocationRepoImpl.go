package assetrepositories

import (
	"database/sql"
	"finalproject/main/master/models"

	// constanta "finalproject/utils/constant"
	"log"
)

type AssetsLocationRepoImpl struct {
	db *sql.DB
}

func InitAssetsLocationRepoImpl(mydb *sql.DB) AssetsLocationRepo {
	return &AssetsLocationRepoImpl{db: mydb}
}

func (s *AssetsLocationRepoImpl) GetAssetsByID(id string) (*models.AssetLocation, error) {
	query := `SELECT 
    m_asset.id,
    m_asset.asset_name,
    m_asset.longitude,
    m_asset.latitude,
    m_asset.photo,
    (SELECT 
            AVG(m_review.rating)
        FROM
            m_review
        WHERE
            m_asset.id = asset_id),
    car_capacity - (SELECT 
            COUNT(m_ticket.id)
        FROM
            m_ticket
        WHERE
            vehicle_id = '1'
                AND (status = 'A' OR status = 'B')
                AND m_asset.id = asset_id) car_capacity_avaliable,
    motorcycle_capacity - (SELECT 
            COUNT(m_ticket.id)
        FROM
            m_ticket
        WHERE
            vehicle_id = '2'
                AND (status = 'A' OR status = 'B')
                AND m_asset.id = asset_id) motorcycle_capacity_avaliable,
    bicycle_capacity - (SELECT 
            COUNT(m_ticket.id)
        FROM
            m_ticket
        WHERE
            vehicle_id = '3'
                AND (status = 'A' OR status = 'B')
                AND m_asset.id = asset_id) bicycle_capacity_avaliable
FROM
    m_asset
WHERE
    m_asset.status = 'A' AND id=?`
	var rate sql.NullString
	row := s.db.QueryRow(query, id)
	var asset = models.AssetLocation{}

	err := row.Scan(&asset.ID, &asset.AssetName, &asset.Longitude, &asset.Latitude, &asset.Photo, &rate, &asset.CarCap, &asset.MotorCap, &asset.BicycleCap)
	if err != nil {
		return nil, err
	}
	asset.Rating = rate.String

	return &asset, nil
}

func (s *AssetsLocationRepoImpl) ReadAssetsLocation() ([]*models.AssetLocation, error) {
	query := `SELECT 
    m_asset.id,
    m_asset.asset_name,
    m_asset.longitude,
    m_asset.latitude,
    m_asset.photo,
    (SELECT 
            AVG(m_review.rating)
        FROM
            m_review
        WHERE
            m_asset.id = asset_id),
    car_capacity - (SELECT 
            COUNT(m_ticket.id)
        FROM
            m_ticket
        WHERE
            vehicle_id = '1'
                AND (status = 'A' OR status = 'B')
                AND m_asset.id = asset_id) car_capacity_avaliable,
    motorcycle_capacity - (SELECT 
            COUNT(m_ticket.id)
        FROM
            m_ticket
        WHERE
            vehicle_id = '2'
                AND (status = 'A' OR status = 'B')
                AND m_asset.id = asset_id) motorcycle_capacity_avaliable,
    bicycle_capacity - (SELECT 
            COUNT(m_ticket.id)
        FROM
            m_ticket
        WHERE
            vehicle_id = '3'
                AND (status = 'A' OR status = 'B')
                AND m_asset.id = asset_id) bicycle_capacity_avaliable
FROM
    m_asset
WHERE
    m_asset.status = 'A'`
	rows, err := s.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var rate sql.NullString
	var listAssetsLocation []*models.AssetLocation
	for rows.Next() {
		asset := models.AssetLocation{}
		err := rows.Scan(&asset.ID, &asset.AssetName, &asset.Longitude, &asset.Latitude, &asset.Photo, &rate, &asset.CarCap, &asset.MotorCap, &asset.BicycleCap)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		log.Println(asset.AssetName)
		asset.Rating = rate.String
		listAssetsLocation = append(listAssetsLocation, &asset)
	}

	return listAssetsLocation, nil
}
