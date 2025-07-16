package usecase

import (
	"errors"

	dtos "github.com/savanyv/e-commerce/internal/dto"
	"github.com/savanyv/e-commerce/internal/models"
	"github.com/savanyv/e-commerce/internal/repository"
)

type BrandUsecase interface {
	Create(req *dtos.CreateBrandRequest) error
	Delete(ID int) error
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

func (u *brandUsecase) Delete(ID int) error {
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
