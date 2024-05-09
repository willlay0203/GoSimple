package server

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/willlay0203/gohttp/middleware"

	"github.com/rs/cors"
)

type Server struct {
	Mux         *http.ServeMux
	Port        string
	middlewares []middleware.Middleware
}

func Setup(port string) Server {
	m := Server{
		Mux:  http.NewServeMux(),
		Port: port,
	}

	return m
}

func (mux *Server) Start() {
	m := middleware.Adapt(mux.Mux, mux.middlewares...)

	// Wrap mux with CORs middleware
	handler := cors.Default().Handler(m)

	// Sets the default logging behaviour
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	log.Fatal(http.ListenAndServe(mux.Port, handler))
}

func (mux *Server) Enable(a middleware.Middleware) {
	mux.middlewares = append(mux.middlewares, a)
}

func (mux *Server) GET(p string, handler func(http.ResponseWriter, *http.Request)) {
	mux.Mux.HandleFunc("GET "+p, handler)
}

func (mux *Server) POST(p string, handler func(http.ResponseWriter, *http.Request)) {
	mux.Mux.HandleFunc("POST "+p, handler)
}

func (mux *Server) PUT(p string, handler func(http.ResponseWriter, *http.Request)) {
	mux.Mux.HandleFunc("PUT "+p, handler)
}

func (mux *Server) PATCH(p string, handler func(http.ResponseWriter, *http.Request)) {
	mux.Mux.HandleFunc("PATCH "+p, handler)
}

func (mux *Server) DELETE(p string, handler func(http.ResponseWriter, *http.Request)) {
	mux.Mux.HandleFunc("DELETE "+p, handler)
}
