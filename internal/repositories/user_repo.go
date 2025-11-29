package storage

import (
	"errors"
	"taskBoard_API/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepository) GetByID(id uint) (models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	if err := r.db.Where("LOWER(email) = LOWER(?)", email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) EmailExists(email string) bool {
	var count int64
	r.db.Model(&models.User{}).Where("LOWER(email) = LOWER(?)", email).Count(&count)
	return count > 0
}

func (r *UserRepository) Create(user models.User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) UpdateName(id uint, updatedUser models.User) error {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return err
	}

	user.Name = updatedUser.Name
	return r.db.Save(&user).Error
}

func (r *UserRepository) DeleteByID(id uint) error {
	if result := r.db.Delete(&models.User{}, id); result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return errors.New("NOT FOUND")
	}
	return nil

}
