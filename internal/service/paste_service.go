package service

import (
	"codeDrop/internal/models"
	"codeDrop/internal/repository"
	"codeDrop/internal/storage"
	"codeDrop/internal/utils"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type PasteService struct {
	Repo    *repository.PasteRepository
	Storage *storage.FileStorage
}

func NewPasteService(repo *repository.PasteRepository, storage *storage.FileStorage) *PasteService {
	return &PasteService{
		Repo: repo,
		Storage: storage,
	}
}

func (s *PasteService) Create(userID uuid.UUID, content string, visibility string) (*models.Paste, error) {
	id, err := utils.GenerateID()
	if err != nil {
		return nil, err
	}

	objectKey := fmt.Sprintf("pastes/%s.txt", id)

	err = s.Storage.Save(objectKey, content)
	if err != nil {
		return nil, err
	}

	paste := &models.Paste{
		ID: id,
		UserId: userID,
		ObjectKey: objectKey,
		Visibility: visibility,
		CreatedAt: time.Now(),
	}

	err = s.Repo.Create(paste)
	if err != nil {
		return nil, err
	}

	return paste, nil
}
