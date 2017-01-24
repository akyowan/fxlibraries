package http

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type HandleFunc func(*Request) *Response

type Router struct {
	*mux.Router
}

type Route struct {
	*mux.Route
}

type OriginValidator func(string) bool

var r *Router

var corsOpts []handlers.CORSOption

func init() {
	r = &Router{mux.NewRouter()}
}

func GetRouter() *Router {
	return r
}

func RouteHandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route {
	return &Route{r.HandleFunc(path, f)}
}

func AllowedMethods(methods []string) {
	corsOpts = append(corsOpts, handlers.AllowedMethods(methods))
}

func AllowedHeaders(headers []string) {
	corsOpts = append(corsOpts, handlers.AllowedHeaders(headers))
}

func AllowedOrigins(origins []string) {
	corsOpts = append(corsOpts, handlers.AllowedOrigins(origins))
}

func AllowedOriginValidator(fn OriginValidator) {
	corsOpts = append(corsOpts, handlers.AllowedOriginValidator(handlers.OriginValidator(fn)))
}

func ExposedHeaders(headers []string) {
	corsOpts = append(corsOpts, handlers.ExposedHeaders(headers))
}

func MaxAge(age int) {
	corsOpts = append(corsOpts, handlers.MaxAge(age))
}

func IgnoreOptions() {
	corsOpts = append(corsOpts, handlers.IgnoreOptions())
}

func AllowCredentials() {
	corsOpts = append(corsOpts, handlers.AllowCredentials())
}

// ListenAndServe This function blocks
func ListenAndServe(addr string) error {
	http.Handle("/", handlers.HTTPMethodOverrideHandler(r))
	return http.ListenAndServe(addr, nil)
}

// ListenAndServeCORS This function blocks
func ListenAndServeCORS(addr string) error {
	http.Handle("/", handlers.HTTPMethodOverrideHandler(r))
	return http.ListenAndServe(addr, handlers.CORS(corsOpts...)(r))
}
