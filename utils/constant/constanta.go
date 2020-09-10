package constanta

const (

	// Get all data per tables for admin frontend
	GET_ALL_USERS     = `SELECT id, id_wallet, username, email, fullname, borndate, phone_number, address, created_at, status FROM m_user_account;`
	GET_ALL_PROVIDERS = `SELECT id, username, email, fullname, borndate, phone_number, address, created_at, status FROM m_provider_account;`
	GET_ALL_ASSETS    = `SELECT id, id_wallet, provider_id, asset_name, asset_area, longitude, latitude, car_capacity, motorcycle_capacity, bicycle_capacity, created_at, status FROM m_asset;`
	GET_ALL_REVIEWS   = `SELECT id, user_id, asset_id, rating, comment, created_at, status FROM m_review;`

	GETALLASSETSPERPROVIDER = `SELECT ma.id, ma.asset_name,ma.asset_area, ma.longitude, ma.latitude, ma.car_capacity, ma.motorcycle_capacity,
								ma.bicycle_capacity, ma.photo,mw.saldo,ma.status
								FROM m_asset as ma join m_wallet as mw on ma.id_wallet = mw.id  where ma.provider_id = ?;`
	DELETEUSER                = `UPDATE m_user_account SET status ="I" where id=?;`
	DELETEPROVIDER            = `UPDATE m_provider_account SET status="A" where id=?;`
	DELETEASSET               = `UPDATE m_asset SET status = "I" where id=?;`
	DELETECOMMENT             = `UPDATE m_review SET status = "I" where id=?;`
	APPROVEASSETS             = `UPDATE m_asset SET status="A" where id =?;`
	PROVIDERASSETSDAILYREPORT = `SELECT 
									ma.asset_name,
										DATE(mt.finished_at) date,
										COUNT(mt.id) total_parked ,
										SUM(CEIL(TIME_TO_SEC(TIMEDIFF(mt.finished_at, mt.start_at)) / 3600) * mf.fee) total_rev
									FROM
										m_asset AS ma
											JOIN
										m_wallet mw ON ma.id_wallet = mw.id
											JOIN
										m_ticket mt ON mt.asset_id = ma.id
											JOIN
										m_fee mf ON mf.id = mt.fee_id
										WHERE ma.id = ? AND mt.status = "I" AND MONTH(NOW())-MONTH(mt.finished_at) <= 2 GROUP BY date  ;`
	PROVIDERASSETSMONTHLYREPORT = `SELECT 
										ma.asset_name,
										MONTH(mt.finished_at) AS months,
										COUNT(mt.id) total_parked ,
										SUM(CEIL(TIME_TO_SEC(TIMEDIFF(mt.finished_at, mt.start_at)) / 3600) * mf.fee) total_rev
									FROM
										m_asset AS ma
											JOIN
										m_wallet mw ON ma.id_wallet = mw.id
											JOIN
										m_ticket mt ON mt.asset_id = ma.id
											JOIN
										m_fee mf ON mf.id = mt.fee_id
										WHERE ma.id = ? AND mt.status = "I" AND MONTH(NOW())-MONTH(mt.finished_at) <= 2  GROUP BY months ;`
	ADMINASSETSDAILYREPORT = `SELECT 
							ma.asset_name,
								DATE(mt.finished_at) date,
								COUNT(mt.id) total_parked ,
								SUM(CEIL(TIME_TO_SEC(TIMEDIFF(mt.finished_at, mt.start_at)) / 3600) * mf.fee) total_rev
							FROM
								m_asset AS ma
									JOIN
								m_wallet mw ON ma.id_wallet = mw.id
									JOIN
								m_ticket mt ON mt.asset_id = ma.id
									JOIN
								m_fee mf ON mf.id = mt.fee_id
								WHERE ma.id = ? AND mt.status = "I" AND MONTH(NOW())-MONTH(mt.finished_at) <= 2 GROUP BY date  ;`
	ADMINASSETSMONTHLYREPORT = `SELECT 
								ma.asset_name,
								MONTH(mt.finished_at) AS months,
								COUNT(mt.id) total_parked ,
								SUM(CEIL(TIME_TO_SEC(TIMEDIFF(mt.finished_at, mt.start_at)) / 3600) * mf.fee) total_rev
							FROM
								m_asset AS ma
									JOIN
								m_wallet mw ON ma.id_wallet = mw.id
									JOIN
								m_ticket mt ON mt.asset_id = ma.id
									JOIN
								m_fee mf ON mf.id = mt.fee_id
								WHERE ma.id = ? AND mt.status = "I" AND MONTH(NOW())-MONTH(mt.finished_at) <= 2  GROUP BY months ;`
	READ_LOCATION_ASSET = `SELECT id, asset_name, longitude, latitude FROM m_asset WHERE status="A";`

	// Ticket reservation
	CREATE_NEW_TICKET            = `INSERT INTO m_ticket (id, user_id, asset_id, fee_id, vehicle_id, license_plate, status) VALUES (?,?,?,?,?,?,"B");`
	UPDATE_TICKET_START_PARKING  = `UPDATE m_ticket SET status="A", start_at=NOW() WHERE id=?;`
	UPDATE_TICKET_FINISH_PARKING = `UPDATE m_ticket SET status="I", finished_at=NOW() WHERE id=?;`

	// Asset capacity change due to parking spot booked by car/motor/bike
	READ_CURRENT_CAPACITY = `select ma.car_capacity, ma.car_capacity-(
							select count(m_ticket.id) from m_ticket where vehicle_id="209f6e05-eb5a-11ea-86a5-b4a9fc958140" and status="A")
							as car_capacity_available,
							ma.motorcycle_capacity, ma.motorcycle_capacity-(
							select count(m_ticket.id) from m_ticket where vehicle_id="221be282-eb5a-11ea-86a5-b4a9fc958140" and status="A")
							as motorcycle_capacity_available,
							ma.bicycle_capacity, ma.bicycle_capacity-(
							select count(m_ticket.id) from m_ticket where vehicle_id="22be08ef-eb5a-11ea-86a5-b4a9fc958140" and status="A")
							as bicycle_capacity_available
							from m_asset ma join m_ticket mt on ma.id = mt.asset_id GROUP BY ?;`

	// User review asset
	CREATE_RATING_COMMENT = `INSERT INTO m_review (id, user_id, asset_id, rating, comment, created_at, status) VALUES (?,?,?,?,?,NOW(),"A");`
)
