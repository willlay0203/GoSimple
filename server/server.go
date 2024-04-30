package server

import (
	"gohttp/middleware"
	"log"
	"log/slog"
	"net/http"
	"os"
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

	// Sets the default logging behaviour
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	log.Fatal(http.ListenAndServe(mux.Port, m))
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
