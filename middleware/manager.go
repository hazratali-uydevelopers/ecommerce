package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	allMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		allMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) With(next http.Handler, Middlewares ...Middleware) http.Handler {
	n := next

	for _, middleware := range Middlewares {
		n = middleware(n)
	}
	
	for _, allMiddleware := range mngr.allMiddlewares {
		n = allMiddleware(n)
	}

	return n
}

func (mngr *Manager) Use(Middlewares ...Middleware) *Manager {
	mngr.allMiddlewares = append(mngr.allMiddlewares, Middlewares...)
	return mngr
}

func (mngr *Manager) Wrap(next http.Handler, Middlewares ...Middleware) http.Handler {
	n := next

	for _, middleware := range Middlewares {
		n = middleware(n)
	}

	return n
}