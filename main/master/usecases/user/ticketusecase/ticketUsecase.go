package ticketusecase

import (
	"finalproject/main/master/models"
)

//TicketUsecase usecase interface
type TicketUsecase interface {
	DeleteTicket(id string) error
	GetHistoryTicketByID(offset, limit, id string) ([]*models.TicketView, *int, error)
}
