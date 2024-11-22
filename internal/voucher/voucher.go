package voucher

import (
	"golang-api-service/internal/brand"
	"time"
)



type Voucher struct {
	ID          int         `json:"id" gorm:"primaryKey"`
	BrandID     int         `json:"brand_id" gorm:"not null"`
	Brand       brand.Brand `json:"brand" gorm:"foreignKey:BrandID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Code        string      `json:"code" gorm:"unique;not null"`
	CostInPoints int        `json:"cost_in_points" gorm:"not null"`
	CreatedAt   time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}
