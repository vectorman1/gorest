package service

import (
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
		return common.PasswordTooShortError
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	str := string(password)
	user.Password = str
	user.Valid = true

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

	if user.Description == "" {
		user.Description = existingUser.Description
	}
	if user.AvatarUrl == "" {
		user.AvatarUrl = existingUser.AvatarUrl
	}
	if user.Role == "" {
		user.Role = existingUser.Role
	}
	if user.Username == "" {
		user.Username = existingUser.Username
	} else if user.Username != "" && existingUser.Username != user.Username {
		return common.InvalidModelError
	}
	if user.Password == "" {
		user.Password = existingUser.Password
	} else {
		fail := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
		if fail != nil {
			newPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
			user.Password = string(newPass)
		} else {
			user.Password = existingUser.Password
		}
	}

	return r.userRepository.Update(user)
}

func (r *UserService) DeleteByID(userID uint) (entity.User, error) {
	e, err := r.FindByID(userID)
	if err != nil {
		return entity.User{}, common.EntityNotFoundError
	}

	return r.userRepository.DeleteByID(e.ID)
}

func (r *UserService) Count() (int, error) {
	return r.userRepository.Count()
}
