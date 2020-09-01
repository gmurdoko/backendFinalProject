package constanta

const (
	GETALLASSETSPERPROVIDER = `SELECT ma.asset_name,ma.asset_area, ma.longitude, ma.latitude, ma.car_capacity, ma.motorcycle_capacity,
								ma.bicycle_capacity, ma.photo,mw.saldo,ma.status
								FROM m_asset as ma join m_wallet as mw on ma.id_wallet = mw.id  where ma.provider_id = ?;`
	DELETEUSER     = `UPDATE m_user_account SET status ="I" where id=?;`
	DELETEPROVIDER = `UPDATE m_provider_account SET status="A" where id=?;`
	DELETEASSET    = `UPDATE m_asset SET status = "I" where id=?;`
	DELETECOMMENT  = `UPDATE m_review SET status = "I" where id=?;`
	APPROVEASSETS  = `UPDATE m_asset SET status="A" where id =?;`
)
