package products

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServiceGetAllBySeller(t *testing.T) {
	t.Run("Successful GetAll", func(t *testing.T) {
		// Arrange
		mockedRepo := NewRepositoryMock()
		service := NewService(mockedRepo)

		idToSearch := "mock"
		expectedResponse := []Product{
			{
				ID:          "mock",
				SellerID:    "FEX112AC",
				Description: "generic product",
				Price:       123.55,
			},
		}

		mockedRepo.
			On("GetAllBySeller", idToSearch).
			Return(expectedResponse, nil)

		// Act
		seller, err := service.GetAllBySeller(idToSearch)

		// Assert
		mockedRepo.AssertExpectations(t)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, seller)
	})
	t.Run("Unexpected error from Repository", func(t *testing.T) {
		// Arrange
		mockedRepo := NewRepositoryMock()
		service := NewService(mockedRepo)

		idToSearch := ""

		mockedRepo.
			On("GetAllBySeller", idToSearch).
			Return([]Product{}, fmt.Errorf("something went wrong"))

		// Act
		seller, err := service.GetAllBySeller(idToSearch)

		// Assert
		mockedRepo.AssertExpectations(t)
		assert.Error(t, err)
		assert.Nil(t, seller)
		assert.EqualError(t, err, ErrMockedRepo.Error())
	})
}
