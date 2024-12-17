package middlewares

import "net/http"


type MiddlewaresType func(w http.ResponseWriter, r *http.Request, next http.Handler)
