package tests

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHandlerIntegration(t *testing.T) {
	t.Run("Prey escapes hunter", func(t *testing.T) {
		// Arrange
		expectedHTTPStatusCode := http.StatusOK
		expectedHTTPHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		whiteSharkRequestBody := `{
				"x_position": 10.0,
				"y_position": 10.0,
				"speed": 10.0
		}`
		tunaRequestBody := `{
				"speed": 11.0
		}`
		expectedConfigResponse := `{
				"success": true
		}`
		expectedHuntResponse := `{
				"success": false,
				"message": "prey is faster than hunter, cannot hunt",
				"time": 0
		}`

		router := createServer()
		whiteSharkRequest, whiteSharkResponseRecorder := createRequestTest(http.MethodPut, "/v1/shark", whiteSharkRequestBody)
		tunaRequest, tunaResponseRecorder := createRequestTest(http.MethodPut, "/v1/prey", tunaRequestBody)
		huntRequest, huntResponseRecorder := createRequestTest(http.MethodPost, "/v1/simulate", "")

		// Act
		router.ServeHTTP(whiteSharkResponseRecorder, whiteSharkRequest)
		router.ServeHTTP(tunaResponseRecorder, tunaRequest)
		router.ServeHTTP(huntResponseRecorder, huntRequest)

		// Assert
		assert.Equal(t, expectedHTTPStatusCode, whiteSharkResponseRecorder.Code)
		assert.Equal(t, expectedHTTPStatusCode, tunaResponseRecorder.Code)
		assert.Equal(t, expectedHTTPStatusCode, huntResponseRecorder.Code)

		assert.Equal(t, expectedHTTPHeaders, whiteSharkResponseRecorder.Header())
		assert.Equal(t, expectedHTTPHeaders, tunaResponseRecorder.Header())
		assert.Equal(t, expectedHTTPHeaders, huntResponseRecorder.Header())

		assert.JSONEq(t, expectedConfigResponse, whiteSharkResponseRecorder.Body.String())
		assert.JSONEq(t, expectedConfigResponse, tunaResponseRecorder.Body.String())
		assert.JSONEq(t, expectedHuntResponse, huntResponseRecorder.Body.String())
	})
	t.Run("Hunter is too far", func(t *testing.T) {
		// Arrange
		expectedHTTPStatusCode := http.StatusOK
		expectedHTTPHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		whiteSharkRequestBody := `{
				"x_position": 400.0,
				"y_position": 0.0,
				"speed": 11.0
		}`
		tunaRequestBody := `{
				"speed": 10.0
		}`
		expectedConfigResponse := `{
				"success": true
		}`
		expectedHuntResponse := `{
				"success": false,
				"message": "prey is too far, cannot hunt",
				"time": 0
		}`

		router := createServer()
		whiteSharkRequest, whiteSharkResponseRecorder := createRequestTest(http.MethodPut, "/v1/shark", whiteSharkRequestBody)
		tunaRequest, tunaResponseRecorder := createRequestTest(http.MethodPut, "/v1/prey", tunaRequestBody)
		huntRequest, huntResponseRecorder := createRequestTest(http.MethodPost, "/v1/simulate", "")

		// Act
		router.ServeHTTP(whiteSharkResponseRecorder, whiteSharkRequest)
		router.ServeHTTP(tunaResponseRecorder, tunaRequest)
		router.ServeHTTP(huntResponseRecorder, huntRequest)

		// Assert
		assert.Equal(t, expectedHTTPStatusCode, whiteSharkResponseRecorder.Code)
		assert.Equal(t, expectedHTTPStatusCode, tunaResponseRecorder.Code)
		assert.Equal(t, expectedHTTPStatusCode, huntResponseRecorder.Code)

		assert.Equal(t, expectedHTTPHeaders, whiteSharkResponseRecorder.Header())
		assert.Equal(t, expectedHTTPHeaders, tunaResponseRecorder.Header())
		assert.Equal(t, expectedHTTPHeaders, huntResponseRecorder.Header())

		assert.JSONEq(t, expectedConfigResponse, whiteSharkResponseRecorder.Body.String())
		assert.JSONEq(t, expectedConfigResponse, tunaResponseRecorder.Body.String())
		assert.JSONEq(t, expectedHuntResponse, huntResponseRecorder.Body.String())
	})
	t.Run("Hunter hunts at least after 24 seconds", func(t *testing.T) {
		// Arrange
		expectedHTTPStatusCode := http.StatusOK
		expectedHTTPHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		whiteSharkRequestBody := `{
				"x_position": 25.0,
				"y_position": 0.0,
				"speed": 11.0
		}`
		tunaRequestBody := `{
				"speed": 10.0
		}`
		expectedConfigResponse := `{
				"success": true
		}`
		expectedHuntResponse := `{
				"success": true,
				"message": "the prey has been hunted",
				"time": 25.0
		}`

		router := createServer()
		whiteSharkRequest, whiteSharkResponseRecorder := createRequestTest(http.MethodPut, "/v1/shark", whiteSharkRequestBody)
		tunaRequest, tunaResponseRecorder := createRequestTest(http.MethodPut, "/v1/prey", tunaRequestBody)
		huntRequest, huntResponseRecorder := createRequestTest(http.MethodPost, "/v1/simulate", "")

		// Act
		router.ServeHTTP(whiteSharkResponseRecorder, whiteSharkRequest)
		router.ServeHTTP(tunaResponseRecorder, tunaRequest)
		router.ServeHTTP(huntResponseRecorder, huntRequest)

		// Assert
		assert.Equal(t, expectedHTTPStatusCode, whiteSharkResponseRecorder.Code)
		assert.Equal(t, expectedHTTPStatusCode, tunaResponseRecorder.Code)
		assert.Equal(t, expectedHTTPStatusCode, huntResponseRecorder.Code)

		assert.Equal(t, expectedHTTPHeaders, whiteSharkResponseRecorder.Header())
		assert.Equal(t, expectedHTTPHeaders, tunaResponseRecorder.Header())
		assert.Equal(t, expectedHTTPHeaders, huntResponseRecorder.Header())

		assert.JSONEq(t, expectedConfigResponse, whiteSharkResponseRecorder.Body.String())
		assert.JSONEq(t, expectedConfigResponse, tunaResponseRecorder.Body.String())
		assert.JSONEq(t, expectedHuntResponse, huntResponseRecorder.Body.String())
	})
}

