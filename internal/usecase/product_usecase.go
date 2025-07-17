package usecase

import (
	"errors"

	dtos "github.com/savanyv/e-commerce/internal/dto"
	"github.com/savanyv/e-commerce/internal/models"
	"github.com/savanyv/e-commerce/internal/repository"
)

type ProductUsecase interface {
	CreateProduct(req *dtos.CreateProductRequest) error
	// UpdateProduct(req *dtos.UpdateProductRequest) error
	// DeleteProduct(ID int) error
	// GetByIDProduct(ID int) (*dtos.ProductResponse, error)
	// GetAllProduct(page, limit int) ([]*dtos.ProductResponse, error)
}

type productUsecase struct {
	repo repository.ProductRepository
	brandRepo repository.BrandRepository
}

func NewProductRepository(repo repository.ProductRepository, brandRepo repository.BrandRepository) ProductUsecase {
	return &productUsecase{
		repo: repo,
		brandRepo: brandRepo,
	}
}

func (u *productUsecase) CreateProduct(req *dtos.CreateProductRequest) error {
	_, err := u.brandRepo.FindByID(int(req.IDBrand))
	if err != nil {
		return errors.New("brand not found")
	}

	product := &models.Product{
		Name:     req.Name,
		Price:    req.Price,
		Quantity: req.Quantity,
		IDBrand:  req.IDBrand,
	}

	if err := u.repo.CreateProduct(product); err != nil {
		return errors.New("error creating product")
	}

	return nil
}
