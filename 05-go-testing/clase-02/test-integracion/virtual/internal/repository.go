package internal

import (
	"fmt"
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-02/test-integracion/virtual/pkg/store"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"Age"`
}

var users []User

type Repository interface {
	GetAll() ([]User, error)
	Store(id int, name string, email string, age int) (User, error)
}

type repository struct {
	db store.FileStore
}

func NewRepository(db store.FileStore) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]User, error) {
	err := r.db.Read(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *repository) Store(id int, nombre string, email string, age int) (User, error) {
	repo.db.Read(&users)

	usr := User{id, nombre, email, age}

	users = append(users, usr)
	err := repo.db.Write(users)

	if err != nil {
		return User{}, err
	}

	return usr, nil
}

func (r *repository) Update(id int, name string, email string, age int) (User, error) {
	err := r.db.Read(&users)

	if err != nil {
		return User{}, err
	}

	usr := User{id, name, email, age}
	for i, v := range users {
		if v.ID == id {
			users[i] = usr
			err := r.db.Write(users)
			if err != nil {
				return User{}, err
			}
			return usr, nil
		}
	}
	return User{}, fmt.Errorf("The user %d doesn't exists", id)

}

func (repo *repository) Delete(id int) error {
	err := repo.db.Read(&users)
	if err != nil {
		return err
	}

	index := 0
	for i, v := range users {
		if v.ID == id {
			index = i
			users = append(users[:index], users[index+1:]...)
			err := repo.db.Write(users)

			return err
		}
	}
	return fmt.Errorf("The user %d doesn't exists", id)

}
