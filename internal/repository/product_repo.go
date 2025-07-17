package repository

import (
	"github.com/savanyv/e-commerce/internal/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(ID int) error
	FindByID(ID int) (*models.Product, error)
	FindAll(limit, offset int) ([]*models.Product, error)
	Count() (int64, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) CreateProduct(product *models.Product) error {
	if err := r.db.Create(product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) UpdateProduct(product *models.Product) error {
	if err := r.db.Save(product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) DeleteProduct(ID int) error {
	if err := r.db.Where("id = ?", ID).Delete(&models.Product{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) FindByID(ID int) (*models.Product, error) {
	var product models.Product
	if err := r.db.Preload("Brand").Where("id = ?", ID).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) FindAll(limit, offset int) ([]*models.Product, error) {
	var products []*models.Product
	if err := r.db.Preload("Brand").Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) Count() (int64, error) {
	var total int64
	err := r.db.Model(&models.Product{}).Count(&total).Error
	if err != nil {
		return 0, err
	}

	return total, nil
}
