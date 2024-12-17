package api

import (
	"context"
	"net/http"
	"regexp"
	"strings"

	"github.com/sajad-dev/go-web-framework/pkg/exception"
	"github.com/sajad-dev/go-web-framework/pkg/middlewares"
)

type methodType string

const (
	POST   methodType = "GET"
	GET    methodType = "GET"
	PUT    methodType = "PUT"
	PATCH  methodType = "PATCH"
	DELETE methodType = "DELETE"
)

func CheckMethod(next http.Handler, method methodType) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != string(method) {
			exception.Response405(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func DaynamicRoute(next http.Handler, pattern []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := strings.Split(r.URL.Path, "/")
		error404 := false
		var parameters = map[string]string{}
		for i, pat := range pattern {
			if pat != url[i] && !strings.Contains(pat, "{") {
				error404 = true
				break
			}
			if strings.Contains(pat, "{") {
				re, _ := regexp.Compile(`\{([^}]*)\}`)

				matches := re.FindStringSubmatch(pat)
				if len(matches) > 1 {
					parameters[matches[1]] = url[i]
				}
			}
		}

		if error404 {
			exception.Response404(w)
			return
		}
		ctx := context.WithValue(r.Context(), "parameters", parameters)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetDYRoute(pattern string) ([]string, string) {
	slice_pat := strings.Split(pattern, "/")
	rou := ""
	dy := false
	for _, pat := range slice_pat {
		if strings.Contains(pat, "{") {
			dy = true
			break
		}
		if pat != "" {
			rou += "/" + pat
		}
	}
	if dy {
		return slice_pat, rou +"/"

	}
	return slice_pat, rou
}
func Route(pattern string, method methodType, controller http.Handler, middlewaresList []func(http.Handler) http.Handler) {
	sli, route := GetDYRoute(pattern)
	http.Handle(route,  middlewares.ConfigWriterAndReader(CheckMethod(DaynamicRoute(middlewares.Handler(middlewaresList,controller), sli), method)))
}

func RouteRun() {
	for _, route := range RouteList {
		Route(route.Pattern, route.Method, route.Controller, route.Middlewares)
	}
}

func RouteAddFunc(Controller func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(Controller)
}
