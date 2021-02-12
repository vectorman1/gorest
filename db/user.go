package db

import (
	"context"
	"gorest/common"
	"gorest/entity"
	"gorm.io/gorm"
	"time"
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

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) FindAll() ([]entity.User, error) {
	var e []entity.User
	timeoutContext, c := context.WithTimeout(context.Background(), time.Second)
	defer c()

	err := r.db.WithContext(timeoutContext).Find(&e).Error
	if err != nil {
		return nil, common.EntityNotFoundError
	}

	return e, nil
}

func (r *UserRepository) FindAllPagedAndSorted(pageNumber int, pageSize int, sortingAttribute string, ascending bool) ([]entity.User, error) {
	var e []entity.User
	timeoutContext, c := context.WithTimeout(context.Background(), time.Second)
	defer c()

	order := common.FormatOrderQuery(sortingAttribute, ascending)
	err := r.db.WithContext(timeoutContext).
		Order(order).
		Offset((pageNumber - 1) * pageSize).
		Limit(pageSize).
		Find(&e).
		Error
	if err != nil {
		return nil, common.EntityNotFoundError
	}

	return e, nil
}

func (r *UserRepository) FindByID(id uint) (entity.User, error) {
	var e entity.User
	timeoutContext, c := context.WithTimeout(context.Background(), time.Second)
	defer c()

	err := r.db.WithContext(timeoutContext).First(&e, id).Error
	if err != nil {
		return entity.User{}, common.EntityNotFoundError
	}

	return e, nil
}

func (r *UserRepository) FindByUsername(username string) (entity.User, error) {
	var e entity.User
	timeoutContext, c := context.WithTimeout(context.Background(), time.Second)
	defer c()

	err := r.db.WithContext(timeoutContext).First(&e, "username = ?", username).Error
	if err != nil {
		return entity.User{}, common.EntityNotFoundError
	}

	return e, nil
}

func (r *UserRepository) Create(user *entity.User) error {
	timeoutContext, c := context.WithTimeout(context.Background(), time.Second)
	defer c()

	err := r.db.WithContext(timeoutContext).Create(&user).Error
	if err != nil {
		return common.InvalidModelError
	}

	return nil
}

func (r *UserRepository) Update(user *entity.User) error {
	timeoutContext, c := context.WithTimeout(context.Background(), time.Second)
	defer c()

	err := r.db.WithContext(timeoutContext).Save(&user).Error
	if err != nil {
		return common.InvalidModelError
	}

	return nil
}

func (r *UserRepository) DeleteByID(userID uint) error {
	e, err := r.FindByID(userID)
	if err != nil {
		return err
	}

	timeoutContext, c := context.WithTimeout(context.Background(), time.Second)
	defer c()

	e.Valid = false
	err = r.db.WithContext(timeoutContext).Delete(&e).Error
	if err != nil {
		return common.EntityNotFoundError
	}

	return nil
}

func (r *UserRepository) Count() (int, error) {
	e, err := r.FindAll()
	if err != nil {
		return 0, common.EntityNotFoundError
	}

	return len(e), nil
}
