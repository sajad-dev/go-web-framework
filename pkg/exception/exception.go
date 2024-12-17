package exception

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/fatih/color"
	"github.com/sajad-dev/go-web-framework/pkg/helpers"
)

func Response500(w http.ResponseWriter, exception string) {
	_, file, line, ok := runtime.Caller(1)
	if helpers.IfThenElse(os.Getenv("DEBUG") == "true", true, false).(bool) {
		res := fmt.Sprintf("Error occurred in %s:%d - %s", file, line, exception)
		json.NewEncoder(w).Encode(CustomError{Message: res, Code: 500, Status: false})

		return
	}

	if ok {
		log.Printf("Error occurred in %s:%d - %s", file, line, exception)
	} else {
		log.Println("Error occurred:", exception)
	}

	json.NewEncoder(w).Encode(CustomError{Message: "Internal Server Error", Code: 500, Status: false})
}

func Response405(w http.ResponseWriter) {
	_, file, line, ok := runtime.Caller(1)
	if helpers.IfThenElse(os.Getenv("DEBUG") == "true", true, false).(bool) {
		res := fmt.Sprintf("Error occurred in %s:%d - %s", file, line, "Method Not Allowed")
		json.NewEncoder(w).Encode(CustomError{Message: res, Code: 500, Status: false})

		return
	}

	if ok {
		log.Printf("Error occurred in %s:%d - %s", file, line, "Method Not Allowed")
	} else {
		log.Println("Error occurred:", "Method Not Allowed")
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	
	json.NewEncoder(w).Encode(CustomError{Message: "Method Not Allowed", Code: 405, Status: false})
}

func Response404(w http.ResponseWriter) {

	json.NewEncoder(w).Encode(CustomError{Message: "Not Found", Code: 404, Status: false})
}

func Log(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		if !helpers.IfThenElse(os.Getenv("DEBUG") == "true", true, false).(bool) {
			log.Panicln(err)
		} else {
			erro := fmt.Sprintf("%s - line %d - file %s",err.Error(),line,file)
			color.Red(erro)
		}
	}
}
