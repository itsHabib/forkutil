package art

import (
	"fmt"
	"net/http"
)

// RouterFunc defines a callback router with optional content in the
// second parameter
type RouterFunc func(*http.Response, interface{}) error

// CBRouter maps status codes to specific callback functions
type CBRouter struct {
	Routers       map[int]RouterFunc
	DefaultRouter RouterFunc
}

// NewRouter returns a new router
func NewRouter() *CBRouter {
	return &CBRouter{
		Routers: make(map[int]RouterFunc),
		DefaultRouter: func(resp *http.Response, _ interface{}) error {
			return fmt.Errorf("From: %s received unkown status: %d",
				resp.Request.URL.String(), resp.StatusCode)
		},
	}
}

// RegisterFunc registers a function with a status code
func (r *CBRouter) RegisterFunc(status int, fn RouterFunc) {
	r.Routers[status] = fn
}

// CallFunc calls a registered func in routers map based on status code
func (r *CBRouter) CallFunc(resp *http.Response, content interface{}) error {
	fn, ok := r.Routers[resp.StatusCode]
	if !ok {
		fn = r.DefaultRouter
	}
	return fn(resp, content)
}
