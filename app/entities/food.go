package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Food struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt time.Time `json:"deleted_at" gorm:"index"`
	Products  []Product `json:"products" gorm:"json"`
	TSV       string    `json:"-" gorm:"type:tsvector"`
}

type CreateFood struct {
	Products []Product `json:"products" gorm:"json"`
}

type Product struct {
	ID         uuid.UUID       `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt  time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  time.Time       `json:"deleted_at" gorm:"index"`
	Name       string          `json:"name" validate:"required"`
	CategoryID uuid.UUID       `json:"category_id" validate:"required,uuid4"`
	Category   ProductCategory `json:"category"`
	Protein    int             `json:"protein" validate:"required"`
	Fat        int             `json:"fat" validate:"required"`
	Carbs      int             `json:"carbs" validate:"required"`
	Calories   int             `json:"calories" validate:"required"`
	TSV        string          `json:"-" gorm:"type:tsvector"`
}

type CreateProduct struct {
	Name     string    `json:"name" validate:"required"`
	Category uuid.UUID `json:"category" validate:"required,uuid4"`
	Protein  int       `json:"protein" validate:"required"`
	Fat      int       `json:"fat" validate:"required"`
	Carbs    int       `json:"carbs" validate:"required"`
	Calories int       `json:"calories" validate:"required"`
}

type ProductCategory struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt time.Time `json:"deleted_at" gorm:"index"`
	Name      string    `json:"name" validate:"required"`
	TSV       string    `json:"-" gorm:"type:tsvector"`
	Products  []Product `json:"products" gorm:"foreignKey:CategoryID"`
}

func (p *Product) ProductDelete(tx *gorm.DB) error {
	p.DeletedAt = time.Now()
	return tx.Save(p).Error
}

func (p *ProductCategory) ProductCategoryDelete(tx *gorm.DB) error {
	p.DeletedAt = time.Now()
	return tx.Save(p).Error
}

func (p *Food) FoodDelete(tx *gorm.DB) error {
	p.DeletedAt = time.Now()
	return tx.Save(p).Error
}
