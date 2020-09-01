package ticketrepository

//TicketRepository interface for ticket
type TicketRepository interface {
	Delete(id string) error
}
