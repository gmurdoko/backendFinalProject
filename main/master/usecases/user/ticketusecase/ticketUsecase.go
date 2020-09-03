package ticketusecase

import (
	"finalproject/main/master/models"
)

//TicketUsecase usecase interface
type TicketUsecase interface {
	DeleteTicket(id string) error
	GetHistoryTicketByID(offset, limit, id string) ([]*models.TicketView, *int, error)
	CreateNewTicket(ticket *models.Ticket) (*models.Ticket, error)
	UpdateTicketStatusActive(ticketID string) (string, error)
	UpdateTicketStatusInactive(ticketID string) (string, error)
}
