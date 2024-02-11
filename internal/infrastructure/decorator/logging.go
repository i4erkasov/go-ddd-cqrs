package decorator

import (
	"context"
	"fmt"

	"github.com/i4erkasov/go-ddd-cqrs/internal/application/command"
	"github.com/i4erkasov/go-ddd-cqrs/internal/application/query"
	"github.com/juju/errors"
	"go.uber.org/zap"
)

type loggingCommand[C any] struct {
	command command.Command[C]
	logger  *zap.Logger
}

func ApplyLoggingCommand[C any](handler command.Command[C], logger *zap.Logger) command.Command[C] {
	return &loggingCommand[C]{
		command: handler,
		logger:  logger,
	}
}

func (l *loggingCommand[C]) Execute(ctx context.Context, cmd C) (err error) {
	logger := l.logger.With([]zap.Field{
		zap.String("command", GetActionName(cmd)),
		zap.String("command_body", fmt.Sprintf("%#v", cmd)),
	}...)

	logger.Debug("Executing command")
	defer func() {
		if err == nil {
			logger.Info("Command executed successfully")
		} else {
			logger.Error("Failed to execute command", zap.Error(errors.Unwrap(err)))
		}
	}()

	err = l.command.Execute(ctx, cmd)
	return err
}

type loggingQuery[Q any, R any] struct {
	query  query.Query[Q, R]
	logger *zap.Logger
}

func ApplyLoggingQuery[Q any, R any](query query.Query[Q, R], logger *zap.Logger) query.Query[Q, R] {
	return &loggingQuery[Q, R]{
		query:  query,
		logger: logger,
	}
}

func (l *loggingQuery[Q, R]) Get(ctx context.Context, query Q) (result R, err error) {
	logger := l.logger.With([]zap.Field{
		zap.String("query", GetActionName(query)),
		zap.String("query_body", fmt.Sprintf("%#v", query)),
	}...)

	logger.Debug("Executing query")
	defer func() {
		if err == nil {
			logger.Info("Query executed successfully")
		} else {
			logger.Error("Failed to execute query", zap.Error(errors.Unwrap(err)))
		}
	}()

	result, err = l.query.Get(ctx, query)
	return result, err
}
