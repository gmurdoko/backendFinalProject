package ticket_usecases

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/user/ticketrepository"
)

type TicketUsecaseImpl struct {
	ticketRepo ticketrepository.TicketRepository
}

func InitTicketUsecase(repo ticketrepository.TicketRepository) TicketUsecase {
	return &TicketUsecaseImpl{ticketRepo: repo}
}

func (t *TicketUsecaseImpl) CreateNewTicket(ticket *models.Ticket) (*models.Ticket, error) {
	data, err := t.ticketRepo.CreateNewTicket(ticket)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (t *TicketUsecaseImpl) UpdateTicketStatusActive(ticketID string) (string, error) {
	err := t.ticketRepo.UpdateTicketStatusActive(ticketID)
	if err != nil {
		return "Failed changing ticket status", err
	}
	return "Success change ticket status: Active", nil
}

func (t *TicketUsecaseImpl) UpdateTicketStatusInactive(ticketID string) (string, error) {
	err := t.ticketRepo.UpdateTicketStatusInactive(ticketID)
	if err != nil {
		return "Failed changing ticket status", err
	}
	return "Success change ticket status: Inactive", nil
}
