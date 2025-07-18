package usecase

import (
	"errors"

	dtos "github.com/savanyv/e-commerce/internal/dto"
	"github.com/savanyv/e-commerce/internal/models"
	"github.com/savanyv/e-commerce/internal/repository"
)

type ProductUsecase interface {
	CreateProduct(req *dtos.CreateProductRequest) error
	GetAllProduct(page, limit int) ([]*dtos.ProductResponse, int64, error)
	GetByIDProduct(ID int) (*dtos.ProductResponse, error)
	UpdateProduct(ID int, req *dtos.UpdateProductRequest) error
	DeleteProduct(ID int) error
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

	if req.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	if req.Price <= 0 {
		return errors.New("price must be greater than 0")
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

func (u *productUsecase) GetAllProduct(page, limit int) ([]*dtos.ProductResponse,int64, error) {
	offset := (page - 1) * limit
	products, err := u.repo.FindAll(limit, offset)
	if err != nil {
		return nil, 0, errors.New("error getting all products")
	}

	total, err := u.repo.Count()
	if err != nil {
		return nil,0, errors.New("error getting total products")
	}

	var res []*dtos.ProductResponse
	for _, product := range products {
		res = append(res, &dtos.ProductResponse{
			ID:        product.ID,
			Name:      product.Name,
			Price:     product.Price,
			Quantity:  product.Quantity,
			Brand: dtos.BrandSimple{
				ID:   product.IDBrand,
				Name: product.Brand.Name,
			},
		})
	}

	return res, total, nil
}

func (u *productUsecase) GetByIDProduct(ID int) (*dtos.ProductResponse, error) {
	product, err := u.repo.FindByID(ID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	res := &dtos.ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		Quantity:  product.Quantity,
		Brand: dtos.BrandSimple{
			ID:   product.IDBrand,
			Name: product.Brand.Name,
		},
	}

	return res, nil
}

func (u *productUsecase) UpdateProduct(ID int, req *dtos.UpdateProductRequest) error {
	product, err := u.repo.FindByID(ID)
	if err != nil {
		return errors.New("product not found")
	}

	if req.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	if req.Price <= 0 {
		return errors.New("price must be greater than 0")
	}

	product.Name = req.Name
	product.Price = req.Price
	product.Quantity = req.Quantity
	product.IDBrand = req.IDBrand

	if err := u.repo.UpdateProduct(product); err != nil {
		return errors.New("error updating product")
	}

	return nil
}

func (u *productUsecase) DeleteProduct(ID int) error {
	_, err := u.repo.FindByID(ID)
	if err != nil {
		return errors.New("product not found")
	}

	if err := u.repo.DeleteProduct(ID); err != nil {
		return errors.New("error deleting product")
	}

	return nil
}
