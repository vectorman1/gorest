package common

import (
	"fmt"
	"gorest/entity"
	"sort"
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

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Strings(a)
	sort.Strings(b)
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Unique(arr *[]entity.Recipe) []entity.Recipe {
	m := make(map[uint]bool)
	var result []entity.Recipe
	for _, v := range *arr {
		if !m[v.ID] {
			result = append(result, v)
		}
	}
	return result
}
