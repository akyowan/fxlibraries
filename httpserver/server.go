package httpserver

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type HandleFunc func(*Request) *Response

type Router struct {
	*mux.Router
	corsOpts []handlers.CORSOption
}

type Route struct {
	*mux.Route
}

type OriginValidator func(string) bool

func NewRouter() *Router {
	r := &Router{
		mux.NewRouter(),
		make([]handlers.CORSOption, 0),
	}
	r.RouteAlive()
	return r
}

func (r *Router) RouteAlive() *Route {
	return &Route{r.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		response := NewResponse()
		response.Write(w)
	})}
}

func (r *Router) RouteHandleFunc(path string, f HandleFunc) *Route {
	return &Route{r.HandleFunc(path, HandlerWrapper(f))}
}

func (r *Router) AllowedMethods(methods []string) {
	r.corsOpts = append(r.corsOpts, handlers.AllowedMethods(methods))
}

func (r *Router) AllowedHeaders(headers []string) {
	r.corsOpts = append(r.corsOpts, handlers.AllowedHeaders(headers))
}

func (r *Router) AllowedOrigins(origins []string) {
	r.corsOpts = append(r.corsOpts, handlers.AllowedOrigins(origins))
}

func (r *Router) AllowedOriginValidator(fn OriginValidator) {
	r.corsOpts = append(r.corsOpts, handlers.AllowedOriginValidator(handlers.OriginValidator(fn)))
}

func (r *Router) ExposedHeaders(headers []string) {
	r.corsOpts = append(r.corsOpts, handlers.ExposedHeaders(headers))
}

func (r *Router) MaxAge(age int) {
	r.corsOpts = append(r.corsOpts, handlers.MaxAge(age))
}

func (r *Router) IgnoreOptions() {
	r.corsOpts = append(r.corsOpts, handlers.IgnoreOptions())
}

func (r *Router) AllowCredentials() {
	r.corsOpts = append(r.corsOpts, handlers.AllowCredentials())
}

// ListenAndServe This function blocks
func (r *Router) ListenAndServe(addr string) error {
	s := http.NewServeMux()
	s.Handle("/", handlers.HTTPMethodOverrideHandler(r))
	return http.ListenAndServe(addr, s)
}

// ListenAndServeCORS This function blocks
func (r *Router) ListenAndServeCORS(addr string) error {
	s := http.NewServeMux()
	s.Handle("/", handlers.HTTPMethodOverrideHandler(r))
	return http.ListenAndServe(addr, handlers.CORS(r.corsOpts...)(r))
}
