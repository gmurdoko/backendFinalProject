package ticket_usecases

import (
	"finalproject/main/master/models"
)

type TicketUsecase interface {
	CreateNewTicket(ticket *models.Ticket) (*models.Ticket, error)
	UpdateTicketStatusActive(ticketID string) (string, error)
	UpdateTicketStatusInactive(ticketID string) (string, error)
}
