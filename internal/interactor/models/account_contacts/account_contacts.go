package account_contacts

import (
	"app.eirc/internal/interactor/models/page"
	"app.eirc/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 帳戶ID
	AccountID string `json:"account_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 聯絡人ID
	ContactID string `json:"contact_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 帳戶聯絡人ID
	AccountContactID string `json:"account_contact_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty" form:"account_id"`
	// 聯絡人ID
	ContactID *string `json:"contact_id,omitempty" form:"contact_id"`
	// 帳戶聯絡人是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty" form:"is_deleted"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	// 搜尋結構檔
	Field
	// 分頁搜尋結構檔
	page.Pagination
}

// List is multiple return structure files
type List struct {
	// 多筆
	AccountContacts []*struct {
		// 帳戶聯絡人ID
		AccountContactID string `json:"account_contact_id,omitempty"`
		// 帳戶ID
		AccountID string `json:"account_id,omitempty"`
		// 聯絡人ID
		ContactID string `json:"contact_id,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"account_contacts"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 帳戶聯絡人ID
	AccountContactID string `json:"account_contact_id,omitempty"`
	// 帳戶ID
	AccountID string `json:"account_id,omitempty"`
	// 聯絡人ID
	ContactID string `json:"contact_id,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
}

// Update struct is used to update achieves
type Update struct {
	// 帳戶聯絡人ID
	AccountContactID string `json:"account_contact_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 聯絡人ID
	ContactID *string `json:"contact_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 帳戶聯絡人是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty" swaggerignore:"true"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// AccountSingle return structure file for opportunities
type AccountSingle struct {
	// 聯絡人ID
	ContactID string `json:"contact_id,omitempty"`
	// 聯絡人名稱
	ContactName string `json:"contact_name,omitempty"`
	// 聯絡人職稱
	ContactTitle string `json:"contact_title,omitempty"`
	// 聯絡人電話
	ContactPhoneNumber string `json:"contact_phone_number,omitempty"`
	// 聯絡人行動電話
	ContactCellPhone string `json:"contact_cell_phone,omitempty"`
	// 聯絡人電子郵件
	ContactEmail string `json:"contact_email,omitempty"`
	// 聯絡人稱謂
	ContactSalutation string `json:"contact_salutation,omitempty"`
	// 聯絡人部門
	ContactDepartment string `json:"contact_department,omitempty"`
}
