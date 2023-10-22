package repository

import (
	"github.com/junshintakeda/internet-forum/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *models.User, email string) error
	GetUserByUserName(users *[]models.User, userName string) error
	CreateUser(user *models.User) error
	DeleteUser(userId uint) error
}

type userrepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userrepository{db}
}

func (ur *userrepository) GetUserByEmail(user *models.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userrepository) GetUserByUserName(users *[]models.User, userName string) error {
	if err := ur.db.Where("user_name=?", userName).Find(users).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userrepository) CreateUser(user *models.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userrepository) DeleteUser(userId uint) error {
	result := ur.db.Delete(&models.User{}, userId)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
