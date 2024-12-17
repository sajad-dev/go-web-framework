package migration

import "fmt"

func Char(name string, size int, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s CHAR(%d) %s DEFAULT %s %s %s", name, size, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func Varchar(name string, size int, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s VARCHAR(%d) %s DEFAULT %s %s %s", name, size, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func Binary(name string, size int, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s BINARY(%d) %s DEFAULT %s %s %s", name, size, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func Varbinary(name string, size int, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s VARBINARY(%d) %s DEFAULT %s %s %s", name, size, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func TinyBlob(name string, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s TINYBLOB %s DEFAULT %s %s %s", name, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func TinyText(name string, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s TINYTEXT %s DEFAULT %s %s %s", name, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func Text(name string, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s TEXT %s DEFAULT %s %s %s", name, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func Blob(name string, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s BLOB %s DEFAULT %s %s %s", name, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func MediumText(name string, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s MEDIUMTEXT %s DEFAULT %s %s %s", name, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func MediumBlob(name string, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s MEDIUMBLOB %s DEFAULT %s %s %s", name, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func LongText(name string, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s LONGTEXT %s DEFAULT %s %s %s", name, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func LongBlob(name string, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s LONGBLOB %s DEFAULT %s %s %s", name, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func Enum(name string, values []string, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s ENUM(%s) %s DEFAULT %s %s %s", name, joinValues(values), nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func Set(name string, values []string, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s SET(%s) %s DEFAULT %s %s %s", name, joinValues(values), nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

// Numeric types
func Bit(name string, size int, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s BIT(%d) %s DEFAULT %s %s %s", name, size, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func TinyInt(name string, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s TINYINT %s DEFAULT %s %s %s", name, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func Bool(name string, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s BOOL %s DEFAULT %s %s %s", name, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func Boolean(name string, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s BOOLEAN %s DEFAULT %s %s %s", name, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func SmallInt(name string, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s SMALLINT %s DEFAULT %s %s %s", name, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func MediumInt(name string, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s MEDIUMINT %s DEFAULT %s %s %s", name, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func Int(name string, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s INT %s DEFAULT %s %s %s", name, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func BigInt(name string, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s BIGINT %s DEFAULT %s %s %s", name, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func Float(name string, size int, d int, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s FLOAT(%d, %d) %s DEFAULT %s %s %s", name, size, d, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func Double(name string, size int, d int, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s DOUBLE(%d, %d) %s DEFAULT %s %s %s", name, size, d, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func Decimal(name string, size int, d int, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s DECIMAL(%d, %d) %s DEFAULT %s %s %s", name, size, d, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

// Date and Time types
func Date(name string, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s DATE %s DEFAULT %s %s %s", name, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func DateTime(name string, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s DATETIME %s DEFAULT %s %s %s", name, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func Time(name string, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s TIME %s DEFAULT %s %s %s", name, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func Year(name string, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s YEAR %s DEFAULT %s %s %s", name, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

func Timestamp(name string, nullable bool, defaultValue string, isUnique bool, isPrimaryKey bool) string {
	return fmt.Sprintf("%s TIMESTAMP %s DEFAULT %s %s %s", name, nullConstraint(nullable), defaultValueConstraint(defaultValue, nullable), uniqueConstraint(isUnique), primaryKeyConstraint(isPrimaryKey))
}

// Helper functions
func joinValues(values []string) string {
	result := ""
	for i, value := range values {
		result += fmt.Sprintf("'%s'", value)
		if i < len(values)-1 {
			result += ", "
		}
	}
	return result
}

func nullConstraint(nullable bool) string {
	if nullable {
		return "NULL"
	}
	return "NOT NULL"
}

func uniqueConstraint(isUnique bool) string {
	if isUnique {
		return "UNIQUE"
	}
	return ""
}

func primaryKeyConstraint(isPrimaryKey bool) string {
	if isPrimaryKey {
		return "PRIMARY KEY"
	}
	return ""
}

func defaultValueConstraint(value string, nullable bool) string {
	if value == "" {
		if nullable {
			return "NULL"
		} else {
			return "''"
		}
	}
	return fmt.Sprintf("'%s'", value)
}

func Forgin(name string,table string, nullable bool, tableForgin string, field_table string, cascade_ondelete string, cascade_onupdate string) string {

	return fmt.Sprintf(" %s INT %s , CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s(%s) ON DELETE %s ON UPDATE %s", name, nullConstraint(nullable), fmt.Sprintf("%s_%s_fg",table,name), name, tableForgin, field_table, cascade_ondelete, cascade_onupdate)
}
func IntPrimary(name string, increment bool) string {
	getincrement := ""
	if increment {
		getincrement = "AUTO_INCREMENT"
	} else {
		getincrement = ""

	}
	return fmt.Sprintf("%s INT NOT NULL %s PRIMARY KEY", name, getincrement)
}