func TestConfigurationHandlers(t *testing.T) {
	t.Run("ConfigureShark invalid types", func(t *testing.T) {
		// Arrange
		expectedHTTPStatusCode := http.StatusUnprocessableEntity
		expectedHTTPHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		whiteSharkRequestBody := `{
				"x_position": "invalid_type",
				"y_position": 10.0,
				"speed": 10.0
		}`
		expectedConfigResponse := `{
				"message": "invalid request"
		}`

		router := createServer()
		whiteSharkRequest, whiteSharkResponseRecorder := createRequestTest(http.MethodPut, "/v1/shark", whiteSharkRequestBody)

		// Act
		router.ServeHTTP(whiteSharkResponseRecorder, whiteSharkRequest)

		// Assert
		assert.Equal(t, expectedHTTPStatusCode, whiteSharkResponseRecorder.Code)

		assert.Equal(t, expectedHTTPHeaders, whiteSharkResponseRecorder.Header())

		assert.JSONEq(t, expectedConfigResponse, whiteSharkResponseRecorder.Body.String())
	})
	t.Run("ConfigurePrey invalid types", func(t *testing.T) {
		// Arrange
		expectedHTTPStatusCode := http.StatusUnprocessableEntity
		expectedHTTPHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		tunaRequestBody := `{
				"speed": "invalid_type"
		}`
		expectedConfigResponse := `{
				"message": "invalid request"
		}`

		router := createServer()
		tunaRequest, tunaResponseRecorder := createRequestTest(http.MethodPut, "/v1/prey", tunaRequestBody)

		// Act
		router.ServeHTTP(tunaResponseRecorder, tunaRequest)

		// Assert
		assert.Equal(t, expectedHTTPStatusCode, tunaResponseRecorder.Code)

		assert.Equal(t, expectedHTTPHeaders, tunaResponseRecorder.Header())

		assert.JSONEq(t, expectedConfigResponse, tunaResponseRecorder.Body.String())
	})
}
