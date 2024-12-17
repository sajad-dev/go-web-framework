package runserver

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	godotenv.Load(".env")

	file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(file)
}
