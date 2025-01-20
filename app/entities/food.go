package entities

import (
	"time"

	"gorm.io/gorm"
)

type Food struct {
	ID         int64               `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time           `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time           `json:"-" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt      `json:"-" gorm:"index;default:NULL"`
	ImageUrl   *string             `json:"image" validate:"http_url"`
	Name       string              `json:"name" validate:"required" gorm:"index"`
	Products   []ProductCreateFood `json:"products" gorm:"json"`
	Protein    float64             `json:"protein" validate:"required"`
	Fat        float64             `json:"fat" validate:"required"`
	Carbs      float64             `json:"carbs" validate:"required"`
	Calories   float64             `json:"calories" validate:"required"`
	UserID     *int64              `json:"user_id" gorm:"index"`
	User       *User               `json:"-" gorm:"foreignKey:UserID"`
	IsPublic   bool                `json:"is_public" gorm:"index"`
	Recepit    *string             `json:"receipt" gorm:"default:NULL;type:text"`
	ReceiptTSV string              `json:"-" gorm:"type:tsvector;index"`
	TSV        string              `json:"-" gorm:"type:tsvector;index"`
}

type CreateFood struct {
	Products *[]ProductCreateFood `json:"products"`
	Name     string               `json:"name" validate:"required"`
	Receipt  *string              `json:"receipt"`
	IsPublic bool                 `json:"is_public"`
	Info     *ProductStat         `json:"info"` // always 100g
}

type ProductStat struct {
	Protein  float64 `json:"protein"`
	Fat      float64 `json:"fat"`
	Carbs    float64 `json:"carbs"`
	Calories float64 `json:"calories"`
}

type ProductCreateFood struct {
	Weight    float64      `json:"weight" validate:"required"`
	ProductID *int64       `json:"product_id"`
	Name      *string      `json:"name"`
	Info      *ProductStat `json:"info" validate:"json"` // ALWAYS 100gr
}

type Product struct {
	ID         int64           `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time       `json:"-" gorm:"autoCreateTime"`
	UpdatedAt  time.Time       `json:"-" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt  `json:"-" gorm:"index,default:NULL"`
	Name       string          `json:"name" validate:"required"`
	CategoryID int64           `json:"category_id" validate:"required"`
	Category   ProductCategory `json:"category" gorm:"foreignKey:CategoryID"`
	ImageUrl   string          `json:"image" validate:"http_url"`
	Protein    float64         `json:"protein" validate:"required"`
	Fat        float64         `json:"fat" validate:"required"`
	Carbs      float64         `json:"carbs" validate:"required"`
	Calories   float64         `json:"calories" validate:"required"`
	TSV        string          `json:"-" gorm:"type:tsvector;index"`
}

type CreateProduct struct {
	Name       string  `json:"name" validate:"required"`
	CategoryID int64   `json:"category" validate:"required"`
	Protein    float64 `json:"protein" validate:"required"`
	Fat        float64 `json:"fat" validate:"required"`
	Carbs      float64 `json:"carbs" validate:"required"`
	Calories   float64 `json:"calories" validate:"required"`
	ImageUrl   string  `json:"image" validate:"http_url"`
}

type ProductCategory struct {
	ID        int64          `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index,default:NULL"`
	Name      string         `json:"name" validate:"required" gorm:"unique"`
	TSV       string         `json:"-" gorm:"type:tsvector;index"`
	Products  []Product      `json:"products" gorm:"foreignKey:CategoryID"`
	ImageUrl  string         `json:"image" validate:"http_url"`
}

type CreateProductCategory struct {
	Name string `json:"name" validate:"required"`
}
