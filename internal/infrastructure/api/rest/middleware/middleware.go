package middleware

type HttpMiddleware struct {
}

func NewHttpMiddleware(opts ...func(*HttpMiddleware)) *HttpMiddleware {
	m := &HttpMiddleware{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}
