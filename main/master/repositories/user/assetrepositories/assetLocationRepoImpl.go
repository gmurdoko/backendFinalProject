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
	query := `SELECT ma.id,ma.asset_name, ma.longitude, ma.latitude, ma.photo,
	ma.car_capacity-(
	select count(m_ticket.id) from m_ticket where vehicle_id="209f6e05-eb5a-11ea-86a5-b4a9fc958140" and status="A")
	as car_capacity_available,
	ma.motorcycle_capacity-(
	select count(m_ticket.id) from m_ticket where vehicle_id="221be282-eb5a-11ea-86a5-b4a9fc958140" and status="A")
	as motorcycle_capacity_available,
	ma.bicycle_capacity-(
	select count(m_ticket.id) from m_ticket where vehicle_id="22be08ef-eb5a-11ea-86a5-b4a9fc958140" and status="A")
	as bicycle_capacity_available
	from m_asset ma join m_ticket mt  where ma.id = ? group by id;`
	row := s.db.QueryRow(query, id)
	var asset = models.AssetLocation{}
	err := row.Scan(&asset.ID, &asset.AssetName, &asset.Longitude, &asset.Latitude, &asset.Photo, &asset.CarCap, &asset.MotorCap, &asset.MotorCap)
	if err != nil {
		return nil, err
	}
	return &asset, nil
}

func (s *AssetsLocationRepoImpl) ReadAssetsLocation() ([]*models.AssetLocation, error) {
	query := `SELECT ma.id,ma.asset_name, ma.longitude, ma.latitude, ma.photo,
	ma.car_capacity-(
	select count(m_ticket.id) from m_ticket where vehicle_id="209f6e05-eb5a-11ea-86a5-b4a9fc958140" and status="A")
	as car_capacity_available,
	ma.motorcycle_capacity-(
	select count(m_ticket.id) from m_ticket where vehicle_id="221be282-eb5a-11ea-86a5-b4a9fc958140" and status="A")
	as motorcycle_capacity_available,
	ma.bicycle_capacity-(
	select count(m_ticket.id) from m_ticket where vehicle_id="22be08ef-eb5a-11ea-86a5-b4a9fc958140" and status="A")
	as bicycle_capacity_available
	from m_asset ma join m_ticket mt group by id;`
	rows, err := s.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	println("masuk")
	var listAssetsLocation []*models.AssetLocation
	for rows.Next() {
		asset := models.AssetLocation{}
		err := rows.Scan(&asset.ID, &asset.AssetName, &asset.Longitude, &asset.Latitude, &asset.Photo, &asset.CarCap, &asset.MotorCap, &asset.MotorCap)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		log.Println(asset.AssetName)
		listAssetsLocation = append(listAssetsLocation, &asset)
	}
	return listAssetsLocation, nil
}
