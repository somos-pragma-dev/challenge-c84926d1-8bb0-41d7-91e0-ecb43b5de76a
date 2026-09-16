package handlers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	dto "github.com/empresa/product-api/internal/dto"
	repository "github.com/empresa/product-api/internal/repository"
	"github.com/empresa/product-api/internal/service"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(repo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{
		service: service.NewProductService(repo),
	}
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Error al obtener productos",
			Message: err.Error(),
		})
		return
	}

	if len(products) == 0 {
		c.JSON(http.StatusOK, dto.ProductListResponse{
			Products: []dto.ProductResponse{},
			Total:    0,
		})
		return
	}

	var productResponses []dto.ProductResponse
	for _, p := range products {
		productResponses = append(productResponses, dto.ProductResponse{
			ID:       p.ID,
			Name:     p.Name,
			Price:    p.Price,
			Stock:    p.Stock,
			Category: p.Category,
		})
	}

	c.JSON(http.StatusOK, dto.ProductListResponse{
		Products: productResponses,
		Total:    len(productResponses),
	})
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "ID inválido",
			Message: "El ID debe ser un número entero positivo",
		})
		return
	}

	product, err := h.service.GetProductByID(uint(id))
	if err != nil {
		if err.Error() == "producto no encontrado" {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{
				Error:   "Producto no encontrado",
				Message: "No existe un producto con el ID proporcionado",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Error al obtener producto",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Stock:    product.Stock,
		Category: product.Category,
	})
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Datos inválidos",
			Message: formatValidationErrors(err),
		})
		return
	}

	product, err := h.service.CreateProduct(req)
	if err != nil {
		switch err.Error() {
		case "el nombre del producto ya existe":
			c.JSON(http.StatusConflict, dto.ErrorResponse{
				Error:   "Conflicto de datos",
				Message: err.Error(),
			})
		case "el precio no puede ser negativo":
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Error:   "Validación fallida",
				Message: err.Error(),
			})
		case "el nombre no puede estar vacío":
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Error:   "Validación fallida",
				Message: err.Error(),
			})
		default:
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Error:   "Error al crear producto",
				Message: err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusCreated, dto.ProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Stock:    product.Stock,
		Category: product.Category,
	})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "ID inválido",
			Message: "El ID debe ser un número entero positivo",
		})
		return
	}

	var req dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Datos inválidos",
			Message: formatValidationErrors(err),
		})
		return
	}

	product, err := h.service.UpdateProduct(uint(id), req)
	if err != nil {
		switch err.Error() {
		case "producto no encontrado":
			c.JSON(http.StatusNotFound, dto.ErrorResponse{
				Error:   "Producto no encontrado",
				Message: "No existe un producto con el ID proporcionado",
			})
		case "el nombre del producto ya existe":
			c.JSON(http.StatusConflict, dto.ErrorResponse{
				Error:   "Conflicto de datos",
				Message: err.Error(),
			})
		case "el precio no puede ser negativo":
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Error:   "Validación fallida",
				Message: err.Error(),
			})
		default:
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Error:   "Error al actualizar producto",
				Message: err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, dto.ProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Stock:    product.Stock,
		Category: product.Category,
	})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "ID inválido",
			Message: "El ID debe ser un número entero positivo",
		})
		return
	}

	err = h.service.DeleteProduct(uint(id))
	if err != nil {
		if err.Error() == "producto no encontrado" {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{
				Error:   "Producto no encontrado",
				Message: "No existe un producto con el ID proporcionado",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Error al eliminar producto",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Producto eliminado exitosamente",
	})
}

func formatValidationErrors(err error) string {
	if err == nil {
		return "Error de validación desconocido"
	}
	return err.Error()
}