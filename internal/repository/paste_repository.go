package repository

import (
	"codeDrop/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PasteRepository struct {
	DB *gorm.DB
}

func NewPasteRepository(db *gorm.DB) *PasteRepository {
	return &PasteRepository{DB: db}
}

func (r *PasteRepository) Create(paste *models.Paste) error {
	return r.DB.Create(paste).Error
}

func (r *PasteRepository) FindById(id string) (*models.Paste, error) {
	var paste models.Paste

	err := r.DB.First(&paste, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &paste, nil
}

func (r *PasteRepository) FindByUser(userID uuid.UUID) ([]models.Paste, error) {
	var pastes []models.Paste

	err := r.DB.Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&pastes).Error

	return pastes, err
}

func (r *PasteRepository) Delete(id string, userId uuid.UUID) error {
	result := r.DB.Where("id = ? AND user_id = ?", id, userId).
		Delete(&models.Paste{})

	return result.Error
}
