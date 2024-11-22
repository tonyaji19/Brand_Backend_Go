package brand

import (
	"gorm.io/gorm"
)

type BrandRepository interface {
	CreateBrand(brand *Brand) error
	GetBrandByID(brandID int) (*Brand, error)
	GetAllBrands() ([]Brand, error)
}

type brandRepository struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) BrandRepository {
	return &brandRepository{db: db}
}

// Create a new brand
func (r *brandRepository) CreateBrand(brand *Brand) error {
	if err := r.db.Create(brand).Error; err != nil {
		return err
	}
	return nil
}

// Get brand by ID
func (r *brandRepository) GetBrandByID(brandID int) (*Brand, error) {
	var brand Brand
	err := r.db.First(&brand, brandID).Error
	if err != nil {
		return nil, err
	}
	return &brand, nil
}

// Get all brands
func (r *brandRepository) GetAllBrands() ([]Brand, error) {
	var brands []Brand
	err := r.db.Find(&brands).Error
	if err != nil {
		return nil, err
	}
	return brands, nil
}
