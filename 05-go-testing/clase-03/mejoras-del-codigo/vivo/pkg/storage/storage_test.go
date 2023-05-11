package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorageGetValue(t *testing.T) {
	t.Run("Successful read", func(t *testing.T) {
		// Arrange
		expectedResponse := 144.0
		storage := NewStorage()

		// Act
		value := storage.GetValue("white_shark_speed")

		// Assert
		assert.Equal(t, value, expectedResponse)
	})
	t.Run("File does not exists", func(t *testing.T) {
		// Arrange
		expectedResponse := ErrReadingFile
		storage := storage{file: "nothing"}

		// Act
		value := storage.GetValue("white_shark_speed")

		// Assert
		assert.Error(t, value.(error), expectedResponse)
		assert.EqualError(t, value.(error), expectedResponse.Error())
	})
	t.Run("Invalid config JSON", func(t *testing.T) {
		// Arrange
		expectedResponse := ErrParseJSON
		storage := storage{file: "invalid_config.json"}

		// Act
		value := storage.GetValue("white_shark_speed")

		// Assert
		assert.Error(t, value.(error), expectedResponse)
		assert.EqualError(t, value.(error), expectedResponse.Error())
	})
	t.Run("No value in JSON Key", func(t *testing.T) {
		// Arrange
		expectedResponse := ErrKeyNotExists
		storage := NewStorage()

		// Act
		value := storage.GetValue("null_data")

		// Assert
		assert.Error(t, value.(error), expectedResponse)
		assert.EqualError(t, value.(error), expectedResponse.Error())
	})
}
