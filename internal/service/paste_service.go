package service

import (
	"codeDrop/internal/models"
	"codeDrop/internal/repository"
	"codeDrop/internal/storage"
	"codeDrop/internal/utils"
	"errors"
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

func (s *PasteService) FindById(id string) (string, error) {

	paste, err := s.Repo.FindById(id)
	if err != nil {
		return "", err
	}

	if paste.ExpiresAt != nil && time.Now().After(*paste.ExpiresAt) {
		return "", errors.New("paste expired")
	}
	content, err := s.Storage.Read(paste.ObjectKey)
	if err != nil {
		return "", err
	}

	return content, nil
}

func (s *PasteService) FindByUser(user_id uuid.UUID) ([]models.Paste, error) {
	return s.Repo.FindByUser(user_id)
}


func (s *PasteService) DeletePaste(id string, user_id uuid.UUID) error {
	paste, err := s.Repo.FindById(id)
	if err != nil {
		return err
	}

	if paste.UserId != user_id {
		return errors.New("unauthorized")
	}

	err = s.Storage.Delete(paste.ObjectKey)
	if err != nil {
		return err
	}

	return s.Repo.Delete(id, user_id)
}