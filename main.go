package main

import (
	"fmt"
	"goroute/server"
	"log"
	"net/http"
)

// This file is only DEV testing, not actual test file
func main() {
	s := server.Setup(":8080") //gokit.Setup(":8080")

	fmt.Printf("Setup server at port %s", s.Port)

	s.GET("/test/{id}", testHandler)
	s.POST("/test", testHandler)
	s.PUT("/test", testHandler)
	s.PATCH("/test", testHandler)
	s.DELETE("/test", testHandler)

	log.Fatal(http.ListenAndServe(s.Port, s.Mux))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("hi")
	test := r.PathValue("id")
	fmt.Print(test)
}
