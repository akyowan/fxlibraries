package http

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

//Router
type Router struct {
	*mux.Router
}

type Route struct {
	*mux.Route
}

var r *Router

func init() {
	r = &Router{mux.NewRouter()}
}

//GetRouter
func GetRouter() *Router {
	return r
}

//RouteHandlerFunc
func RouteHandlerFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return &Route{r.HandleFunc(path, f)}
}

//ListenAndServe, This function will blocks
func ListenAndServe(addr string) error {
	http.Handle("/", handlers.HTTPMethodOverrideHandler(r))
	return http.ListenAndServe(addr, nil)
}
