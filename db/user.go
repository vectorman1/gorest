package db

import (
	"gorest/common"
	"gorest/entity"
	"gorm.io/gorm"
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
	err := r.db.Find(&e).Error

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *UserRepository) FindAllPagedAndSorted(pageNumber int, pageSize int, sortingAttribute string, ascending bool) ([]entity.User, error) {
	var e []entity.User
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

func (r *UserRepository) FindByID(id uint) (entity.User, error) {
	var e entity.User
	err := r.db.First(&e, id).Error
	if err != nil {
		return entity.User{}, err
	}

	return e, nil
}

func (r *UserRepository) FindByUsername(username string) (entity.User, error) {
	var e entity.User
	err := r.db.First(&e, "username = ?", username).Error
	if err != nil {
		return entity.User{}, err
	}

	return e, nil
}

func (r *UserRepository) Create(user *entity.User) error {
	err := r.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Update(user *entity.User) error {
	err := r.db.Save(user).Error
	if err != nil {
		return err
	}

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

	return len(e), nil
}
