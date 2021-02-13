package common

import (
	"fmt"
	"gorest/entity"
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

func Unique(arr *[]entity.Recipe) []entity.Recipe {
	m := make(map[uint]bool)
	var result []entity.Recipe
	for _, v := range *arr {
		if !m[v.ID] {
			m[v.ID] = true
			result = append(result, v)
		}
	}
	return result
}
