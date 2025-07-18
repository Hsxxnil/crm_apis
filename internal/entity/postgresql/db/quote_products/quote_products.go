package quote_products

import (
	"crm/internal/entity/postgresql/db/products"
	"crm/internal/entity/postgresql/db/users"
	"crm/internal/interactor/models/special"
)

// Table struct is quote_products database table struct
type Table struct {
	// 報價產品ID
	QuoteProductID string `gorm:"<-:create;column:quote_product_id;type:uuid;not null;primaryKey;" json:"quote_product_id"`
	// 報價ID
	QuoteID string `gorm:"column:quote_id;type:text;not null;" json:"quote_id"`
	// 產品ID
	ProductID string `gorm:"column:product_id;type:text;not null;" json:"product_id"`
	// products data
	Products products.Table `gorm:"foreignKey:ProductID;references:ProductID" json:"products,omitempty"`
	// 報價產品數量
	Quantity int `gorm:"column:quantity;type:int;not null;" json:"quantity"`
	// 報價產品單價
	UnitPrice float64 `gorm:"column:unit_price;type:numeric;not null;" json:"unit_price"`
	// 報價產品小計
	SubTotal float64 `gorm:"column:sub_total;type:numeric;not null;" json:"sub_total"`
	// 報價產品總價
	TotalPrice float64 `gorm:"column:total_price;type:numeric;not null;" json:"total_price"`
	// 報價產品折扣
	Discount float64 `gorm:"column:discount;type:text;not null;" json:"discount"`
	// 報價產品描述
	Description string `gorm:"column:description;type:text;" json:"description"`
	// 報價產品號碼
	Code string `gorm:"column:code;type:text;not null;" json:"code"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:CreatedBy;references:UserID" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:UpdatedBy;references:UserID" json:"updated_by_users,omitempty"`
	special.Table
}

// Base struct is corresponding to quote_products table structure file
type Base struct {
	// 報價產品ID
	QuoteProductID *string `json:"quote_product_id,omitempty"`
	// 報價ID
	QuoteID *string `json:"quote_id,omitempty"`
	// 產品ID
	ProductID *string `json:"product_id,omitempty"`
	// products data
	Products products.Base `json:"products,omitempty"`
	// 報價產品數量
	Quantity *int `json:"quantity,omitempty"`
	// 報價產品單價
	UnitPrice *float64 `json:"unit_price,omitempty"`
	// 報價產品小計
	SubTotal *float64 `json:"sub_total,omitempty"`
	// 報價產品總價
	TotalPrice *float64 `json:"total_price,omitempty"`
	// 報價產品折扣
	Discount *float64 `json:"discount,omitempty"`
	// 報價產品描述
	Description *string `json:"description,omitempty"`
	// 報價產品號碼
	Code *string `json:"code,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.Base
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "quote_products"
}
