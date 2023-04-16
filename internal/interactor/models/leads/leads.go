package leads

import (
	"app.eirc/internal/interactor/models/page"
	"app.eirc/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 線索狀態
	Status string `json:"status,omitempty" binding:"required" validate:"required"`
	// 線索客戶名稱
	CompanyName string `json:"company_name,omitempty" binding:"required" validate:"required"`
	// 線索來源ID
	SourceID string `json:"source_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 線索客戶行業ID
	IndustryID string `json:"industry_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 線索分級
	Rating string `json:"rating,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}

// Field is structure file for search
type Field struct {
	// 線索ID
	LeadID string `json:"lead_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 線索狀態
	Status *string `json:"status,omitempty" from:"status"`
	// 線索客戶名稱
	CompanyName *string `json:"company_name,omitempty" from:"company_name"`
	// 線索來源ID
	SourceID *string `json:"source_id,omitempty" from:"source_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 線索客戶行業ID
	IndustryID *string `json:"industry_id,omitempty" from:"industry_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 線索分級
	Rating *string `json:"rating,omitempty" from:"rating"`
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
	Leads []*struct {
		// 線索ID
		LeadID string `json:"lead_id,omitempty"`
		// 線索狀態
		Status string `json:"status,omitempty"`
		// 線索客戶名稱
		CompanyName string `json:"company_name,omitempty"`
		// 線索來源ID
		SourceID string `json:"source_id,omitempty"`
		// 線索客戶行業ID
		IndustryID string `json:"industry_id,omitempty"`
		// 線索分級
		Rating string `json:"rating,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"leads"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 線索ID
	LeadID string `json:"lead_id,omitempty"`
	// 線索狀態
	Status string `json:"status,omitempty"`
	// 線索客戶名稱
	CompanyName string `json:"company_name,omitempty"`
	// 線索來源ID
	SourceID string `json:"source_id,omitempty"`
	// 線索客戶行業ID
	IndustryID string `json:"industry_id,omitempty"`
	// 線索分級
	Rating string `json:"rating,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
}

// Update struct is used to update achieves
type Update struct {
	// 線索ID
	LeadID string `json:"lead_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 線索狀態
	Status *string `json:"status,omitempty"`
	// 線索客戶名稱
	CompanyName *string `json:"company_name,omitempty"`
	// 線索來源ID
	SourceID *string `json:"source_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 線索客戶行業ID
	IndustryID string `json:"industry_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 線索分級
	Rating *string `json:"rating,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}
