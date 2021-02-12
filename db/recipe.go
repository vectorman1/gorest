package db

import (
	"errors"
	"fmt"
	"gorest/common"
	"gorest/entity"
	"gorm.io/gorm"
)

type Recipe interface {
	FindAll() ([]entity.Recipe, error)
	FindAllPagedAndSorted(
		pageNumber int,
		pageSize int,
		sortingAttribute string,
		ascending bool) ([]entity.Recipe, error)
	FindByID(id uint) (entity.Recipe, error)
	FindAllByTitle(str string) ([]entity.Recipe, error)
	FindAllByProducts(products *entity.Products) ([]entity.Recipe, error)
	FindAllByTags(tags *entity.Tags) ([]entity.Recipe, error)
	Create(recipe *entity.Recipe) error
	CreateBatch(recipes *[]entity.Recipe) error
	Update(user *entity.Recipe) error
	DeleteByID(recipeID uint) (entity.Recipe, error)
	Count() (int, error)
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
	var e []entity.Recipe
	err := r.db.Find(&e).Error

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *RecipeRepository) FindAllPagedAndSorted(pageNumber int, pageSize int, sortingAttribute string, ascending bool) ([]entity.Recipe, error) {
	var e []entity.Recipe
	order := common.FormatOrderQuery(sortingAttribute, ascending)
	err := r.db.
		Order(order).
		Offset((pageNumber - 1) * pageSize).
		Limit(pageSize).
		Find(&e).
		Error

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *RecipeRepository) FindByID(id uint) (entity.Recipe, error) {
	var e entity.Recipe
	err := r.db.First(&e, id).Error
	if err != nil {
		return entity.Recipe{}, err
	}

	return e, nil
}

func (r *RecipeRepository) FindAllByTitle(title string) ([]entity.Recipe, error) {
	var e []entity.Recipe
	err := r.db.
		Where("title LIKE ?", fmt.Sprintf("%%%s%%", title)).
		Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *RecipeRepository) FindAllByProducts(products *entity.Products) ([]entity.Recipe, error) {
	var e []entity.Recipe
	err := r.db.
		Where("products <> ?", nil).Find(&e).Error
	if err != nil {
		return nil, err
	}

	var res []entity.Recipe
	for _, r := range e {
		for _, p := range *products {
			if common.Contains(r.Products, p) {
				res = append(res, r)
			}
		}
	}
	if len(res) == 0 {
		return nil, errors.New("no recipes found")
	}

	return res, nil
}

func (r *RecipeRepository) FindAllByTags(tags *entity.Tags) ([]entity.Recipe, error) {
	var e []entity.Recipe
	err := r.db.
		Where("tags <> ?", nil).Find(&e).Error
	if err != nil {
		return nil, err
	}

	var res []entity.Recipe
	for _, r := range e {
		for _, t := range *tags {
			if common.Contains(r.Tags, t) {
				res = append(res, r)
			}
		}
	}
	if len(res) == 0 {
		return nil, errors.New("no recipes found")
	}

	return res, nil
}

func (r *RecipeRepository) Create(recipe *entity.Recipe) error {
	err := r.db.Create(&recipe).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *RecipeRepository) CreateBatch(recipes *[]entity.Recipe) error {
	err := r.db.CreateInBatches(&recipes, 20).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *RecipeRepository) Update(recipe *entity.Recipe) error {
	_, err := r.FindByID(recipe.ID)
	if err != nil {
		return err
	}

	err = r.db.Save(&recipe).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *RecipeRepository) DeleteByID(recipeID uint) (entity.Recipe, error) {
	e, err := r.FindByID(recipeID)
	if err != nil {
		return entity.Recipe{}, err
	}

	err = r.db.Delete(&e).Error
	if err != nil {
		return entity.Recipe{}, err
	}

	return e, nil
}

func (r *RecipeRepository) Count() (int, error) {
	e, err := r.FindAll()
	if err != nil {
		return 0, err
	}

	return len(e), nil
}
