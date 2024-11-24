package tickets

import "context"

type tickets interface {
	CreateTickets(ctx context.Context) (tickets, error)
}
