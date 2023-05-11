package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	ErrReadingFile  = fmt.Errorf("error reading file")
	ErrParseJSON    = fmt.Errorf("error parsing json")
	ErrKeyNotExists = fmt.Errorf("error key not exists")
)

type Storage interface {
	GetValue(key string) interface{}
}

type storage struct {
	file string
}

func (s *storage) GetValue(key string) interface{} {
	file, err := os.ReadFile(s.file)
	if err != nil {
		return ErrReadingFile
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(file, &data)
	if err != nil {
		return ErrParseJSON
	}

	if v, ok := data[key]; ok {
		return v
	}

	return ErrKeyNotExists
}

func NewStorage() Storage {
	file := "config.json"
	return &storage{file: file}
}
