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
	query := `SELECT m_asset.id,m_asset.asset_name, m_asset.longitude, m_asset.latitude, m_asset.photo,
	m_asset.car_capacity-(
	select count(m_ticket.id) from m_ticket where vehicle_id="1" and status="A" or status="B")
	as car_capacity_available,
	m_asset.motorcycle_capacity-(
	select count(m_ticket.id) from m_ticket where vehicle_id="2" and status="A" or status="B")
	as motorcycle_capacity_available,
	m_asset.bicycle_capacity-(
	select count(m_ticket.id) from m_ticket where vehicle_id="3" and status="A" or status="B")
	as bicycle_capacity_available from m_asset where id=?`
	row := s.db.QueryRow(query, id)
	var asset = models.AssetLocation{}
	err := row.Scan(&asset.ID, &asset.AssetName, &asset.Longitude, &asset.Latitude, &asset.Photo, &asset.CarCap, &asset.MotorCap, &asset.BicycleCap)
	if err != nil {
		return nil, err
	}
	return &asset, nil
}

func (s *AssetsLocationRepoImpl) ReadAssetsLocation() ([]*models.AssetLocation, error) {
	query := `SELECT m_asset.id,m_asset.asset_name, m_asset.longitude, m_asset.latitude, m_asset.photo,
	m_asset.car_capacity-(
	select count(m_ticket.id) from m_ticket where vehicle_id="1" and status="A" or status="B")
	as car_capacity_available,
	m_asset.motorcycle_capacity-(
	select count(m_ticket.id) from m_ticket where vehicle_id="2" and status="A" or status="B")
	as motorcycle_capacity_available,
	m_asset.bicycle_capacity-(
	select count(m_ticket.id) from m_ticket where vehicle_id="3" and status="A" or status="B")
	as bicycle_capacity_available from m_asset`
	rows, err := s.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	println("masuk")
	var listAssetsLocation []*models.AssetLocation
	for rows.Next() {
		asset := models.AssetLocation{}
		err := rows.Scan(&asset.ID, &asset.AssetName, &asset.Longitude, &asset.Latitude, &asset.Photo, &asset.CarCap, &asset.MotorCap, &asset.BicycleCap)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		log.Println(asset.AssetName)
		listAssetsLocation = append(listAssetsLocation, &asset)
	}
	return listAssetsLocation, nil
}
