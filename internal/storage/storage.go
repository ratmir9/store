package storage

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ratmir9/storage/internal/file"
)

type Storage struct {
	Files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		Files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, fileData []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, fileData)
	if err != nil {
		return nil, err
	}
	s.Files[newFile.ID] = newFile
	return newFile, nil
}

func (s *Storage) GetById(FileID uuid.UUID) (*file.File, error) {
	foundFile, ok := s.Files[FileID]
	if !ok {
		return nil, fmt.Errorf("File % not found", FileID)
	}
	return foundFile, nil
}
