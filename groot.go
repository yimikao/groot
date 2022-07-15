package groot

import (
	"context"
	"net/http"
	"regexp"
)

type reqContext string

var (
	rp reqContext = "url_params"
)

// A route that can be matched against
type RouteEntry struct {
	Path    *regexp.Regexp
	Method  string
	Handler http.HandlerFunc
}

// hold all routes that can be matched against
type Router struct {
	Routes []RouteEntry
}

// registers a route with the specified method and handler
func (rt *Router) Route(path *regexp.Regexp, method string, handler http.HandlerFunc) {
	e := RouteEntry{
		Path:    path,
		Method:  method,
		Handler: handler,
	}
	rt.Routes = append(rt.Routes, e)
}

func (e *RouteEntry) Match(r *http.Request) map[string]string {
	match := e.Path.FindStringSubmatch(r.URL.Path)
	if match == nil {
		return nil
	}

	params := map[string]string{}
	paramNames := e.Path.SubexpNames()
	for i, pn := range match {
		params[paramNames[i]] = pn
	}

	return params

}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, e := range rt.Routes {
		params := e.Match(r)
		if params == nil {
			continue
		}

		ctx := context.WithValue(r.Context(), rp, params)
		e.Handler.ServeHTTP(w, r.WithContext(ctx))
		return
	}

}
