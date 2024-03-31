package repository

import (
	"gofinance/api/pkg/interfaces"
	"gofinance/api/pkg/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) Create(user *models.User) (*models.User, error) {
	result := ur.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (ur *userRepository) FindByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := ur.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) FindByGoogleID(googleID string) (*models.User, error) {
	var user models.User
	err := ur.db.First(&user, "google_id = ?", googleID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) FindAll() ([]*models.User, error) {
	var users []*models.User
	err := ur.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := ur.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) Update(id uuid.UUID, updates interface{}) error {
	result := ur.db.Where("id = ?", id).Updates(updates)
	return result.Error
}
