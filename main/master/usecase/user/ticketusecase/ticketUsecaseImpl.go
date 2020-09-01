package ticketusecase

import "finalproject/main/master/repository/user/ticketrepository"

//TicketUsecaseImpl app
type TicketUsecaseImpl struct {
	ticketRepository ticketrepository.TicketRepository
}

//DeleteTicket app
func (s TicketUsecaseImpl) DeleteTicket(id string) error {
	err := s.ticketRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

//InitTicketUsecaseImpl app
func InitTicketUsecaseImpl(ticketRepository ticketrepository.TicketRepository) TicketUsecase {
	return &TicketUsecaseImpl{ticketRepository}
}
