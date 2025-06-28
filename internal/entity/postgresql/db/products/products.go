package products

import (
	"crm/internal/entity/postgresql/db/users"
	model "crm/internal/interactor/models/products"
	"crm/internal/interactor/models/sort"
	"crm/internal/interactor/models/special"
)

// Table struct is products database table struct
type Table struct {
	// 產品ID
	ProductID string `gorm:"<-:create;column:product_id;type:uuid;not null;primaryKey;" json:"product_id"`
	// 產品名稱
	Name string `gorm:"column:name;type:text;not null;" json:"name"`
	// 產品識別碼
	Code string `gorm:"column:code;type:text;" json:"code"`
	// 產品是否啟用
	IsEnable bool `gorm:"column:is_enable;type:bool;not null;" json:"is_enable"`
	// 產品描述
	Description string `gorm:"column:description;type:text;" json:"description"`
	// 產品價格
	Price float64 `gorm:"column:price;type:numeric;not null;" json:"price"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:CreatedBy;references:UserID" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:UpdatedBy;references:UserID" json:"updated_by_users,omitempty"`
	special.Table
}

// Base struct is corresponding to products table structure file
type Base struct {
	// 產品ID
	ProductID *string `json:"product_id,omitempty"`
	// 產品名稱
	Name *string `json:"name,omitempty"`
	// 產品識別碼
	Code *string `json:"code,omitempty"`
	// 產品是否啟用
	IsEnable *bool `json:"is_enable,omitempty"`
	// 產品描述
	Description *string `json:"description,omitempty"`
	// 產品價格
	Price *float64 `json:"price,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.Base
	// 搜尋欄位
	model.Filter `json:"filter"`
	// 排序欄位
	sort.Sort `json:"sort"`
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "products"
}
