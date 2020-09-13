package utils

const (
	SELECT_NEW_USER       = `SELECT * FROM m_user_account WHERE id=?`
	SELECT_NEW_PROVIDER   = `SELECT * FROM m_provider_account WHERE id=?`
	SELECT_USER_LOGIN     = `SELECT * FROM m_user_account WHERE username=? OR email=?`
	SELECT_PROVIDER_LOGIN = `SELECT * FROM m_provider_account WHERE username=? OR email=?`
	SELECT_PROVIDER_EXIST = `SELECT username, email FROM m_provider_account WHERE username=? OR email=?`
	SELECT_USER_EXIST     = `SELECT username, email FROM m_user_account WHERE username=? OR email=?`
	INSERT_USER_ACCOUNT   = `INSERT INTO m_user_account (id,id_wallet,username,password,email,fullname,phone_number,created_at) 
	VALUES (?,?,?,?,?,?,?,?)`
	INSERT_PROVIDER_ACCOUNT = `INSERT INTO m_provider_account (id,username,password,email,fullname,phone_number,created_at) 
	VALUES (?,?,?,?,?,?,?)`
	INSERT_ASSET = `INSERT INTO m_asset (id,id_wallet,provider_id,asset_name,asset_area,longitude,latitude,car_capacity,motorcycle_capacity,bicycle_capacity,
		photo, created_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`
	INSERT_WALLET             = `INSERT INTO m_wallet (id) VALUES(?)`
	SELECT_USER_SALDO         = `SELECT m_user_account.id, saldo from m_wallet JOIN m_user_account ON m_wallet.id = m_user_account.id_wallet WHERE m_user_account.id=?`
	UPDATE_USER_SALDO_TOPUP   = `UPDATE m_wallet SET debit=?, saldo=saldo+debit, edited_at=? WHERE id=?`
	SELECT_PHOTO_USER         = `SELECT photo FROM m_user_account WHERE id=?`
	UPDATE_PHOTO_USER         = `UPDATE m_user_account SET photo=?,edited_at=? WHERE id=?`
	DELETE_PHOTO_USER         = `UPDATE m_user_account SET photo =' ',deleted_at=? WHERE id=?`
	UPDATE_DATA_USER          = `UPDATE m_user_account SET address =?, borndate=?,edited_at=? WHERE id =?`
	SELECT_UPDATED_SALDO_USER = "SELECT saldo FROM m_wallet WHERE id=?"
	SELECT_PROVIDER_SALDO     = `SELECT SUM(mw.saldo) FROM m_asset AS ma JOIN m_wallet mw ON ma.id_wallet = mw.id JOIN 
	m_provider_account mpa on ma.provider_id = mpa.id WHERE mpa.id = ?`
	DATE_FORMAT     = `2006-01-02 15:04:05`
	GET_USER_TICKET = `SELECT 
	m_ticket.id, user_id,
    m_asset.asset_name,
    m_vehicle_type.vehicle_type,
    license_plate,
	book_at, m_ticket.status,
	m_ticket.asset_id,
	m_ticket.fee_id,
	m_ticket.vehicle_id
FROM
    m_ticket
        JOIN
    m_asset ON m_ticket.asset_id = m_asset.id
        JOIN
    m_vehicle_type ON m_vehicle_type.id = m_ticket.vehicle_id
WHERE
    m_ticket.user_id = ?
		AND (m_ticket.status = "A" OR m_ticket.status="B")`
	SELECT_USER_TICKET  = `select id,user_id,asset_id,fee_id,vehicle_id,license_plate,book_at from m_ticket where user_id=?`
	SELECT_ASSET_REVIEW = `SELECT 
    m_provider_account.id,m_review.rating, m_review.comment, m_asset.asset_name
FROM
    m_review
        JOIN
    m_asset ON m_review.asset_id = m_asset.id
        JOIN
    m_provider_account ON m_asset.provider_id = m_provider_account.id WHERE m_provider_account.id=?;`
)
