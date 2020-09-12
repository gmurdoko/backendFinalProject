package ticketrepository

import (
	"database/sql"
	"finalproject/main/master/models"
	constanta "finalproject/utils/constant"
	"fmt"
	"log"

	"github.com/google/uuid"
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
func (s ticketRepositoryImpl) SelectHistoryTicketByUserID(offset, limit, id string) ([]*models.TicketView, *int, error) {

	queryIn := fmt.Sprintf(`SELECT mt.id, mua.username, ma.asset_name, mt.license_plate, mf.fee, ceil(time_to_sec(timediff(mt.finished_at, mt.start_at))/3600), (ceil(time_to_sec(timediff(mt.finished_at, mt.start_at))/3600) * mf.fee), mt.book_at, mt.start_at, mt.finished_at FROM m_ticket mt JOIN m_user_account mua ON mt.user_id = mua.id JOIN m_asset ma ON mt.asset_id = ma.id JOIN m_fee mf ON mt.fee_id = mf.id WHERE mt.status = 'I' AND mua.id = ? ORDER BY mt.finished_at DESC LIMIT %s , %s;`, offset, limit)
	data, err := s.db.Query(queryIn, id)
	if err != nil {
		return nil, nil, err
	}
	// fmt.Println("Repo", data)
	defer data.Close()
	var result = []*models.TicketView{}

	for data.Next() {
		var ticketView = models.TicketView{}
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

func (t *ticketRepositoryImpl) CreateNewTicket(ticket *models.Ticket) (*models.Ticket, error) {
	query := constanta.CREATE_NEW_TICKET
	ticket.ID = uuid.New().String()
	tx, err := t.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec(query, ticket.ID, ticket.UserID, ticket.AssetID,
		ticket.FeeID, ticket.VehicleID, ticket.LicensePlate)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ticket, tx.Commit()
}

func (t *ticketRepositoryImpl) UpdateTicketStatusActive(ticketID string) error {
	query := constanta.UPDATE_TICKET_START_PARKING
	tx, err := t.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = tx.Exec(query, ticketID)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	return tx.Commit()
}

func (t *ticketRepositoryImpl) UpdateTicketStatusInactive(ticketID string) error {
	query := constanta.UPDATE_TICKET_FINISH_PARKING
	tx, err := t.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = tx.Exec(query, ticketID)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	return tx.Commit()
}

func (t *ticketRepositoryImpl) SelectTicketViewByID(id string) (*models.TicketView, error) {
	query := "SELECT mt.id, mua.username, ma.asset_name, mt.license_plate, mf.fee, ceil(time_to_sec(timediff(mt.finished_at, mt.start_at))/3600), (ceil(time_to_sec(timediff(mt.finished_at, mt.start_at))/3600) * mf.fee), mt.book_at, mt.start_at, mt.finished_at FROM m_ticket mt JOIN m_user_account mua ON mt.user_id = mua.id JOIN m_asset ma ON mt.asset_id = ma.id JOIN m_fee mf ON mt.fee_id = mf.id WHERE mt.status = 'A' AND mt.id = ?;"
	ticketView := new(models.TicketView)
	var bookAt, StartAt, FinishedAt sql.NullString
	err := t.db.QueryRow(query, id).Scan(&ticketView.ID, &ticketView.Username, &ticketView.AssetName, &ticketView.LicensePlate,
		&ticketView.BasedFee, &ticketView.ParkingDurationHour, &ticketView.PayFee, &bookAt, &StartAt, &FinishedAt)
	ticketView.BookAt = bookAt.String
	ticketView.StartAt = StartAt.String
	ticketView.FinishedAt = FinishedAt.String
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return ticketView, nil
}

//InitTicketRepositoryImpl is init gate for repository
func InitTicketRepositoryImpl(db *sql.DB) TicketRepository {
	return &ticketRepositoryImpl{db}
}
