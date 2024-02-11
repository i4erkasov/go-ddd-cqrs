package handler

import "github.com/i4erkasov/go-ddd-cqrs/internal/application"

type HttpHandler struct {
	app *application.App
}

func NewHttpHandler(opts ...func(*HttpHandler)) *HttpHandler {
	h := &HttpHandler{}
	for _, opt := range opts {
		opt(h)
	}
	return h
}
