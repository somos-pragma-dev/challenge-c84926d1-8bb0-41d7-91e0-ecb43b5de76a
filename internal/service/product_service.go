package service

import (
	"errors"
	repository "github.com/empresa/product-api/internal/repository"
	dto "github.com/empresa/product-api/internal/dto"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetAllProducts() ([]dto.ProductDTO, error) {
	products, err := s.repo.FindAll()
	if err != nil {
		return nil, errors.New("error al consultar la base de datos")
	}

	var dtos []dto.ProductDTO
	for _, p := range products {
		dtos = append(dtos, dto.ProductDTO{
			ID:       p.ID,
			Name:     p.Name,
			Price:    p.Price,
			Stock:    p.Stock,
			Category: p.Category,
		})
	}

	return dtos, nil
}

func (s *ProductService) GetProductByID(id uint) (dto.ProductDTO, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return dto.ProductDTO{}, errors.New("producto no encontrado")
	}

	return dto.ProductDTO{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Stock:    product.Stock,
		Category: product.Category,
	}, nil
}

func (s *ProductService) CreateProduct(req dto.CreateProductRequest) (dto.ProductDTO, error) {
	if req.Name == "" {
		return dto.ProductDTO{}, errors.New("el nombre no puede estar vacío")
	}

	if req.Price < 0 {
		return dto.ProductDTO{}, errors.New("el precio no puede ser negativo")
	}

	exists, err := s.repo.ExistsByName(req.Name, 0)
	if err != nil {
		return dto.ProductDTO{}, errors.New("error al verificar el nombre del producto")
	}
	if exists {
		return dto.ProductDTO{}, errors.New("el nombre del producto ya existe")
	}

	product := &domain.Product{
		Name:     req.Name,
		Price:    req.Price,
		Stock:    req.Stock,
		Category: req.Category,
	}

	if err := s.repo.Create(product); err != nil {
		return dto.ProductDTO{}, errors.New("error al guardar el producto")
	}

	return dto.ProductDTO{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Stock:    product.Stock,
		Category: product.Category,
	}, nil
}

func (s *ProductService) UpdateProduct(id uint, req dto.UpdateProductRequest) (dto.ProductDTO, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return dto.ProductDTO{}, errors.New("producto no encontrado")
	}

	if req.Name != nil && *req.Name != "" {
		exists, err := s.repo.ExistsByName(*req.Name, id)
		if err != nil {
			return dto.ProductDTO{}, errors.New("error al verificar el nombre del producto")
		}
		if exists {
			return dto.ProductDTO{}, errors.New("el nombre del producto ya existe")
		}
		product.Name = *req.Name
	}

	if req.Price != nil {
		if *req.Price < 0 {
			return dto.ProductDTO{}, errors.New("el precio no puede ser negativo")
		}
		product.Price = *req.Price
	}

	if req.Stock != nil {
		product.Stock = *req.Stock
	}

	if req.Category != nil {
		product.Category = *req.Category
	}

	if err := s.repo.Update(product); err != nil {
		return dto.ProductDTO{}, errors.New("error al actualizar el producto")
	}

	return dto.ProductDTO{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Stock:    product.Stock,
		Category: product.Category,
	}, nil
}

func (s *ProductService) DeleteProduct(id uint) error {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("producto no encontrado")
	}

	if err := s.repo.Delete(product.ID); err != nil {
		return errors.New("error al eliminar el producto")
	}

	return nil
}