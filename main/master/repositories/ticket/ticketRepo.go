package ticket

import "finalproject/main/master/models"

type TicketRepo interface {
	CreateNewTicket(ticket *models.Ticket) (*models.Ticket, error)
	UpdateTicketStatusActive(ticket *models.Ticket, ticketID string) (*models.Ticket, error)
	UpdateTicketStatusInactive(ticket *models.Ticket, ticketID string) (*models.Ticket, error)
}


