package ticketrepository

import (
	"database/sql"
)

//TicketRepositoryImpl is for init Repository
type TicketRepositoryImpl struct {
	db *sql.DB
}

//Delete si for deleting ticket history for user
func (s TicketRepositoryImpl) Delete(id string) error {
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

//InitTicketRepositoryImpl is init gate for repository
func InitTicketRepositoryImpl(db *sql.DB) TicketRepository {
	return &TicketRepositoryImpl{db}
}
