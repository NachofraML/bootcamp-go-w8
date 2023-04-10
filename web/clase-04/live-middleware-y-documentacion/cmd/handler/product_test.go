package handler

import (
	"bytes"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-04/live-middleware-y-documentacion/internal/product"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func createServer() *gin.Engine {
	_ = godotenv.Load("../../.env")
	jsonStoragePath := os.Getenv("JSON_STORAGE_PATH_TEST")
	rp := product.NewJsonStorageRepository(jsonStoragePath)
	s := product.NewDefaultService(rp)
	cr := NewProductController(s)

	sv := gin.Default()

	p := sv.Group("/products")
	{
		p.GET("", cr.GetAll())
		p.GET("/:id", cr.GetId())
		p.POST("", cr.Save())
		p.DELETE("/:id", cr.Delete())
	}
	return sv
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "12345")

	return req, httptest.NewRecorder()
}

func Test_GetAllProducts_OK(t *testing.T) {
	sv := createServer()
	var (
		expectedResponse = `{
		  "data": [
			{
				"code_value":"1A",
				"expiration":"", 
				"id":1, 
				"is_published":false, 
				"name":"Producto de ejemplo", 
				"price":9.99, 
				"quantity":0
	        }
		  ],
		  "message": "success"
		}`
	)

	req, rec := createRequestTest(http.MethodGet, "/products", "")
	sv.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
	assert.JSONEq(t, expectedResponse, rec.Body.String())
}

func Test_GetProductByID_OK(t *testing.T) {
	sv := createServer()
	var (
		expectedResponse = `{
			"data":{
				"code_value":"1A",
				"expiration":"", 
				"id":1, 
				"is_published":false, 
				"name":"Producto de ejemplo", 
				"price":9.99, 
				"quantity":0
			}, 
			"message":"success"
		}`
	)

	req, rec := createRequestTest(http.MethodGet, "/products/1", "")
	sv.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
	assert.JSONEq(t, expectedResponse, rec.Body.String())
}

func Test_SaveProduct_OK(t *testing.T) {
	sv := createServer()
	req, rec := createRequestTest(http.MethodPost, "/products", `{
            "name":"test",
            "quantity":5,
            "code_value":"1D",
            "is_published":false,
            "expiration":"28/10/2020",
			"price":2.5
	}`)

	sv.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
}

func Test_DeleteProduct_OK(t *testing.T) {
	sv := createServer()
	req, rec := createRequestTest(http.MethodDelete, "/products/2", "")

	sv.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func Test_UpdateProductById_BAD_REQUEST(t *testing.T) {
	sv := createServer()
	req, rec := createRequestTest(http.MethodDelete, "/products/fdfasfdf", "")

	sv.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func Test_SaveProductById_UNAUTHORIZED(t *testing.T) {
	sv := createServer()
	req, rec := createRequestTest(http.MethodPost, "/products", "")
	req.Header.Set("token", "invalid_token")

	sv.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func Test_GetProductById_NOT_FOUND(t *testing.T) {
	sv := createServer()
	req, rec := createRequestTest(http.MethodGet, "/products/12345", "")

	sv.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}
