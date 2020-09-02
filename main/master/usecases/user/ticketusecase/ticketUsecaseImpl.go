package ticketusecase

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/user/ticketrepository"
	"fmt"
)

//ticketUsecaseImpl app
type ticketUsecaseImpl struct {
	ticketRepository ticketrepository.TicketRepository
}

//DeleteTicket app
func (s ticketUsecaseImpl) DeleteTicket(id string) error {
	err := s.ticketRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

//GetHistoryTicketByID app
func (s ticketUsecaseImpl) GetHistoryTicketByID(offset, limit, id string) ([]*models.TicketView, *int, error) {
	println("INI USECASE", offset, limit, id)
	ticketView, totalField, err := s.ticketRepository.SelectHistoryTicketByUserID(offset, limit, id)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("useCase", *totalField)
	return ticketView, totalField, nil
}

//InitTicketUsecaseImpl app
func InitTicketUsecaseImpl(ticketRepository ticketrepository.TicketRepository) TicketUsecase {
	return &ticketUsecaseImpl{ticketRepository}
}
