package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorest/common"
	"gorest/db"
	"gorest/entity"
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
		return errors.New("password is too short")
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
		return err
	}

	if existingUser.Username != user.Username {
		return errors.New("can't change username")
	}

	fail := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if fail != nil {
		newPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		existingUser.Password = string(newPass)
	}

	if user.AvatarUrl == "" {
		user.AvatarUrl = common.DEFAULT_AVATAR_URL
	}

	existingUser.Gender = user.Gender
	existingUser.Role = user.Role
	existingUser.AvatarUrl = user.AvatarUrl
	existingUser.Description = user.Description
	existingUser.Valid = user.Valid
	existingUser.Recipes = user.Recipes

	return r.userRepository.Update(&existingUser)
}

func (r *UserService) DeleteByID(userID uint) error {
	// TODO rework
	return r.userRepository.DeleteByID(userID)
}

func (r *UserService) Count() (int, error) {
	return r.userRepository.Count()
}
