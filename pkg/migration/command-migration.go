package migration

import (
	"fmt"
	"os"
	"strings"

	"github.com/sajad-dev/go-web-framework/pkg/connection"
	"github.com/sajad-dev/go-web-framework/pkg/exception"
)



func CreateAll() {

	for _, v := range ArrMigrations {
		CreateTable(v.Funcation(), v.Table)
	}
}

func CreateTable(strSlice []string, table string) {
	co := strings.Join(strSlice, " , ")

	str_sql := fmt.Sprintf("CREATE TABLE  IF NOT EXISTS %s (%s)", table, co)
	_, err := connection.Database.Query(str_sql)
	exception.Log(err)

}

func DropTable() {
	database := os.Getenv("DATABASE_DB")
	sql := fmt.Sprintf(` 
SELECT table_name
FROM information_schema.tables 
WHERE table_schema = '%s';
 `, database)
	row, err := connection.Database.Query(sql)
	exception.Log(err)

	for row.Next() {
		var tb string

		err := row.Scan(&tb)
		exception.Log(err)

		_, err = connection.Database.Query(fmt.Sprintf("DROP TABLE IF EXISTS %s", tb))
		exception.Log(err)

	}
}
