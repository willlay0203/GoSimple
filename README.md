# GoHttp

Example usage

```
func main() {
	s := server.Setup(":8080") //gokit.Setup(":8080")

	fmt.Printf("Setup server at port %s\n", s.Port)

	// Assigning routes
	s.GET("/test/{id}", testHandler)
	s.POST("/test", testHandler)
	s.PUT("/test", testHandler)
	s.PATCH("/test", testHandler)
	s.DELETE("/test", testHandler)

	// Enabling middleware
	s.Enable(middleware.RequestLogger())
	s.Start()
}

func testHandler(w http.ResponseWriter, r *http.Request) error {
	fmt.Print("hi")
	test := r.PathValue("id")
	fmt.Print(test)
	// APIError{Status: 500, Msg: "something"}
	return nil
}
```
