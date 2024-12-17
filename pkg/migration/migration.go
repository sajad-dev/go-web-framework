package migration

import (
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/sajad-dev/go-web-framework/pkg/connection"
	"github.com/sajad-dev/go-web-framework/pkg/exception"
	"golang.org/x/exp/slices"
)

type MigrateFunc func() []string

type Migrate struct {
	Table     string
	Funcation MigrateFunc
}

type Migration struct {
	MigrateList []*Migrate
}

type MigrationInterface interface {
	AppendToMigrationList()
	CheckTable()
	AddTable(strSlice string, table string)
	checkDeletedTable()
	UpdateTable(table string, column_name_old string, parametr string)
	HandelUpdate(field string, fieldtype string, collection sql.NullString,
		null string, defult sql.NullString, key string, extera string, privileges string, comment string, function MigrateFunc, table string, x int)
}

func (migrate *Migration) AppendToMigrationList() {
	migrate.MigrateList = ArrMigrations
}

func (migrate *Migration) CheckTable() {
	database := os.Getenv("DATABASE_DB")
	migrate.checkDeletedTable()

	sqlqu := fmt.Sprintf(` 
SELECT table_name
FROM information_schema.tables 
WHERE table_schema = '%s';
 `, database)

	row, err := connection.Database.Query(sqlqu)
	exception.Log(err)

	for row.Next() {
		var tb string

		err := row.Scan(&tb)
		exception.Log(err)

	}
	for _, m := range migrate.MigrateList {

		sql_qu := fmt.Sprintf("SHOW FULL COLUMNS FROM %s", m.Table)
		qu, err := connection.Database.Query(sql_qu)
		exception.Log(err)

		x := 0

		for qu.Next() {

			var (
				field, fieldtype, null, key, extra, privileges, comment string
				collection, defult                                      sql.NullString
			)
			err = qu.Scan(&field, &fieldtype, &collection, &null, &key, &defult, &extra, &privileges, &comment)
			exception.Log(err)

			migrate.HandelUpdate(field, fieldtype, collection, null, defult, key, extra, privileges, comment, m.Funcation, m.Table, x)
			x++

		}
		strSlice := m.Funcation()
		if len(strSlice) > x {
			for i := x; i < len(strSlice); i++ {
				migrate.AddTable(strSlice[i], m.Table)
			}
		}

	}
}

func (migrate *Migration) HandelUpdate(field string, fieldtype string, collection sql.NullString,
	null string, defult sql.NullString, key string, extera string, privileges string, comment string, function MigrateFunc, table string, x int) {
	strSlice := function()

	if x >= len(strSlice) {
		sql_qu := fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s", table, field)
		_, err := connection.Database.Query(sql_qu)
		exception.Log(err)
		return
	}

	if key == "UNI" && strings.Contains(strSlice[x], "UNIQUE") {

		sql_qu := fmt.Sprintf("SELECT DISTINCT COLUMN_NAME, INDEX_NAME   FROM information_schema.statistics  WHERE table_name = '%s'", table)
		fmt.Println(table)
		qu, err := connection.Database.Query(sql_qu)
		exception.Log(err)

		unilist := []string{}
		for qu.Next() {
			var column, key string
			qu.Scan(&column, &key)
			if column == field {
				unilist = append(unilist, key)
			}
		}
		for _, v := range unilist {
			sql_qu := fmt.Sprintf("ALTER TABLE %s DROP INDEX %s", table, v)
			connection.Database.Query(sql_qu)
			exception.Log(err)
		}

	}

	if key == "PRI" {
		strSlice[x] = strings.ReplaceAll(strSlice[x], "PRIMARY KEY", "")
	}
	migrate.UpdateTable(table, field, strSlice[x])
}

func (migrate *Migration) checkDeletedTable() {
	row, err := connection.Database.Query("SHOW TABLES")
	exception.Log(err)

	tableArr := []string{}
	for _, v := range migrate.MigrateList {
		tableArr = append(tableArr, v.Table)
	}
	for row.Next() {
		var table string
		row.Scan(&table)

		if !slices.Contains(tableArr, table) {
			connection.Database.Query(fmt.Sprintf("DROP TABLE %s", table))
		}

	}

}
func Handel() {
	var migra MigrationInterface
	migra = &Migration{}
	migra.AppendToMigrationList()
	migra.CheckTable()
}

func (migrate *Migration) AddTable(strSlice string, table string) {
	if strings.Contains(strSlice, "FOREIGN") {
		re := regexp.MustCompile(`\((.*?)\)`)
		matches := re.FindStringSubmatch(strSlice)
		sql_qu := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s INT", table, matches[1])
		_, err := connection.Database.Query(sql_qu)
		exception.Log(err)
		sql_qu = fmt.Sprintf("ALTER TABLE %s ADD CONSTRAINT %s %s", table, matches[1], strSlice)
		_, err = connection.Database.Query(sql_qu)
		exception.Log(err)
		return
	}
	sql_qu := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s", table, strSlice)
	_, err := connection.Database.Query(sql_qu)
	exception.Log(err)
}

func (migrate *Migration) UpdateTable(table string, column_name_old string, parametr string) {
	sql_qu := fmt.Sprintf("ALTER TABLE %s CHANGE %s %s ", table, column_name_old, parametr)
	if strings.Contains(sql_qu, "FOREIGN") {
		sql_du2 := fmt.Sprintf("ALTER TABLE %s DROP FOREIGN KEY %s", table, table+"_"+column_name_old+"_fg")
		connection.Database.Query(sql_du2)
		sql_qu = fmt.Sprintf("ALTER TABLE %s CHANGE COLUMN %s %s ", table, column_name_old, parametr)
		sql_qu = strings.Replace(sql_qu, "CONSTRAINT", "ADD CONSTRAINT ", 1)
	}

	_, err := connection.Database.Query(sql_qu)
	exception.Log(err)
}
