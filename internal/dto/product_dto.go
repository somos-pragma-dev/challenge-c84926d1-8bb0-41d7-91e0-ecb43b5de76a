package dto

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/empresa/product-api/internal/domain"
	"github.com/empresa/product-api/internal/repository"
)

type ProductDTO struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
	Category string  `json:"category"`
}

type CreateProductRequest struct {
	Name     string  `json:"name" binding:"required,min=3,max=255"`
	Price    float64 `json:"price" binding:"required,gte=0"`
	Stock    int     `json:"stock" binding:"gte=0"`
	Category string  `json:"category" binding:"required,min=1,max=100"`
}

type UpdateProductRequest struct {
	Name     string  `json:"name" binding:"omitempty,min=3,max=255"`
	Price    float64 `json:"price" binding:"omitempty,gte=0"`
	Stock    int     `json:"stock" binding:"omitempty,gte=0"`
	Category string  `json:"category" binding:"omitempty,min=1,max=100"`
}

type ProductListResponse struct {
	Success bool        `json:"success"`
	Data    []ProductDTO `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

type ProductResponse struct {
	Success bool       `json:"success"`
	Data    ProductDTO `json:"data,omitempty"`
	Message string     `json:"message,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type ProductHandler struct {
	repository *repository.ProductRepository
}

func NewProductHandler(repo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repository: repo}
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.repository.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Error al obtener la lista de productos",
		})
		return
	}

	productDTOs := make([]ProductDTO, 0, len(products))
	for _, p := range products {
		productDTOs = append(productDTOs, ProductDTO{
			ID:       p.ID,
			Name:     p.Name,
			Price:    p.Price,
			Stock:    p.Stock,
			Category: p.Category,
		})
	}

	c.JSON(http.StatusOK, ProductListResponse{
		Success: true,
		Data:    productDTOs,
	})
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   "ID de producto inválido",
		})
		return
	}

	product, err := h.repository.FindByID(uint(id))
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, ErrorResponse{
				Success: false,
				Error:   "Producto no encontrado",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Error al obtener el producto",
		})
		return
	}

	c.JSON(http.StatusOK, ProductResponse{
		Success: true,
		Data: ProductDTO{
			ID:       product.ID,
			Name:     product.Name,
			Price:    product.Price,
			Stock:    product.Stock,
			Category: product.Category,
		},
	})
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := formatValidationErrors(err)
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   validationErrors,
		})
		return
	}

	product := &domain.Product{
		Name:     req.Name,
		Price:    req.Price,
		Stock:    req.Stock,
		Category: req.Category,
	}

	if err := h.repository.Create(product); err != nil {
		if errors.Is(err, repository.ErrProductAlreadyExists) {
			c.JSON(http.StatusConflict, ErrorResponse{
				Success: false,
				Error:   "Ya existe un producto con ese nombre",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Error al crear el producto",
		})
		return
	}

	c.JSON(http.StatusCreated, ProductResponse{
		Success: true,
		Data: ProductDTO{
			ID:       product.ID,
			Name:     product.Name,
			Price:    product.Price,
			Stock:    product.Stock,
			Category: product.Category,
		},
		Message: "Producto creado exitosamente",
	})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   "ID de producto inválido",
		})
		return
	}

	existingProduct, err := h.repository.FindByID(uint(id))
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, ErrorResponse{
				Success: false,
				Error:   "Producto no encontrado",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Error al obtener el producto",
		})
		return
	}

	var req UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := formatValidationErrors(err)
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   validationErrors,
		})
		return
	}

	if req.Name != "" {
		existingProduct.Name = req.Name
	}
	if req.Price != 0 {
		existingProduct.Price = req.Price
	}
	if req.Stock != 0 {
		existingProduct.Stock = req.Stock
	}
	if req.Category != "" {
		existingProduct.Category = req.Category
	}

	if err := h.repository.Update(existingProduct); err != nil {
		if errors.Is(err, repository.ErrProductAlreadyExists) {
			c.JSON(http.StatusConflict, ErrorResponse{
				Success: false,
				Error:   "Ya existe un producto con ese nombre",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Error al actualizar el producto",
		})
		return
	}

	c.JSON(http.StatusOK, ProductResponse{
		Success: true,
		Data: ProductDTO{
			ID:       existingProduct.ID,
			Name:     existingProduct.Name,
			Price:    existingProduct.Price,
			Stock:    existingProduct.Stock,
			Category: existingProduct.Category,
		},
		Message: "Producto actualizado exitosamente",
	})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   "ID de producto inválido",
		})
		return
	}

	if err := h.repository.Delete(uint(id)); err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, ErrorResponse{
				Success: false,
				Error:   "Producto no encontrado",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Error al eliminar el producto",
		})
		return
	}

	c.JSON(http.StatusOK, ProductResponse{
		Success: true,
		Message: "Producto eliminado exitosamente",
	})
}

func formatValidationErrors(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var errorMessages []string
		for _, e := range validationErrors {
			switch e.Field() {
			case "Name":
				switch e.Tag() {
				case "required":
					errorMessages = append(errorMessages, "El nombre es requerido")
				case "min":
					errorMessages = append(errorMessages, "El nombre debe tener al menos 3 caracteres")
				case "max":
					errorMessages = append(errorMessages, "El nombre no puede exceder 255 caracteres")
				}
			case "Price":
				switch e.Tag() {
				case "required":
					errorMessages = append(errorMessages, "El precio es requerido")
				case "gte":
					errorMessages = append(errorMessages, "El precio no puede ser negativo")
				}
			case "Stock":
				switch e.Tag() {
				case "gte":
					errorMessages = append(errorMessages, "El stock no puede ser negativo")
				}
			case "Category":
				switch e.Tag() {
				case "required":
					errorMessages = append(errorMessages, "La categoría es requerida")
				case "min":
					errorMessages = append(errorMessages, "La categoría debe tener al menos 1 caracter")
				case "max":
					errorMessages = append(errorMessages, "La categoría no puede exceder 100 caracteres")
				}
			}
		}
		if len(errorMessages) > 0 {
			return errorMessages[0]
		}
	}
	return "Error de validación"
}