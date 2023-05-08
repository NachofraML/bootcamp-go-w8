package warehouses

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"vivo/mocks"
)

func Test_sqlGetAll(t *testing.T) {
	//arrange
	dbMock, _ := mocks.InitDb()
	repoMock := NewMySqlRepositoty(dbMock)
	firstWarehouse := Warehouse{
		ID:        1,
		Name:      "Main Warehouse",
		Address:   "221 Baker Street",
		Telephone: "4555666",
		Capacity:  100,
	}
	warehouseQuantity := 2

	//act
	warehouses, err := repoMock.GetAll()
	//assert
	assert.NoError(t, err)
	assert.Equal(t, warehouseQuantity, len(warehouses))
	assert.Equal(t, firstWarehouse.Name, warehouses[0].Name)
}

func Test_sqlGet(t *testing.T) {
	//arrange
	dbMock, err := mocks.InitDb()
	repoMock := NewMySqlRepositoty(dbMock)
	firstWarehouse := Warehouse{
		ID:        1,
		Name:      "Main Warehouse",
		Address:   "221 Baker Street",
		Telephone: "4555666",
		Capacity:  100,
	}

	//act
	warehouse, err := repoMock.Get(firstWarehouse.ID)

	//assert
	assert.NoError(t, err)
	assert.Equal(t, firstWarehouse.Name, warehouse.Name)
	assert.Equal(t, firstWarehouse.ID, warehouse.ID)
}

func Test_sqlCreate(t *testing.T) {
	//arrange
	dbMock, _ := mocks.InitDb()
	repoMock := NewMySqlRepositoty(dbMock)
	warehouseToCreate := Warehouse{
		Name:      "Corn Flakes",
		Address:   "221 Baker Street",
		Telephone: "123456789",
		Capacity:  200,
	}

	//act
	err := repoMock.Create(&warehouseToCreate)
	//assert
	assert.NoError(t, err)
	assert.Equal(t, 10, warehouseToCreate.ID)
}
