package connection

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sajad-dev/go-web-framework/pkg/exception"
	// "github.com/sajad-dev/go-web-framework/pkg/middlewares"
)

var Database *sql.DB

func Connection() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/", os.Getenv("USERNAME_DB"), os.Getenv("PASSWORD_DB"), os.Getenv("IP_DB"), os.Getenv("PORT_DB")))
	exception.Log(err)

	databasename := os.Getenv("DATABASE_DB")
	_, _ = db.Exec("CREATE DATABASE " + databasename)

	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("USERNAME_DB"), os.Getenv("PASSWORD_DB"), os.Getenv("IP_DB"), os.Getenv("PORT_DB"), databasename))
	exception.Log(err)

	if err := db.Ping(); err != nil {
		exception.Log(err)
	}

	Database = db
}
