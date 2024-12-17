package api

import (

	"net/http"


)

type MiddlewaresListType []func(http.Handler) http.Handler
type ApiType struct {
	Pattern     string
	Method      methodType
	Controller  http.Handler
	Middlewares MiddlewaresListType
}



var RouteList = []ApiType{
}
