package usecase

import (
	"errors"
	"strings"

	dtos "github.com/savanyv/e-commerce/internal/dto"
	"github.com/savanyv/e-commerce/internal/models"
	"github.com/savanyv/e-commerce/internal/repository"
)

type BrandUsecase interface {
	CreateBrand(req *dtos.CreateBrandRequest) error
	DeleteBrand(ID int) error
	GetAllBrands() ([]*dtos.BrandResponse, error)
}

type brandUsecase struct {
	repo repository.BrandRepository
}

func NewBrandUsecase(repo repository.BrandRepository) BrandUsecase {
	return &brandUsecase{
		repo: repo,
	}
}

func (u *brandUsecase) CreateBrand(req *dtos.CreateBrandRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name is required")
	}

	existing, _ := u.repo.FindByName(req.Name)
	if existing != nil {
		return errors.New("brand with this name already exists")
	}

	brand := &models.Brand{
		Name: req.Name,
	}

	if err := u.repo.Create(brand); err != nil {
		return errors.New("error creating brand")
	}

	return nil
}

func (u *brandUsecase) DeleteBrand(ID int) error {
	_, err := u.repo.FindByID(ID)
	if err != nil {
		return errors.New("brand not found")
	}

	used, err := u.repo.IsUsedByProduct(ID)
	if err != nil {
		return errors.New("error checking if brand is used by product")
	}

	if used {
		return errors.New("brand is used by product")
	}

	if err := u.repo.Delete(ID); err != nil {
		return errors.New("error deleting brand")
	}

	return nil
}

func (u *brandUsecase) GetAllBrands() ([]*dtos.BrandResponse, error) {
	brands, err := u.repo.FindAll()
	if err != nil {
		return nil, errors.New("error getting all brands")
	}

	var res []*dtos.BrandResponse
	for _, brand := range brands {
		res = append(res, &dtos.BrandResponse{
			ID:        brand.ID,
			Name:      brand.Name,
		})
	}

	return res, nil
}
