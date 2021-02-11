package db

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorest/entity"
	"gorm.io/gorm"
)

type User interface {
	FindAll() (*[]entity.User, error)
	FindAllPagedAndSorted(
		pageNumber int,
		pageSize int,
		sortingAttribute string,
		ascending bool) (*[]entity.User, error)
	FindByID(id uint) (*entity.User, error)
	FindByUsername(username string) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	DeleteByID(userID uint) error
	Count() (int, error)
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) FindAll() (*[]entity.User, error) {
	var e []entity.User
	res := r.db.Find(&e).Error
	if res != nil {
		return nil, res
	}
	return &e, nil
}

func (r *UserRepository) FindAllPagedAndSorted(pageNumber int, pageSize int, sortingAttribute string, ascending bool) (*[]entity.User, error) {
	var e []entity.User
	var d string
	if ascending {
		d = "asc"
	} else {
		d = "desc"
	}
	order := fmt.Sprintf("%s %s", sortingAttribute, d)
	res := r.db.
		Order(order).
		Offset((pageNumber - 1) * pageSize).
		Limit(pageSize).
		Find(&e).
		Error
	if res != nil {
		return nil, res
	}
	return &e, nil
}

func (r *UserRepository) FindByID(id uint) (*entity.User, error) {
	var e entity.User
	res := r.db.First(&e, id)
	if res.Error != nil {
		return nil, res.Error
	}
	if &e != nil {
		return &e, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (r *UserRepository) FindByUsername(username string) (*entity.User, error) {
	var e entity.User
	res := r.db.First(&e, "username = ?", username).Error
	if res != nil {
		return nil, res
	}
	return &e, nil
}

func (r *UserRepository) Create(user *entity.User) error {
	if len(user.Password) <= 8 {
		return errors.New("password is too short")
	}
	newUser := entity.User{
		Username:    user.Username,
		Gender:      user.Gender,
		Role:        "user",
		AvatarUrl:   user.AvatarUrl,
		Description: user.Description,
		Valid:       true,
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	str := string(password)
	newUser.Password = str
	result := r.db.Create(&newUser)
	if result.Error != nil {
		return result.Error
	}
	*user = newUser
	return nil
}

func (r *UserRepository) Update(user *entity.User) error {
	var existingUser entity.User
	res := r.db.First(&existingUser, "id = ?", &user.ID)
	if res.Error != nil {
		return res.Error
	}

	if existingUser.Username != user.Username {
		return errors.New("can't change username")
	}

	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		newPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		existingUser.Password = string(newPass)
	}

	existingUser.Gender = user.Gender
	existingUser.Role = user.Role
	existingUser.AvatarUrl = user.AvatarUrl
	existingUser.Description = user.Description
	existingUser.Valid = user.Valid
	existingUser.Recipes = user.Recipes
	err = r.db.Save(existingUser).Error
	if err != nil {
		return err
	}

	user = &existingUser
	return nil
}

func (r *UserRepository) DeleteByID(userID uint) error {
	e, err := r.FindByID(userID)
	if err != nil {
		return err
	}
	e.Valid = false
	err = r.db.Delete(&e).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Count() (int, error) {
	e, err := r.FindAll()
	if err != nil {
		return 0, err
	}
	return len(*e), nil
}
