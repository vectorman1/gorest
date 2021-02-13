package models

import (
	"github.com/lib/pq"
)

type RecipesSearchRequest struct {
	Title    string
	Products pq.StringArray
	Tags     pq.StringArray
}
