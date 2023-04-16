package lead_contacts

import (
	"app.eirc/internal/interactor/models/page"
	"app.eirc/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 線索聯絡人名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	// 線索聯絡人職稱
	Title string `json:"title,omitempty"`
	// 線索聯絡人電話
	PhoneNumber string `json:"phone_number,omitempty" binding:"required" validate:"required"`
	// 線索聯絡人行動電話
	CellPhone string `json:"cell_phone,omitempty"`
	// 線索聯絡人電子郵件
	Email string `json:"email,omitempty"`
	// 線索聯絡人LINE
	Line string `json:"line,omitempty"`
	// 線索ID
	LeadID string `json:"lead_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}

// Field is structure file for search
type Field struct {
	// 線索聯絡人ID
	LeadContactID string `json:"lead_contact_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 線索聯絡人名稱
	Name *string `json:"name,omitempty" from:"name"`
	// 線索聯絡人職稱
	Title *string `json:"title,omitempty" from:"title"`
	// 線索聯絡人電話
	PhoneNumber *string `json:"phone_number,omitempty" from:"phone_number"`
	// 線索聯絡人行動電話
	CellPhone *string `json:"cell_phone,omitempty" from:"cell_phone"`
	// 商機線索聯絡人電子郵件
	Email *string `json:"email,omitempty" from:"email"`
	// 商機線索聯絡人LINE
	Line *string `json:"line,omitempty" from:"line"`
	// 線索ID
	LeadID *string `json:"lead_id,omitempty" from:"lead_id"`
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
	LeadContacts []*struct {
		// 商機線索聯絡人ID
		LeadContactID string `json:"lead_contact_id,omitempty"`
		// 商機線索聯絡人名稱
		Name string `json:"name,omitempty"`
		// 商機線索聯絡人職稱
		Title string `json:"title,omitempty"`
		// 商機線索聯絡人電話
		PhoneNumber string `json:"phone_number,omitempty"`
		// 商機線索聯絡人行動電話
		CellPhone string `json:"cell_phone,omitempty"`
		// 商機線索聯絡人電子郵件
		Email string `json:"email,omitempty"`
		// 商機線索聯絡人LINE
		Line string `json:"line,omitempty"`
		// 線索ID
		LeadID string `json:"lead_id,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"lead_contacts"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 商機線索聯絡人ID
	LeadContactID string `json:"lead_contact_id,omitempty"`
	// 商機線索聯絡人名稱
	Name string `json:"name,omitempty"`
	// 商機線索聯絡人職稱
	Title string `json:"title,omitempty"`
	// 商機線索聯絡人電話
	PhoneNumber string `json:"phone_number,omitempty"`
	// 商機線索聯絡人行動電話
	CellPhone string `json:"cell_phone,omitempty"`
	// 商機線索聯絡人電子郵件
	Email string `json:"email,omitempty"`
	// 商機線索聯絡人LINE
	Line string `json:"line,omitempty"`
	// 線索ID
	LeadID string `json:"lead_id,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
}

// Update struct is used to update achieves
type Update struct {
	// 商機線索聯絡人ID
	LeadContactID string `json:"lead_contact_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 商機線索聯絡人名稱
	Name *string `json:"name,omitempty"`
	// 商機線索聯絡人職稱
	Title *string `json:"title,omitempty"`
	// 商機線索聯絡人電話
	PhoneNumber *string `json:"phone_number,omitempty"`
	// 商機線索聯絡人行動電話
	CellPhone *string `json:"cell_phone,omitempty"`
	// 商機線索聯絡人電子郵件
	Email *string `json:"email,omitempty"`
	// 商機線索聯絡人LINE
	Line *string `json:"line,omitempty"`
	// 線索ID
	LeadID *string `json:"lead_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}
