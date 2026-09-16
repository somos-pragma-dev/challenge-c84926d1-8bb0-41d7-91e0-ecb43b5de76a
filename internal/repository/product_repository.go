package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/empresa/product-api/internal/domain"
)

var (
	ErrProductNotFound       = errors.New("producto no encontrado")
	ErrProductAlreadyExists = errors.New("el producto ya existe")
	ErrDatabaseOperation    = errors.New("error en la operación de base de datos")
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *domain.Product) error {
	if err := product.Validate(); err != nil {
		return fmt.Errorf("validación fallida: %w", err)
	}

	exists, err := r.ExistsByName(product.Name, 0)
	if err != nil {
		return err
	}
	if exists {
		return ErrProductAlreadyExists
	}

	result := r.db.Create(product)
	if result.Error != nil {
		return fmt.Errorf("%w: %v", ErrDatabaseOperation, result.Error)
	}

	return nil
}

func (r *ProductRepository) FindByID(id uint) (*domain.Product, error) {
	var product domain.Product

	result := r.db.First(&product, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrDatabaseOperation, result.Error)
	}

	return &product, nil
}

func (r *ProductRepository) FindAll() ([]domain.Product, error) {
	var products []domain.Product

	result := r.db.Find(&products)
	if result.Error != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabaseOperation, result.Error)
	}

	return products, nil
}

func (r *ProductRepository) FindByName(name string) (*domain.Product, error) {
	var product domain.Product

	result := r.db.Where("name = ?", name).First(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrDatabaseOperation, result.Error)
	}

	return &product, nil
}

func (r *ProductRepository) FindByCategory(category string) ([]domain.Product, error) {
	var products []domain.Product

	result := r.db.Where("category = ?", category).Find(&products)
	if result.Error != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabaseOperation, result.Error)
	}

	return products, nil
}

func (r *ProductRepository) Update(product *domain.Product) error {
	if err := product.Validate(); err != nil {
		return fmt.Errorf("validación fallida: %w", err)
	}

	exists, err := r.ExistsByName(product.Name, product.ID)
	if err != nil {
		return err
	}
	if exists {
		return ErrProductAlreadyExists
	}

	result := r.db.Save(product)
	if result.Error != nil {
		return fmt.Errorf("%w: %v", ErrDatabaseOperation, result.Error)
	}

	if result.RowsAffected == 0 {
		return ErrProductNotFound
	}

	return nil
}

func (r *ProductRepository) Delete(id uint) error {
	result := r.db.Delete(&domain.Product{}, id)
	if result.Error != nil {
		return fmt.Errorf("%w: %v", ErrDatabaseOperation, result.Error)
	}

	if result.RowsAffected == 0 {
		return ErrProductNotFound
	}

	return nil
}

func (r *ProductRepository) ExistsByName(name string, excludeID uint) (bool, error) {
	var count int64

	query := r.db.Model(&domain.Product{}).Where("name = ?", name)
	if excludeID > 0 {
		query = query.Where("id != ?", excludeID)
	}

	if err := query.Count(&count).Error; err != nil {
		return false, fmt.Errorf("%w: %v", ErrDatabaseOperation, err)
	}

	return count > 0, nil
}