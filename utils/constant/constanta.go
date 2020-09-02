package constanta

const (
	GETALLASSETSPERPROVIDER = `SELECT ma.asset_name,ma.asset_area, ma.longitude, ma.latitude, ma.car_capacity, ma.motorcycle_capacity,
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
)
