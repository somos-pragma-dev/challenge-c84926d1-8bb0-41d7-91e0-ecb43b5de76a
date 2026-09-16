package validators

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/empresa/product-api/internal/domain"
	"github.com/empresa/product-api/internal/repository"
)

type ProductValidator struct {
	repo *repository.ProductRepository
}

func NewProductValidator(repo *repository.ProductRepository) *ProductValidator {
	return &ProductValidator{repo: repo}
}

func (v *ProductValidator) ValidateProduct(product *domain.Product) error {
	validate := validator.New()

	err := validate.Struct(product)
	if err != nil {
		return v.formatValidationErrors(err)
	}

	if err := v.validateUniqueName(product); err != nil {
		return err
	}

	return nil
}

func (v *ProductValidator) validateUniqueName(product *domain.Product) error {
	exists, err := v.repo.ExistsByName(product.Name, product.ID)
	if err != nil {
		return fmt.Errorf("error al verificar nombre del producto: %w", err)
	}

	if exists {
		return fmt.Errorf("el nombre '%s' ya existe en la base de datos", product.Name)
	}

	return nil
}

func (v *ProductValidator) ValidateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("el nombre del producto no puede estar vacío")
	}

	if len(name) < 3 {
		return fmt.Errorf("el nombre del producto debe tener al menos 3 caracteres")
	}

	if len(name) > 100 {
		return fmt.Errorf("el nombre del producto no puede exceder 100 caracteres")
	}

	return nil
}

func (v *ProductValidator) ValidatePrice(price float64) error {
	if price < 0 {
		return fmt.Errorf("el precio no puede ser negativo")
	}

	if price > 999999.99 {
		return fmt.Errorf("el precio excede el límite permitido")
	}

	return nil
}

func (v *ProductValidator) ValidateStock(stock int) error {
	if stock < 0 {
		return fmt.Errorf("el stock no puede ser negativo")
	}

	if stock > 999999 {
		return fmt.Errorf("el stock excede el límite permitido")
	}

	return nil
}

func (v *ProductValidator) ValidateCategory(category string) error {
	validCategories := []string{"electronics", "clothing", "food", "books", "home", "sports", "toys", "other"}

	categoryLower := strings.ToLower(category)

	for _, valid := range validCategories {
		if categoryLower == valid {
			return nil
		}
	}

	return fmt.Errorf("categoría '%s' no válida. Las categorías válidas son: %s", category, strings.Join(validCategories, ", "))
}

func (v *ProductValidator) formatValidationErrors(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var errorMessages []string

		for _, fieldErr := range validationErrors {
			switch fieldErr.Field() {
			case "Name":
				switch fieldErr.Tag() {
				case "required":
					errorMessages = append(errorMessages, "El nombre del producto es obligatorio")
				case "min":
					errorMessages = append(errorMessages, "El nombre debe tener al menos 3 caracteres")
				case "max":
					errorMessages = append(errorMessages, "El nombre no puede exceder 100 caracteres")
				}
			case "Price":
				switch fieldErr.Tag() {
				case "required":
					errorMessages = append(errorMessages, "El precio es obligatorio")
				case "gte":
					errorMessages = append(errorMessages, "El precio debe ser mayor o igual a 0")
				}
			case "Stock":
				switch fieldErr.Tag() {
				case "required":
					errorMessages = append(errorMessages, "El stock es obligatorio")
				case "gte":
					errorMessages = append(errorMessages, "El stock debe ser mayor o igual a 0")
				}
			case "Category":
				switch fieldErr.Tag() {
				case "required":
					errorMessages = append(errorMessages, "La categoría es obligatoria")
				}
			default:
				errorMessages = append(errorMessages, fmt.Sprintf("Error en el campo %s", fieldErr.Field()))
			}
		}

		return strings.Join(errorMessages, "; ")
	}

	return err.Error()
}

func (v *ProductValidator) ValidateProductForUpdate(existingProduct *domain.Product, updatedProduct *domain.Product) error {
	validate := validator.New()

	err := validate.Struct(updatedProduct)
	if err != nil {
		return v.formatValidationErrors(err)
	}

	if updatedProduct.Name != existingProduct.Name {
		if err := v.validateUniqueName(updatedProduct); err != nil {
			return err
		}
	}

	return nil
}