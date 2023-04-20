package products

import (
	"app.eirc/internal/interactor/models/special"
)

// Table struct is products database table struct
type Table struct {
	// 產品ID
	ProductID string `gorm:"<-:create;column:product_id;type:uuid;not null;primaryKey;" json:"product_id"`
	// 產品名稱
	Name string `gorm:"column:name;type:text;not null;" json:"name"`
	// 產品編號
	Code string `gorm:"column:code;type:text;" json:"code"`
	// 是否啟用
	IsEnable bool `gorm:"column:is_enable;type:bool;not null;" json:"is_enable"`
	// 產品描述
	Description string `gorm:"column:description;type:text;" json:"description"`
	special.UseTable
}

// Base struct is corresponding to products table structure file
type Base struct {
	// 產品ID
	ProductID *string `json:"product_id,omitempty"`
	// 產品名稱
	Name *string `json:"name,omitempty"`
	// 產品編號
	Code *string `json:"code,omitempty"`
	// 是否啟用
	IsEnable *bool `json:"is_enable,omitempty"`
	// 產品描述
	Description *string `json:"description,omitempty"`
	special.UseBase
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "products"
}
