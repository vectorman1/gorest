package service

import (
	"gorest/common"
	"gorest/db"
	"gorest/delivery/models"
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
	Search(req models.RecipesSearchRequest) ([]entity.Recipe, error)
}

type RecipeService struct {
	recipeRepository db.RecipeRepository
	userRepository   db.UserRepository
}

func NewRecipeService(recipeRepo *db.RecipeRepository, userRepo *db.UserRepository) *RecipeService {
	return &RecipeService{
		recipeRepository: *recipeRepo,
		userRepository:   *userRepo,
	}
}

type RecipeRepository struct {
	db *gorm.DB
}

// External r.recipeRepository methods
func (r *RecipeService) FindAll() ([]entity.Recipe, error) {
	return r.recipeRepository.FindAll()
}

func (r *RecipeService) FindAllPagedAndSorted(pageNumber int, pageSize int, sortingAttribute string, ascending bool) ([]entity.Recipe, error) {
	return r.recipeRepository.FindAllPagedAndSorted(pageSize, pageSize, sortingAttribute, ascending)
}

func (r *RecipeService) FindByID(id uint) (entity.Recipe, error) {
	return r.recipeRepository.FindByID(id)
}

func (r *RecipeService) FindAllByTitle(title string) ([]entity.Recipe, error) {
	return r.recipeRepository.FindAllByTitle(title)
}

func (r *RecipeService) FindAllByProducts(products *entity.Products) ([]entity.Recipe, error) {
	return r.recipeRepository.FindAllByProducts(products)
}

func (r *RecipeService) FindAllByTags(tags *entity.Tags) ([]entity.Recipe, error) {
	return r.recipeRepository.FindAllByTags(tags)
}

func (r *RecipeService) Create(recipe *entity.Recipe) error {
	if recipe.ImageUrl == "" {
		recipe.ImageUrl = common.DEFAULT_AVATAR_URL
	}
	return r.recipeRepository.Create(recipe)
}

func (r *RecipeService) CreateBatch(recipes *[]entity.Recipe) error {
	for _, v := range *recipes {
		if v.ImageUrl == "" {
			v.ImageUrl = common.DEFAULT_AVATAR_URL
		}
	}

	return r.recipeRepository.CreateBatch(recipes)
}

func (r *RecipeService) Update(recipe *entity.Recipe) error {
	return r.recipeRepository.Update(recipe)
}

func (r *RecipeService) DeleteByID(recipeID uint) (entity.Recipe, error) {
	return r.recipeRepository.DeleteByID(recipeID)
}

func (r *RecipeService) Count() (int, error) {
	return r.recipeRepository.Count()
}

// Service methods
func (r *RecipeService) Search(req models.RecipesSearchRequest) ([]entity.Recipe, error) {
	var result []entity.Recipe
	if req.Products != nil {
		recipesByProducts, err := r.recipeRepository.FindAllByProducts(&req.Products)
		if err == nil {
			result = append(result, recipesByProducts...)
		}
	}
	if req.Tags == nil {
		recipesByTags, err := r.recipeRepository.FindAllByTags(&req.Tags)
		if err == nil {
			result = append(result, recipesByTags...)
		}
	}
	if req.Title != "" {
		recipesByTitle, err := r.recipeRepository.FindAllByTitle(req.Title)
		if err == nil {
			result = append(result, recipesByTitle...)
		}
	}

	uniqueRes := common.Unique(&result)
	return uniqueRes, nil
}
