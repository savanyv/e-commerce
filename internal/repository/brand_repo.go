package repository

import (
	"github.com/savanyv/e-commerce/internal/models"
	"gorm.io/gorm"
)

type BrandRepository interface {
	Create(brand *models.Brand) error
	Delete(id int) error
	FindByID(id int) (*models.Brand, error)
	FindAll() ([]*models.Brand, error)
	IsUsedByProduct(id int) (bool, error)
}

type brandRepository struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) BrandRepository {
	return &brandRepository{
		db: db,
	}
}

func (r *brandRepository) Create(brand *models.Brand) error {
	if err := r.db.Create(brand).Error; err != nil {
		return err
	}

	return nil
}

func (r *brandRepository) Delete(id int) error {
	if err := r.db.Where("id = ?", id).Delete(&models.Brand{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *brandRepository) FindByID(id int) (*models.Brand, error) {
	var brand models.Brand
	if err := r.db.Where("id = ?", id).First(&brand).Error; err != nil {
		return nil, err
	}

	return &brand, nil
}

func (r *brandRepository) FindAll() ([]*models.Brand, error) {
	var brands []*models.Brand
	if err := r.db.Find(&brands).Error; err != nil {
		return nil, err
	}

	return brands, nil
}

func (r *brandRepository) IsUsedByProduct(id int) (bool, error) {
	var count int64
	if err := r.db.Model(&models.Product{}).Where("id_brand = ?", id).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
