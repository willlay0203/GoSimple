package middleware

type Middleware interface {
	ServeHTTP()
}
