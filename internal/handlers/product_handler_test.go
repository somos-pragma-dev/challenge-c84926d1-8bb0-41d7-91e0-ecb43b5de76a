package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	dto "github.com/empresa/product-api/internal/dto"
)

func init() {
	gin.SetMode(gin.TestMode)
}

type MockProductService struct {
	products []dto.ProductDTO
	err      error
}

func (m *MockProductService) GetAllProducts() ([]dto.ProductDTO, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.products, nil
}

func (m *MockProductService) GetProductByID(id uint) (dto.ProductDTO, error) {
	if m.err != nil {
		return dto.ProductDTO{}, m.err
	}
	for _, p := range m.products {
		if p.ID == id {
			return p, nil
		}
	}
	return dto.ProductDTO{}, errors.New("producto no encontrado")
}

func (m *MockProductService) CreateProduct(req dto.CreateProductRequest) (dto.ProductDTO, error) {
	if m.err != nil {
		return dto.ProductDTO{}, m.err
	}
	return dto.ProductDTO{
		ID:       1,
		Name:     req.Name,
		Price:    req.Price,
		Stock:    req.Stock,
		Category: req.Category,
	}, nil
}

func (m *MockProductService) UpdateProduct(id uint, req dto.UpdateProductRequest) (dto.ProductDTO, error) {
	if m.err != nil {
		return dto.ProductDTO{}, m.err
	}
	for _, p := range m.products {
		if p.ID == id {
			result := p
			if req.Name != nil {
				result.Name = *req.Name
			}
			if req.Price != nil {
				result.Price = *req.Price
			}
			if req.Stock != nil {
				result.Stock = *req.Stock
			}
			if req.Category != nil {
				result.Category = *req.Category
			}
			return result, nil
		}
	}
	return dto.ProductDTO{}, errors.New("producto no encontrado")
}

func (m *MockProductService) DeleteProduct(id uint) error {
	if m.err != nil {
		return m.err
	}
	for _, p := range m.products {
		if p.ID == id {
			return nil
		}
	}
	return errors.New("producto no encontrado")
}

func setupTestRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	return router
}

func TestListProducts_Success(t *testing.T) {
	router := setupTestRouter()

	mockService := &MockProductService{
		products: []dto.ProductDTO{
			{ID: 1, Name: "Producto 1", Price: 100.50, Stock: 10, Category: "Electrónica"},
			{ID: 2, Name: "Producto 2", Price: 200.00, Stock: 5, Category: "Ropa"},
		},
	}

	handler := &ProductHandler{service: &ProductServiceMock{mockService}}
	router.GET("/products", handler.ListProducts)

	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response dto.ProductListResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshaling response: %v", err)
	}

	if response.Total != 2 {
		t.Errorf("Expected 2 products, got %d", response.Total)
	}

	if len(response.Products) != 2 {
		t.Errorf("Expected 2 products in array, got %d", len(response.Products))
	}
}

func TestListProducts_Empty(t *testing.T) {
	router := setupTestRouter()

	mockService := &MockProductService{
		products: []dto.ProductDTO{},
	}

	handler := &ProductHandler{service: &ProductServiceMock{mockService}}
	router.GET("/products", handler.ListProducts)

	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response dto.ProductListResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshaling response: %v", err)
	}

	if response.Total != 0 {
		t.Errorf("Expected 0 products, got %d", response.Total)
	}
}

func TestGetProduct_Success(t *testing.T) {
	router := setupTestRouter()

	mockService := &MockProductService{
		products: []dto.ProductDTO{
			{ID: 1, Name: "Producto 1", Price: 100.50, Stock: 10, Category: "Electrónica"},
		},
	}

	handler := &ProductHandler{service: &ProductServiceMock{mockService}}
	router.GET("/products/:id", handler.GetProduct)

	req, _ := http.NewRequest("GET", "/products/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response dto.ProductResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshaling response: %v", err)
	}

	if response.Name != "Producto 1" {
		t.Errorf("Expected product name 'Producto 1', got '%s'", response.Name)
	}
}

func TestGetProduct_InvalidID(t *testing.T) {
	router := setupTestRouter()

	handler := &ProductHandler{service: &ProductServiceMock{}}
	router.GET("/products/:id", handler.GetProduct)

	req, _ := http.NewRequest("GET", "/products/abc", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}

	var response dto.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshaling response: %v", err)
	}

	if response.Error != "ID inválido" {
		t.Errorf("Expected error message 'ID inválido', got '%s'", response.Error)
	}
}

func TestCreateProduct_Success(t *testing.T) {
	router := setupTestRouter()

	mockService := &MockProductService{}

	handler := &ProductHandler{service: &ProductServiceMock{mockService}}
	router.POST("/products", handler.CreateProduct)

	body := dto.CreateProductRequest{
		Name:     "Nuevo Producto",
		Price:    150.00,
		Stock:    20,
		Category: "Hogar",
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}

	var response dto.ProductResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshaling response: %v", err)
	}

	if response.Name != "Nuevo Producto" {
		t.Errorf("Expected product name 'Nuevo Producto', got '%s'", response.Name)
	}
}

func TestCreateProduct_InvalidJSON(t *testing.T) {
	router := setupTestRouter()

	handler := &ProductHandler{service: &ProductServiceMock{}}
	router.POST("/products", handler.CreateProduct)

	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer([]byte("{invalid")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestCreateProduct_DuplicateName(t *testing.T) {
	router := setupTestRouter()

	mockService := &MockProductService{
		err: errors.New("el nombre del producto ya existe"),
	}

	handler := &ProductHandler{service: &ProductServiceMock{mockService}}
	router.POST("/products", handler.CreateProduct)

	body := dto.CreateProductRequest{
		Name:     "Producto Existente",
		Price:    100.00,
		Stock:    10,
		Category: "Electrónica",
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusConflict {
		t.Errorf("Expected status 409, got %d", w.Code)
	}

	var response dto.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshaling response: %v", err)
	}

	if response.Error != "Conflicto de datos" {
		t.Errorf("Expected error 'Conflicto de datos', got '%s'", response.Error)
	}
}

func TestDeleteProduct_Success(t *testing.T) {
	router := setupTestRouter()

	mockService := &MockProductService{
		products: []dto.ProductDTO{
			{ID: 1, Name: "Producto 1", Price: 100.50, Stock: 10, Category: "Electrónica"},
		},
	}

	handler := &ProductHandler{service: &ProductServiceMock{mockService}}
	router.DELETE("/products/:id", handler.DeleteProduct)

	req, _ := http.NewRequest("DELETE", "/products/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestDeleteProduct_NotFound(t *testing.T) {
	router := setupTestRouter()

	mockService := &MockProductService{
		err: errors.New("producto no encontrado"),
	}

	handler := &ProductHandler{service: &ProductServiceMock{mockService}}
	router.DELETE("/products/:id", handler.DeleteProduct)

	req, _ := http.NewRequest("DELETE", "/products/999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}