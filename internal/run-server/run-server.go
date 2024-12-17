package runserver

import (
	"net/http"
	"os"

	"github.com/sajad-dev/go-web-framework/pkg/exception"
	"github.com/sajad-dev/go-web-framework/pkg/helpers"
)

func Run() {
	err := http.ListenAndServe(":8000", nil)
	exception.Log(err)
	if !helpers.IfThenElse(os.Getenv("DEBUG") == "true", true, false).(bool) {
		// defer log.Panicln("END PROGRAM")
	}
}
