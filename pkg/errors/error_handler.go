package errors

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ValidationErrorResponse struct {
	Error   string            `json:"error"`
	Message string            `json:"message"`
	Code    int               `json:"code"`
	Details map[string]string `json:"details,omitempty"`
}

func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	errMsg := err.Error()
	log.Printf("[ERROR] %s", errMsg)

	switch {
	case strings.Contains(errMsg, "not found"):
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "NOT_FOUND",
			Message: errMsg,
			Code:    http.StatusNotFound,
		})
	case strings.Contains(errMsg, "ya existe"):
		c.JSON(http.StatusConflict, ErrorResponse{
			Error:   "CONFLICT",
			Message: errMsg,
			Code:    http.StatusConflict,
		})
	case strings.Contains(errMsg, "no puede estar vacío"):
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "BAD_REQUEST",
			Message: errMsg,
			Code:    http.StatusBadRequest,
		})
	case strings.Contains(errMsg, "no puede ser negativo"):
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "BAD_REQUEST",
			Message: errMsg,
			Code:    http.StatusBadRequest,
		})
	case strings.Contains(errMsg, "al menos"):
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "BAD_REQUEST",
			Message: errMsg,
			Code:    http.StatusBadRequest,
		})
	case strings.Contains(errMsg, "no puede exceder"):
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "BAD_REQUEST",
			Message: errMsg,
			Code:    http.StatusBadRequest,
		})
	case strings.Contains(errMsg, "no válida"):
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "BAD_REQUEST",
			Message: errMsg,
			Code:    http.StatusBadRequest,
		})
	case strings.Contains(errMsg, "error al verificar"):
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "INTERNAL_ERROR",
			Message: "Error al procesar la solicitud",
			Code:    http.StatusInternalServerError,
		})
	default:
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "INTERNAL_ERROR",
			Message: "Ha ocurrido un error interno. Por favor, intente más tarde.",
			Code:    http.StatusInternalServerError,
		})
	}
}

func HandleValidationErrors(c *gin.Context, err error) {
	errMsg := err.Error()
	log.Printf("[VALIDATION ERROR] %s", errMsg)

	details := parseValidationDetails(errMsg)

	response := ValidationErrorResponse{
		Error:   "VALIDATION_ERROR",
		Message: "Error de validación en los datos proporcionados",
		Code:    http.StatusBadRequest,
	}

	if len(details) > 0 {
		response.Details = details
	}

	c.JSON(http.StatusBadRequest, response)
}

func parseValidationDetails(errMsg string) map[string]string {
	details := make(map[string]string)

	errors := strings.Split(errMsg, "; ")
	for _, e := range errors {
		e = strings.TrimSpace(e)
		if e == "" {
			continue
		}

		if strings.Contains(e, "nombre") {
			details["name"] = e
		} else if strings.Contains(e, "precio") {
			details["price"] = e
		} else if strings.Contains(e, "stock") {
			details["stock"] = e
		} else if strings.Contains(e, "categoría") || strings.Contains(e, "categoria") {
			details["category"] = e
		} else {
			details["general"] = e
		}
	}

	return details
}

func HandleNotFound(c *gin.Context, resource string) {
	c.JSON(http.StatusNotFound, ErrorResponse{
		Error:   "NOT_FOUND",
		Message: fmt.Sprintf("%s no encontrado", resource),
		Code:    http.StatusNotFound,
	})
}

func HandleBadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, ErrorResponse{
		Error:   "BAD_REQUEST",
		Message: message,
		Code:    http.StatusBadRequest,
	})
}

func HandleUnauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, ErrorResponse{
		Error:   "UNAUTHORIZED",
		Message: "No autorizado para realizar esta acción",
		Code:    http.StatusUnauthorized,
	})
}

func HandleMethodNotAllowed(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, ErrorResponse{
		Error:   "METHOD_NOT_ALLOWED",
		Message: "Método no permitido para este recurso",
		Code:    http.StatusMethodNotAllowed,
	})
}

// GlobalErrorHandler returns a Gin middleware that handles panics
func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[PANIC] %v", err)
				c.JSON(http.StatusInternalServerError, ErrorResponse{
					Error:   "INTERNAL_ERROR",
					Message: "Se ha producido un error inesperado",
					Code:    http.StatusInternalServerError,
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}