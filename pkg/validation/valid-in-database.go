package validation

import (
	"fmt"

	"github.com/sajad-dev/go-web-framework/pkg/model"
)


func Unique(str string, field string, table string) string {
	output := model.Get([]string{field}, table, []model.Where_st{
		model.Where_st{Key: field, Value: str, After: "", Operator: "="},
	}, "", false)

	if len(output) > 0 {
		return fmt.Sprintf(Message["unique"], str)
	}
	return ""
}
