package application

import (
	"github.com/i4erkasov/go-ddd-cqrs/internal/application/command"
	"github.com/i4erkasov/go-ddd-cqrs/internal/application/query"
)

type App struct {
	Commands *command.Commands
	Queries  *query.Queries
}

func New(
	commands *command.Commands,
	queries *query.Queries,
) *App {
	return &App{
		Commands: commands,
		Queries:  queries,
	}
}
