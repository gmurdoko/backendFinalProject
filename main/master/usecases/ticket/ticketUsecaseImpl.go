package usecases

import (
	"finalproject/main/master/models"
	repositories "finalproject/main/master/repositories/ticket"
)

type TicketUsecaseImpl struct {
	ticketRepo repositories.TicketRepo
}

func InitTicketUsecase(repo repositories.TicketRepo) TicketUsecase {
	return &TicketUsecaseImpl{ticketRepo: repo}
}

func (t *TicketUsecaseImpl) CreateNewTicket(ticket *models.Ticket) (*models.Ticket, error) {
	data, err := t.ticketRepo.CreateNewTicket(ticket)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (t *TicketUsecaseImpl) UpdateTicketStatusActive(ticket *models.Ticket, ticketID string) (*models.Ticket, error) {
	data, err := t.ticketRepo.UpdateTicketStatusActive(ticket, ticketID)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (t *TicketUsecaseImpl) UpdateTicketStatusInactive(ticket *models.Ticket, ticketID string) (*models.Ticket, error) {
	data, err := t.ticketRepo.UpdateTicketStatusInactive(ticket, ticketID)
	if err != nil {
		return nil, err
	}
	return data, nil
}
