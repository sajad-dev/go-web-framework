package middlewares

import (
	"net/http"
)

func (e MiddlewaresType) HandelMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		e(w, r, next)
	})
}

func Handler(middlewares []func(http.Handler) http.Handler, finaly http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		finaly = middlewares[i](finaly)
	}
	return finaly
}

func ConfigWriterAndReader(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		next.ServeHTTP(w, r)
	})
}

func DaynamicRoute(w http.ResponseWriter, r *http.Request, next http.Handler) {

}
