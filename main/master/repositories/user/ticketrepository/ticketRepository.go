package ticketrepository

import (
	"finalproject/main/master/models"
)

//TicketRepository interface for ticket
type TicketRepository interface {
	Delete(id string) error
	SelectHistoryTicketByUserID(offset, limit, id string) ([]*models.TicketView, *int, error)
	CreateNewTicket(ticket *models.Ticket) (*models.Ticket, error)
	UpdateTicketStatusActive(ticketID string) error
	UpdateTicketStatusInactive(ticketID string) error
	SelectTicketViewByID(id string) (*models.TicketView, error)
}
