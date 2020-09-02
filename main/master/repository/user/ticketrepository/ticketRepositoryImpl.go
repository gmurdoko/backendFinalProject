package ticketrepository

import (
	"database/sql"
	"finalproject/main/master/model"
	"fmt"
	"log"
)

//ticketRepositoryImpl is for init Repository
type ticketRepositoryImpl struct {
	db *sql.DB
}

//Delete si for deleting ticket history for user
func (s ticketRepositoryImpl) Delete(id string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	query := "UPDATE m_ticket SET status = 'D' WHERE id = ?;"
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

//SelectHistoryTicketByUserID history user payment
func (s ticketRepositoryImpl) SelectHistoryTicketByUserID(offset, limit, id string) ([]*model.TicketView, *int, error) {
	fmt.Println("REPO 1", offset, limit, id)
	queryIn := fmt.Sprintf(`SELECT mt.id, mua.username, ma.asset_name, mt.license_plate, mf.fee, ceil(time_to_sec(timediff(mt.finished_at, mt.start_at))/3600), (ceil(time_to_sec(timediff(mt.finished_at, mt.start_at))/3600) * mf.fee), mt.book_at, mt.start_at, mt.finished_at FROM m_ticket mt JOIN m_user_account mua ON mt.user_id = mua.id JOIN m_asset ma ON mt.asset_id = ma.id JOIN m_fee mf ON mt.fee_id = mf.id WHERE mt.status = 'I' AND mua.id = ? ORDER BY mt.finished_at DESC LIMIT %s , %s;`, offset, limit)
	fmt.Println(queryIn)
	data, err := s.db.Query(queryIn, id)
	if err != nil {
		return nil, nil, err
	}
	// fmt.Println("Repo", data)
	defer data.Close()
	var result = []*model.TicketView{}
	fmt.Println("SAMPAI SINI")
	for data.Next() {
		var ticketView = model.TicketView{}
		var err = data.Scan(&ticketView.ID, &ticketView.Username, &ticketView.AssetName, &ticketView.LicensePlate, &ticketView.BasedFee, &ticketView.ParkingDurationHour, &ticketView.PayFee, &ticketView.BookAt, &ticketView.StartAt, &ticketView.FinishedAt)
		fmt.Println(ticketView.ID, ticketView.Username)
		if err != nil {
			log.Println(err)
			return nil, nil, err
		}
		result = append(result, &ticketView)
	}
	fmt.Println("HERE WE GO")
	if err = data.Err(); err != nil {
		return nil, nil, err
	}
	var totalField = new(int)
	err = s.db.QueryRow(`select COUNT(*) FROM m_ticket mt JOIN m_user_account mua ON mt.user_id = mua.id JOIN m_asset ma ON mt.asset_id = ma.id JOIN m_fee mf ON mt.fee_id = mf.id WHERE mt.status = 'I' AND mua.id = ?;`, id).Scan(&totalField)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("total Field", *totalField)
	return result, totalField, nil
}

//InitTicketRepositoryImpl is init gate for repository
func InitTicketRepositoryImpl(db *sql.DB) TicketRepository {
	return &ticketRepositoryImpl{db}
}
