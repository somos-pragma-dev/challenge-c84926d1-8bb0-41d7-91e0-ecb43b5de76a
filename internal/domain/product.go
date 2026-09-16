package domain

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `gorm:"size:255;not null;uniqueIndex" json:"name"`
	Price     float64        `gorm:"not null" json:"price"`
	Stock     int            `gorm:"not null;default:0" json:"stock"`
	Category  string         `gorm:"size:100;not null" json:"category"`
}

func (p *Product) TableName() string {
	return "products"
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.New("el nombre del producto es requerido")
	}

	if len(p.Name) < 3 {
		return errors.New("el nombre del producto debe tener al menos 3 caracteres")
	}

	if len(p.Name) > 255 {
		return errors.New("el nombre del producto no puede exceder 255 caracteres")
	}

	if p.Price < 0 {
		return errors.New("el precio del producto no puede ser negativo")
	}

	if p.Stock < 0 {
		return errors.New("el stock del producto no puede ser negativo")
	}

	if p.Category == "" {
		return errors.New("la categoría del producto es requerida")
	}

	return nil
}

func (p *Product) UpdateStock(quantity int) error {
	newStock := p.Stock + quantity
	if newStock < 0 {
		return errors.New("el stock resultante no puede ser negativo")
	}
	p.Stock = newStock
	return nil
}

func (p *Product) IsAvailable() bool {
	return p.Stock > 0
}

func (p *Product) CanDecreaseStock(quantity int) bool {
	return p.Stock >= quantity
}

type ProductRepository interface {
	Create(product *Product) error
	FindByID(id uint) (*Product, error)
	FindAll() ([]Product, error)
	FindByName(name string) (*Product, error)
	FindByCategory(category string) ([]Product, error)
	Update(product *Product) error
	Delete(id uint) error
	ExistsByName(name string, excludeID uint) (bool, error)
}