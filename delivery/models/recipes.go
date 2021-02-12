package models

import "gorest/entity"

type RecipesSearchRequest struct {
	Title    string
	Products entity.Products
	Tags     entity.Tags
}
