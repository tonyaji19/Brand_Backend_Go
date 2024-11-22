package voucher

import (
	"gorm.io/gorm"
)

// VoucherRepository interface
type VoucherRepository interface {
	CreateVoucher(voucher *Voucher) error
	GetVoucherByID(voucherID int) (*Voucher, error)
	GetAllVouchers() ([]Voucher, error)
	GetVouchersByBrandID(brandID int) ([]Voucher, error)
}

// voucherRepository struct
type voucherRepository struct {
	db *gorm.DB
}

// NewVoucherRepository initializes a new voucher repository
func NewVoucherRepository(db *gorm.DB) VoucherRepository {
	return &voucherRepository{db: db} // Sesuai dengan tipe *voucherRepository
}

// CreateVoucher creates a new voucher
func (r *voucherRepository) CreateVoucher(voucher *Voucher) error {
	if err := r.db.Create(voucher).Error; err != nil {
		return err
	}
	return nil
}

// GetVoucherByID retrieves a voucher by its ID
func (r *voucherRepository) GetVoucherByID(id int) (*Voucher, error) { // Tipe receiver diperbaiki
	var voucher Voucher
	err := r.db.Preload("Brand").First(&voucher, id).Error
	if err != nil {
		return nil, err
	}
	return &voucher, nil
}

// GetAllVouchers retrieves all vouchers
func (r *voucherRepository) GetAllVouchers() ([]Voucher, error) { // Tipe receiver diperbaiki
	var vouchers []Voucher
	err := r.db.Preload("Brand").Find(&vouchers).Error
	return vouchers, err
}

// Implementasi di struct voucherRepository
func (r *voucherRepository) GetVouchersByBrandID(brandID int) ([]Voucher, error) {
    var vouchers []Voucher
    err := r.db.Where("brand_id = ?", brandID).Preload("Brand").Find(&vouchers).Error
    return vouchers, err
}