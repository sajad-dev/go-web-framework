package helpers

func IfThenElse(condition bool, a any, b any) any {
	if condition {
		return a
	}
	return b
}
