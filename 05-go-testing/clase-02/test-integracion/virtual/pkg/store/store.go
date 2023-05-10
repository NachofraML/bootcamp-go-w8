package store

import (
	"encoding/json"
	"io/ioutil"
)

type FileStore struct {
	FileName string
	Mock     *Mock
}

type Mock struct {
	Data []byte
	Err  error
}

func (fs *FileStore) AddMock(mock *Mock) {
	fs.Mock = mock
}

func (fs *FileStore) ClearMock() {
	fs.Mock = nil
}

func (fs *FileStore) Read(data interface{}) error {
	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}
		return json.Unmarshal(fs.Mock.Data, data)
	}
	file, err := ioutil.ReadFile(fs.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, data)
}

func (fs *FileStore) Write(data interface{}) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}
		fs.Mock.Data = file
		return nil
	}

	return ioutil.WriteFile(fs.FileName, file, 0644)
}
