package constanta

const (
	GETALLASSETSPERPROVIDER = `SELECT ma.asset_name,ma.asset_area, ma.longitude, ma.latitude, ma.car_capacity, ma.motorcycle_capacity,
								ma.bicycle_capacity, ma.photo,mw.saldo,ma.status
								FROM m_asset as ma join m_wallet as mw on ma.id_wallet = mw.id  where ma.provider_id = ?;`
)
