package common

import (
	"fmt"
)

func FormatOrderQuery(attr string, asc bool) string {
	var d string
	if asc {
		d = "asc"
	} else {
		d = "desc"
	}

	return fmt.Sprintf("%s %s", attr, d)
}

func Contains(src []string, str string) bool {
	for _, s := range src {
		if s == str {
			return true
		}
	}
	return false
}
