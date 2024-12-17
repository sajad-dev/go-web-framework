package validation

import (
	"encoding/json"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/sajad-dev/go-web-framework/pkg/controllers"
)

type TestFunc func() (string, string)

var ValidationFuncs = map[string]func(...interface{}) string{
	"required": func(args ...interface{}) string {
		if len(args) == 0 {
			return ""
		}
		return Required(args[0])
	},
	"accepted": func(args ...interface{}) string {
		if len(args) == 0 {
			return ""
		}
		return Accepted(args[0].(bool))
	},
	"active_url": func(args ...interface{}) string {
		if len(args) == 0 {
			return ""
		}
		return ActiveURL(args[0].(string))
	},
	"alpha": func(args ...interface{}) string {
		if len(args) == 0 {
			return ""
		}
		return Alpha(args[0].(string))
	},
	"alpha_numeric": func(args ...interface{}) string {
		if len(args) == 0 {
			return ""
		}
		return AlphaNumeric(args[0].(string))
	},
	"email": func(args ...interface{}) string {
		if len(args) == 0 {
			return ""
		}
		return Email(args[0].(string))
	},
	"numeric": func(args ...interface{}) string {
		if len(args) == 0 {
			return ""
		}
		return Numeric(args[0].(string))
	},
	"ip": func(args ...interface{}) string {
		if len(args) == 0 {
			return ""
		}
		return IP(args[0].(string))
	},
	"boolean": func(args ...interface{}) string {
		if len(args) == 0 {
			return ""
		}
		return Boolean(args[0])
	},
	"url": func(args ...interface{}) string {
		if len(args) == 0 {
			return ""
		}
		return URL(args[0].(string))
	},
	"min": func(args ...interface{}) string {
		if len(args) < 2 {
			return ""
		}
		return Min(args[0].(string), args[1].(int))
	},
	"max": func(args ...interface{}) string {
		if len(args) < 2 {
			return ""
		}
		return Max(args[0].(string), args[1].(int))
	},
	"phone": func(args ...interface{}) string {
		if len(args) == 0 {
			return ""
		}
		return Phone(args[0].(string))
	},
	"confirmed": func(args ...interface{}) string {
		if len(args) < 2 {
			return ""
		}
		return Confirmed(args[0].(string), args[1].(string))
	},
	"file": func(args ...interface{}) string {
		if len(args) == 0 {
			return ""
		}
		return File(args[0].(*Blob))
	},
	"mimes": func(args ...interface{}) string {
		if len(args) < 2 {
			return ""
		}
		return MIMEType(args[0].(*Blob), args[1].([]string))
	},
	"max_size": func(args ...interface{}) string {
		if len(args) < 2 {
			return ""
		}
		return MaxSize(args[0].(*Blob), args[1].(int))
	},
	"extension": func(args ...interface{}) string {
		if len(args) < 2 {
			return ""
		}
		return Extension(args[0].(*Blob), args[1].([]string))
	},
}

func Handel(form map[string]string,w http.ResponseWriter) bool {
	pc, _, _, _ := runtime.Caller(1)
	name := strings.Split(runtime.FuncForPC(pc).Name(), ".")[2] + "Struct"
	structGet := controllers.StructRegistry[name]
	v := reflect.ValueOf(structGet)
	arr := []map[string]string{}

	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i)
			name := strings.ToLower(field.Name)
			valid := strings.Split(field.Tag.Get("validation"), "|")
			val := map[string]string{}
			for _, v := range valid {
				if strings.Contains(v, ":") {
					spl := strings.Split(v, ":")
					index, _ := strconv.Atoi(spl[1])

					val[v] = ValidationFuncs[spl[0]](form[name], index)
				} else {
					val[v] = ValidationFuncs[v](form[name])
				}
			}
			arr = append(arr, val)
		}
	}
	output := []string{}
	for _, v := range arr {
		for _, v := range v {
			if v != "" {
				output = append(output, v)
			}
		}
	}
	if len(output) > 0  {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(output)
		return true
	}
	return false
}
