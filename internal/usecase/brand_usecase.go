package usecase

import (
	"errors"

	dtos "github.com/savanyv/e-commerce/internal/dto"
	"github.com/savanyv/e-commerce/internal/models"
	"github.com/savanyv/e-commerce/internal/repository"
)

type BrandUsecase interface {
	Create(req *dtos.CreateBrandRequest) error
}

type brandUsecase struct {
	repo repository.BrandRepository
}

func NewBrandUsecase(repo repository.BrandRepository) BrandUsecase {
	return &brandUsecase{
		repo: repo,
	}
}

func (u *brandUsecase) Create(req *dtos.CreateBrandRequest) error {
	brand := &models.Brand{
		Name: req.Name,
	}

	if err := u.repo.Create(brand); err != nil {
		return errors.New("error creating brand")
	}

	return nil
}
