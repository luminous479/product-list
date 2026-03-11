package middlewares

import (
	"net/http"
)

type Middleware func (http.Handler) http.Handler

type Manager struct{
	globalMiddleware []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddleware: []Middleware{},
	}
}
func (m *Manager) Use(middleware ...Middleware) {
	m.globalMiddleware = append(m.globalMiddleware, middleware...)
}
func (m *Manager) WrapMux(next http.Handler) http.Handler {
	for _, middleware := range m.globalMiddleware {
		next = middleware(next)
	}
	return next
}
