package products

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"vivo/mocks"
)

func Test_sqlGetAll(t *testing.T) {
	//arrange
	dbMock, _ := mocks.InitDb()
	repoMock := NewMySqlRepositoty(dbMock)
	firstProduct := Product{
		ID:          1,
		Name:        "Corn Shoots",
		Quantity:    244,
		CodeValue:   "0009-1111",
		IsPublished: false,
		Expiration:  "2022-01-08",
		Price:       23.27,
		WarehouseId: 1,
	}
	productQuantity := 220

	//act
	products, err := repoMock.GetAll()
	//assert
	assert.NoError(t, err)
	assert.Equal(t, productQuantity, len(products))
	assert.Equal(t, firstProduct.Name, products[0].Name)
}
