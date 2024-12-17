package validation

var Message = map[string]string{
	"accepted":      "The field must be accepted",
	"active_url":    "The field must be a valid URL",
	"after_date":    "The field must be a date after the specified date",
	"alpha":         "The field must contain only alphabetic characters",
	"alpha_numeric": "The field must contain only alphanumeric characters",
	"required":      "The field is required",
	"email":         "The field must be a valid email address",
	"numeric":       "The field must be a number",
	"ip":            "The field must be a valid IP address",
	"boolean":       "The field must be true or false",
	"url":           "The field must be a valid URL",
	"min":           "The field must be at least %d characters",
	"max":           "The field must not exceed %d characters",
	"phone":         "The field must be a valid phone number",
	"confirmed":     "The field must be confirmed",
	"file_required": "The file is required",
	"mimes":         "The file must be of type: %s",
	"max_size":      "The file must be less than %d bytes",
	"extension":     "The file must have one of the following extensions: %s",
	"unique":        "The field %s exsis now",
}
