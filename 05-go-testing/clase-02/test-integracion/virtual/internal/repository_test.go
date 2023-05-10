package internal

import (
	"encoding/json"
	"errors"
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-02/test-integracion/virtual/pkg/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllSuccesfull(t *testing.T) {
	//Arrange
	var users = []User{
		{
			ID:    1,
			Name:  "Juan",
			Email: "juan@emial.com",
			Age:   32,
		},
		{
			ID:    2,
			Name:  "Pepe",
			Email: "pepe@emial.com",
			Age:   17,
		},
	}
	dataJson, _ := json.Marshal(users)
	dbMock := store.Mock{
		Data: dataJson,
	}

	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	repo := NewRepository(storeStub)

	//Act
	myUsers, _ := repo.GetAll()

	//Assert
	assert.Equal(t, users, myUsers)
}

func TestGetAllFail(t *testing.T) {
	//Arrange
	var user = User{

		ID:    1,
		Name:  "Juan",
		Email: "juan@emial.com",
		Age:   32,
	}

	expectedError := errors.New("error for Storage")
	dbMock := store.Mock{
		Err: expectedError,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	rep := NewRepository(storeStub)

	//act
	result, err := rep.Store(user.ID, user.Name, user.Email, user.Age)

	//Assert
	assert.Equal(t, expectedError, err)
	assert.Equal(t, User{}, result)
}
