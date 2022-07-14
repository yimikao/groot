package groot

import "net/http"

// A route that can be matched against
type RouteEntry struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

// hold all routes that can be matched against
type Router struct {
	Routes []RouteEntry
}

// registers a route with the specified method and handler
func (rt *Router) Route(path, method string, handler http.HandlerFunc) {
	e := RouteEntry{
		Path:    path,
		Method:  method,
		Handler: handler,
	}
	rt.Routes = append(rt.Routes, e)
}
