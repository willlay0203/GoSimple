package gokit

import (
	"net/http"
)

type Mux struct {
	mux  *http.ServeMux
	port string
}

func Setup(port string) Mux {
	m := Mux{
		mux:  http.NewServeMux(),
		port: port,
	}

	return m
}
