package ticketusecase

import (
	"finalproject/main/master/model"
)

//TicketUsecase usecase interface
type TicketUsecase interface {
	DeleteTicket(id string) error
	GetHistoryTicketByID(offset, limit, id string) ([]*model.TicketView, *int, error)
}
