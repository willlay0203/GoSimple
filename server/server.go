package server

import (
	"net/http"
)

type Server struct {
	Mux  *http.ServeMux
	Port string
}

func Setup(port string) Server {
	m := Server{
		Mux:  http.NewServeMux(),
		Port: port,
	}

	return m
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
