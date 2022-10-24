package products

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	repository := NewRepository()
	service := NewService(repository)
	p := NewHandler(service)
	r := gin.Default()

	r.GET("/api/v1/products", p.GetProducts)
	return r
}

func createRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func TestGetProducts(t *testing.T) {
	// arrange
	r := createServer()
	sellerId := "FEX112AC"
	url := fmt.Sprintf("/api/v1/products?seller_id=%v", sellerId)
	req, rr := createRequestTest(http.MethodGet, url, "")

	// act
	r.ServeHTTP(rr, req)

	// assert
	assert.Equal(t, 200, rr.Code)
}

func TestGetProductsSellerInvalid(t *testing.T) {
	// arrange
	r := createServer()
	sellerId := ""
	url := fmt.Sprintf("/api/v1/products?seller_id=%v", sellerId)
	req, rr := createRequestTest(http.MethodGet, url, "")
	errExpect := map[string]string{
		"error": "seller_id query param is required",
	}
	var response map[string]string
	// act
	r.ServeHTTP(rr, req)

	// assert
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, 400, rr.Code)
	assert.Equal(t, errExpect, response)
}

// func TestGetProductsError500(t *testing.T) {
// 	// arrange
// 	r := createServer()
// 	sellerId := ""
// 	url := fmt.Sprintf("/api//products?seller_id=%v", sellerId)
// 	req, rr := createRequestTest(http.MethodGet, url, "")
// 	errExpect := map[string]string{
// 		"error": "error",
// 	}
// 	var response map[string]string

// 	// act
// 	r.ServeHTTP(rr, req)

// 	// assert
// 	err := json.Unmarshal(rr.Body.Bytes(), &response)
// 	assert.Nil(t, err)
// 	assert.Equal(t, 500, rr.Code)
// 	assert.Equal(t, errExpect, response)
// }
