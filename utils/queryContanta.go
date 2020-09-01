package utils

const (
	INSERT_USER_ACCOUNT = `INSERT INTO m_user_account (id,id_wallet,username,password,email,fullname,phone_number,created_at) 
	VALUES (?,?,?,?,?,?,?,?)`
	INSERT_PROVIDER_ACCOUNT = `INSERT INTO m_provider_account (id,username,password,email,fullname,phone_number,created_at) 
	VALUES (?,?,?,?,?,?,?)`
	INSERT_ASSET = `INSERT INTO m_asset (id,id_wallet,provider_id,asset_name,asset_area,longitude,latitude,car_capacity,motorcycle_capacity,bicycle_capacity,
		created_at, photo) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`
	INSERT_WALLET           = `INSERT INTO m_wallet (id) VALUES(?)`
	SELECT_USER_SALDO       = `SELECT saldo from m_wallet JOIN m_user_account ON m_wallet.id = m_user_account.id_wallet WHERE m_user_account.id=?`
	UPDATE_USER_SALDO_TOPUP = `UPDATE m_wallet SET saldo=saldo+?, edited_at=? WHERE id=?`
	SELECT_PHOTO_USER       = `SELECT photo FROM m_user_account WHERE id=?`
	UPDATE_PHOTO_USER       = `UPDATE m_user_account SET photo=?,edited_at=? WHERE id=?`
	DELETE_PHOTO_USER       = `UPDATE m_user_account SET photo = null,deleted_at=? WHERE id=?`
	UPDATE_DATA_USER        = `UPDATE m_user_account SET address =?, borndate=?,edited_at=? WHERE id =?`
)
