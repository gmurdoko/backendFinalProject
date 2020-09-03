package ticket_repositories

import "finalproject/main/master/models"

type TicketRepo interface {
	CreateNewTicket(ticket *models.Ticket) (*models.Ticket, error)
	UpdateTicketStatusActive(ticketID string) error
	UpdateTicketStatusInactive(ticketID string) error
}



