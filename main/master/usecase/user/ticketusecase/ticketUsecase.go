package ticketusecase

//TicketUsecase usecase interface
type TicketUsecase interface {
	DeleteTicket(id string) error
}
