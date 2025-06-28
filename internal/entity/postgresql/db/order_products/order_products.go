package order_products

import (
	"crm/internal/entity/postgresql/db/products"
	"crm/internal/entity/postgresql/db/users"
	"crm/internal/interactor/models/special"
)

// Table struct is order_products database table struct
type Table struct {
	// 訂單產品ID
	OrderProductID string `gorm:"<-:create;column:order_product_id;type:uuid;not null;primaryKey;" json:"order_product_id"`
	// 訂單ID
	OrderID string `gorm:"column:order_id;type:text;not null;" json:"order_id"`
	// 產品ID
	ProductID string `gorm:"column:product_id;type:text;not null;" json:"product_id"`
	// products data
	Products products.Table `gorm:"foreignKey:ProductID;references:ProductID" json:"products,omitempty"`
	// 訂單產品數量
	Quantity int `gorm:"column:quantity;type:int;not null;" json:"quantity"`
	// 訂單產品單價
	UnitPrice float64 `gorm:"column:unit_price;type:numeric;not null;" json:"unit_price"`
	// 訂單產品報價
	QuotePrice float64 `gorm:"column:quote_price;type:numeric;not null;" json:"quote_price"`
	// 訂單產品小計
	SubTotal float64 `gorm:"column:sub_total;type:numeric;not null;" json:"sub_total"`
	// 訂單產品描述
	Description string `gorm:"column:description;type:text;" json:"description"`
	// 訂單產品號碼
	Code string `gorm:"column:code;type:text;not null;" json:"code"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:CreatedBy;references:UserID" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:UpdatedBy;references:UserID" json:"updated_by_users,omitempty"`
	special.Table
}

// Base struct is corresponding to order_products table structure file
type Base struct {
	// 訂單產品ID
	OrderProductID *string `json:"order_product_id,omitempty"`
	// 訂單ID
	OrderID *string `json:"order_id,omitempty"`
	// 產品ID
	ProductID *string `json:"product_id,omitempty"`
	// products data
	Products products.Base `json:"products,omitempty"`
	// 訂單產品數量
	Quantity *int `json:"quantity,omitempty"`
	// 訂單產品單價
	UnitPrice *float64 `json:"unit_price,omitempty"`
	// 訂單產品報價
	QuotePrice *float64 `json:"quote_price,omitempty"`
	// 訂單產品小計
	SubTotal *float64 `json:"sub_total,omitempty"`
	// 訂單產品描述
	Description *string `json:"description,omitempty"`
	// 訂單產品號碼
	Code *string `json:"code,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.Base
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "order_products"
}
