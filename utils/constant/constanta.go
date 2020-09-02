package constant

const (
	// Searching for parking lot
	READ_LOCATION_ASSET = `SELECT asset_name, longitude, latitude FROM m_asset WHERE status="A";`

	// Ticket reservation
	CREATE_NEW_TICKET   = `INSERT INTO m_ticket (id, user_id, asset_id, fee_id, vehicle_id, license_plate, book_at, status) VALUES (?,?,?,?,?,?,?);`
	UPDATE_TICKET_START_PARKING = `UPDATE m_ticket SET status="A", start_at=NOW() WHERE id=?;`
	UPDATE_TICKET_FINISH_PARKING = `UPDATE m_ticket SET status="I", finished_at=NOW() WHERE id=?;`

	// Asset capacity change due to parking spot booking by car/motor/bike
	UPDATE_CAPACITY_CAR_START = `UPDATE m_asset SET car_capacity=car_capacity-1 WHERE id=?;`
	UPDATE_CAPACITY_MOTOR_START = `UPDATE m_asset SET motorcycle_capacity=motorcycle_capacity-1 WHERE id=?;`
	UPDATE_CAPACITY_BIKE_START = `UPDATE m_asset SET bicycle_capacity=bicycle_capacity-1 WHERE id=?;`
	UPDATE_CAPACITY_CAR_FINISH = `UPDATE m_asset SET car_capacity=car_capacity+1 WHERE id=?;`
	UPDATE_CAPACITY_MOTOR_FINISH = `UPDATE m_asset SET motorcycle_capacity=motorcycle_capacity+1 WHERE id=?;`
	UPDATE_CAPACITY_BIKE_FINISH = `UPDATE m_asset SET bicycle_capacity=bicycle_capacity+1 WHERE id=?;`

	// User review asset
	CREATE_RATING_COMMENT = `INSERT INTO m_review (id, user_id, asset_id, rating, comment, created_at, status) VALUES (uuid(), ?,?,?,?,NOW(),"A");`
	)
