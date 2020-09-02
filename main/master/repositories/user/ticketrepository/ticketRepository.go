package ticketrepository

import "finalproject/main/master/models"

//TicketRepository interface for ticket
type TicketRepository interface {
	Delete(id string) error
	SelectHistoryTicketByUserID(offset, limit, id string) ([]*models.TicketView, *int, error)
}
