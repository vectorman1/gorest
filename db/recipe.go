package db

import (
	"gorest/entity"
	"gorm.io/gorm"
)

type Recipe interface {
	FindAll() ([]entity.Recipe, error)
	FindAllPagedAndSorted(pageNumber int, pageSize int, sortingAttribute string, ascending bool) ([]entity.Recipe, error)
	FindByID(id uint) (entity.Recipe, error)
	FindAllByTitle(str string) ([]entity.Recipe, error)
	FindAllByProducts(products *entity.Products) ([]entity.Recipe, error)
	FindAllByTags(tags *entity.Tags) ([]entity.Recipe, error)
	Create(recipe *entity.Recipe) error
	CreateBatch(recipes *[]entity.Recipe) error
	Update(user *entity.Recipe) error
	DeleteByID(recipeID uint) error
	Count() int
}

func NewRecipeRepository(db *gorm.DB) *RecipeRepository {
	return &RecipeRepository{
		db: db,
	}
}

type RecipeRepository struct {
	db *gorm.DB
}

func (r *RecipeRepository) FindAll() ([]entity.Recipe, error) {
	return nil, nil
}

func (r *RecipeRepository) FindAllPagedAndSorted(
	pageNumber int,
	pageSize int,
	sortingAttribute string,
	ascending bool) ([]entity.Recipe, error) {
	return nil, nil
}

func (r *RecipeRepository) FindByID(id uint) (entity.Recipe, error) {
	return entity.Recipe{}, nil
}

func (r *RecipeRepository) FindAllByTitle(str string) ([]entity.Recipe, error) {
	return nil, nil
}

func (r *RecipeRepository) FindAllByProducts(products *entity.Products) ([]entity.Recipe, error) {
	return nil, nil
}

func (r *RecipeRepository) FindAllByTags(tags *entity.Tags) ([]entity.Recipe, error) {
	return nil, nil
}

func (r *RecipeRepository) Create(recipe *entity.Recipe) error {
	return nil
}

func (r *RecipeRepository) CreateBatch(recipes *[]entity.Recipe) error {
	return nil
}

func (r *RecipeRepository) Update(user *entity.Recipe) error {
	return nil
}

func (r *RecipeRepository) DeleteByID(recipeID uint) error {
	return nil
}

func (r *RecipeRepository) Count() int {
	return 0
}
