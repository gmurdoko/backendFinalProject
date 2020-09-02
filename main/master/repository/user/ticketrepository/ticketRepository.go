package ticketrepository

import (
	"finalproject/main/master/model"
)

//TicketRepository interface for ticket
type TicketRepository interface {
	Delete(id string) error
	SelectHistoryTicketByUserID(offset, limit, id string) ([]*model.TicketView, *int, error)
}
