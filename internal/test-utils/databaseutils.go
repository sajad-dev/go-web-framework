package testutils

import (
	"bytes"
	"os/exec"
	"strings"
)

type Mig struct{}

func MiggarionListAppend() ([]string, []string) {
	ou := exec.Command("go", "doc", "migration", "MigrationInterface")

	output, _ := ou.Output()
	res_table := []string{}
	res_func := []string{}
	slice := bytes.Split(output, []byte("\n"))
	for _, s := range slice {
		text := string(strings.ReplaceAll(string(s), "func", ""))
		index := strings.Index(text, "(")
		migration_index := strings.Index(text, "Migration")
		if index > 0 && migration_index > 0 {
			table_name := strings.ToLower(text[:index][:migration_index])
			res_table = append(res_table, strings.TrimSpace(table_name))
			res_func = append(res_func, strings.TrimSpace(string(text[:index])))
		}
	}
	return res_table, res_func
}

