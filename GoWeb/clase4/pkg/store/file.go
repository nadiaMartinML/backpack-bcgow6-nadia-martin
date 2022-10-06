package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	ReadFile(data interface{}) error
	WriteFile(data interface{}) error
}

type Type string

const (
	FileType Type = "file"
)

func New(store Type, fileName string) Store {
	return &fileStore{fileName}
}

type fileStore struct {
	FilePath string
}

func (fs *fileStore) WriteFile(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, fileData, 0644)
}

func (fs *fileStore) ReadFile(data interface{}) error {
	file, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}
