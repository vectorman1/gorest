package service

import (
	"golang.org/x/crypto/bcrypt"
	"gorest/common"
	"gorest/db"
	"gorest/entity"
	"reflect"
)

type User interface {
	FindAll() ([]entity.User, error)
	FindAllPagedAndSorted(
		pageNumber int,
		pageSize int,
		sortingAttribute string,
		ascending bool) ([]entity.User, error)
	FindByID(id uint) (entity.User, error)
	FindByUsername(username string) (entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	DeleteByID(userID uint) error
	Count() (int, error)
	Login(username string, password string) (entity.User, error)
}

type UserService struct {
	userRepository db.UserRepository
}

func NewUserService(userRepo *db.UserRepository) *UserService {
	return &UserService{
		userRepository: *userRepo,
	}
}

func (r *UserService) FindAll() ([]entity.User, error) {
	return r.userRepository.FindAll()
}

func (r *UserService) FindAllPagedAndSorted(pageNumber int, pageSize int, sortingAttribute string, ascending bool) ([]entity.User, error) {
	return r.userRepository.FindAllPagedAndSorted(pageNumber, pageSize, sortingAttribute, ascending)
}

func (r *UserService) FindByID(id uint) (entity.User, error) {
	return r.userRepository.FindByID(id)
}

func (r *UserService) FindByUsername(username string) (entity.User, error) {
	return r.userRepository.FindByUsername(username)
}

func (r *UserService) Create(user *entity.User) error {
	if len(user.Password) <= 8 {
		return common.PasswordTooShortError
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	str := string(password)
	user.Password = str

	if user.AvatarUrl == "" {
		user.AvatarUrl = common.DEFAULT_AVATAR_URL
	}

	return r.userRepository.Create(user)
}

func (r *UserService) Update(user *entity.User) error {
	existingUser, err := r.FindByID(user.ID)
	if err != nil {
		return common.EntityNotFoundError
	}

	if existingUser.Username != user.Username {
		return common.InvalidModelError
	}

	fail := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if fail != nil {
		newPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		existingUser.Password = string(newPass)
	}

	if user.AvatarUrl == "" {
		user.AvatarUrl = common.DEFAULT_AVATAR_URL
	}

	if existingUser.Gender != user.Gender {
		existingUser.Gender = user.Gender
	}
	if existingUser.Role != user.Role {
		existingUser.Role = user.Role
	}
	if existingUser.AvatarUrl != user.AvatarUrl {
		existingUser.AvatarUrl = user.AvatarUrl
	}
	if existingUser.Description != user.Description {
		existingUser.Description = user.Description
	}
	if existingUser.Valid != user.Valid {
		existingUser.Valid = user.Valid
	}
	if !reflect.DeepEqual(existingUser.Recipes, user.Recipes) {
		existingUser.Recipes = user.Recipes
	}

	return r.userRepository.Update(&existingUser)
}

func (r *UserService) DeleteByID(userID uint) error {
	e, err := r.FindByID(userID)
	if err != nil {
		return common.EntityNotFoundError
	}

	return r.userRepository.DeleteByID(e.ID)
}

func (r *UserService) Count() (int, error) {
	return r.userRepository.Count()
}
