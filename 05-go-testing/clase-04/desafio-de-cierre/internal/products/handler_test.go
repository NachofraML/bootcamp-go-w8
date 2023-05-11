package products

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerGetProducts(t *testing.T) {
	t.Run("Successful response", func(t *testing.T) {
		// Arrange
		queryParamKey := "seller_id"
		queryParamValue := "1"
		expectedHTTPStatusCode := http.StatusOK
		expectedHTTPHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedResponse := `[
			{
				"id": "mock",
				"seller_id": "FEX112AC",
				"description": "generic product",
				"price": 123.55
			}
		]`
		responseFromService := []Product{
			{
				ID:          "mock",
				SellerID:    "FEX112AC",
				Description: "generic product",
				Price:       123.55,
			},
		}
		var errFromService error

		serviceMock := NewServiceMock()
		handler := NewHandler(serviceMock)

		serviceMock.
			On("GetAllBySeller", queryParamValue).
			Return(responseFromService, errFromService)

		router := gin.New()
		router.GET("/api/v1/products", handler.GetProducts)

		responseRecorder := httptest.NewRecorder()

		// Act
		router.ServeHTTP(responseRecorder, httptest.NewRequest(
			http.MethodGet,
			"/api/v1/products?"+queryParamKey+"="+queryParamValue,
			nil,
		))

		// Assert
		serviceMock.AssertExpectations(t)
		assert.Equal(t, expectedHTTPStatusCode, responseRecorder.Code)
		assert.Equal(t, expectedHTTPHeaders, responseRecorder.Header())
		assert.JSONEq(t, expectedResponse, responseRecorder.Body.String())
	})
	t.Run("Request without seller_id", func(t *testing.T) {
		// Arrange
		sellerIdToSearch := ""
		expectedHTTPStatusCode := http.StatusBadRequest
		expectedHTTPHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedResponse := `{
				"error": "seller_id query param is required"
		}`

		serviceMock := NewServiceMock()
		handler := NewHandler(serviceMock)

		router := gin.New()
		router.GET("/api/v1/products", handler.GetProducts)

		responseRecorder := httptest.NewRecorder()

		// Act
		router.ServeHTTP(responseRecorder, httptest.NewRequest(
			http.MethodGet,
			"/api/v1/products?"+sellerIdToSearch,
			nil,
		))

		// Assert
		assert.Equal(t, expectedHTTPStatusCode, responseRecorder.Code)
		assert.Equal(t, expectedHTTPHeaders, responseRecorder.Header())
		assert.JSONEq(t, expectedResponse, responseRecorder.Body.String())
	})
	t.Run("Unexpected error", func(t *testing.T) {
		// Arrange
		queryParamKey := "seller_id"
		queryParamValue := "1"
		expectedHTTPStatusCode := http.StatusInternalServerError
		expectedHTTPHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedResponse := `{
				"error": "something went wrong"
		}`

		var responseFromService []Product
		errFromService := fmt.Errorf("something went wrong")

		serviceMock := NewServiceMock()
		handler := NewHandler(serviceMock)

		serviceMock.
			On("GetAllBySeller", queryParamValue).
			Return(responseFromService, errFromService)

		router := gin.New()
		router.GET("/api/v1/products", handler.GetProducts)

		responseRecorder := httptest.NewRecorder()

		// Act
		router.ServeHTTP(responseRecorder, httptest.NewRequest(
			http.MethodGet,
			"/api/v1/products?"+queryParamKey+"="+queryParamValue,
			nil,
		))

		// Assert
		serviceMock.AssertExpectations(t)
		assert.Equal(t, expectedHTTPStatusCode, responseRecorder.Code)
		assert.Equal(t, expectedHTTPHeaders, responseRecorder.Header())
		assert.JSONEq(t, expectedResponse, responseRecorder.Body.String())
	})
}
