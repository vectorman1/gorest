package db

import (
	"context"
	"github.com/lib/pq"
	"gorest/common"
	"gorest/entity"
	"gorm.io/gorm"
	"time"
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
	timeoutContext, c := context.WithTimeout(context.Background(), time.Second)
	defer c()

	err := r.db.WithContext(timeoutContext).Find(&e).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *RecipeRepository) FindAllPagedAndSorted(pageNumber int, pageSize int, sortingAttribute string, ascending bool) ([]entity.Recipe, error) {
	var e []entity.Recipe
	timeoutContext, c := context.WithTimeout(context.Background(), time.Second)
	defer c()

	order := common.FormatOrderQuery(sortingAttribute, ascending)
	err := r.db.
		WithContext(timeoutContext).
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
	timeoutContext, c := context.WithTimeout(context.Background(), time.Second)
	defer c()

	err := r.db.WithContext(timeoutContext).First(&e, id).Error
	if err != nil {
		return entity.Recipe{}, common.EntityNotFoundError
	}

	return e, nil
}

func (r *RecipeRepository) FindAllByTitle(title string) ([]entity.Recipe, error) {
	var e []entity.Recipe
	timeoutContext, c := context.WithTimeout(context.Background(), time.Second)
	defer c()

	err := r.db.
		WithContext(timeoutContext).
		Where("title LIKE ?", "%"+title+"%").
		Find(&e).
		Error
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *RecipeRepository) FindAllByProducts(products *pq.StringArray) ([]entity.Recipe, error) {
	timeoutContext, c := context.WithTimeout(context.Background(), time.Second)
	defer c()

	var res []entity.Recipe
	err := r.db.
		WithContext(timeoutContext).
		Where("products = ?", pq.Array(*products)).
		Find(&res).Error
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, common.EntityNotFoundError
	}
	return res, nil
}

func (r *RecipeRepository) FindAllByTags(tags *pq.StringArray) ([]entity.Recipe, error) {
	var res []entity.Recipe
	timeoutContext, c := context.WithTimeout(context.Background(), time.Second)
	defer c()

	err := r.db.
		WithContext(timeoutContext).
		Where("tags = ?", pq.Array(*tags)).
		Find(&res).Error
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, common.EntityNotFoundError
	}
	return res, nil
}

func (r *RecipeRepository) Create(recipe *entity.Recipe) error {
	timeoutContext, c := context.WithTimeout(context.Background(), time.Second)
	defer c()

	err := r.db.WithContext(timeoutContext).Create(&recipe).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *RecipeRepository) CreateBatch(recipes *[]entity.Recipe) error {
	timeoutContext, c := context.WithTimeout(context.Background(), time.Second)
	defer c()

	err := r.db.WithContext(timeoutContext).CreateInBatches(&recipes, 20).Error
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

	timeoutContext, c := context.WithTimeout(context.Background(), time.Second)
	defer c()

	err = r.db.WithContext(timeoutContext).Save(&recipe).Error
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

	timeoutContext, c := context.WithTimeout(context.Background(), time.Second)
	defer c()

	err = r.db.WithContext(timeoutContext).Delete(&e).Error
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
