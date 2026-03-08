package storage

import (
	"os"
	"path/filepath"
)

type FileStorage struct {
	BasePath string
}

func NewFileStorage(base string) *FileStorage {
	return &FileStorage{
		BasePath: base,
	}
}

func (s *FileStorage) Save(key string, content string) error {

	path := filepath.Join(s.BasePath, key)

	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(content)

	return err
}

func (s *FileStorage) Read(key string) (string, error) {

	path := filepath.Join(s.BasePath, key)

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (s *FileStorage) Delete(key string) error {

	path := filepath.Join(s.BasePath, key)

	return os.Remove(path)
}