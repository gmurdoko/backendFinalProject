package ticket_repositories

import (
	"database/sql"
	"finalproject/main/master/models"
	"finalproject/utils/constant"
	"github.com/google/uuid"
	"log"
)

type TicketRepoImpl struct{
	db *sql.DB
}

func InitTicketRepoImpl(mydb *sql.DB) TicketRepo {
	return &TicketRepoImpl{db: mydb}
}

func (t *TicketRepoImpl) CreateNewTicket(ticket *models.Ticket) (*models.Ticket, error) {
	query := constant.CREATE_NEW_TICKET
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

func (t *TicketRepoImpl) UpdateTicketStatusActive(ticketID string) error {
	query := constant.UPDATE_TICKET_START_PARKING
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

func (t *TicketRepoImpl) UpdateTicketStatusInactive(ticketID string) error {
	query := constant.UPDATE_TICKET_FINISH_PARKING
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
