package accounts

import (
	"app.eirc/internal/interactor/models/account_contacts"
	"app.eirc/internal/interactor/models/page"
	"app.eirc/internal/interactor/models/section"
	"app.eirc/internal/interactor/models/sort"
)

// Create struct is used to create achieves
type Create struct {
	// 帳戶名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	// 帳戶電話
	PhoneNumber string `json:"phone_number,omitempty"`
	// 帳戶類型
	Type []string `json:"type,omitempty" binding:"required" validate:"required"`
	// 行業ID
	IndustryID string `json:"industry_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 父系帳戶ID
	ParentAccountID string `json:"parent_account_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 帳戶ID
	AccountID string `json:"account_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 帳戶名稱
	Name *string `json:"name,omitempty" form:"name"`
	// 帳戶電話
	PhoneNumber *string `json:"phone_number,omitempty" form:"phone_number"`
	// 帳戶類型
	Type *[]string `json:"type,omitempty" form:"type"`
	// 行業ID
	IndustryID *string `json:"industry_id,omitempty" form:"industry_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 父系帳戶ID
	ParentAccountID *string `json:"parent_account_id,omitempty" form:"parent_account_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 業務員ID
	SalespersonID *string `json:"salesperson_id,omitempty" form:"salesperson_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	// 搜尋結構檔
	Field
	// 搜尋欄位
	Filter `json:"filter"`
	// 分頁搜尋結構檔
	page.Pagination
	// 排序欄位
	sort.Sort `json:"sort"`
}

// FieldsNoPagination is the searched structure file (including filter)
type FieldsNoPagination struct {
	// 搜尋結構檔
	Field
	// 搜尋欄位
	FilterNoPagination `json:"filter"`
}

// Filter struct is used to store the search field
type Filter struct {
	// 帳戶名稱
	FilterName string `json:"name,omitempty"`
	// 帳戶電話
	FilterPhoneNumber string `json:"phone_number,omitempty"`
	// 帳戶類型
	FilterType []string `json:"type,omitempty"`
	// 業務員名稱
	FilterSalespersonName string `json:"salesperson_name,omitempty"`
}

// FilterNoPagination struct is used to store the search field no pagination
type FilterNoPagination struct {
	// 帳戶名稱
	FilterName string `json:"name,omitempty"`
	// 帳戶類型
	FilterType []string `json:"type,omitempty"`
}

// List is multiple return structure files
type List struct {
	// 多筆
	Accounts []*struct {
		// 帳戶ID
		AccountID string `json:"account_id,omitempty"`
		// 帳戶名稱
		Name string `json:"name,omitempty"`
		// 帳戶電話
		PhoneNumber string `json:"phone_number,omitempty"`
		// 帳戶類型
		Type []string `json:"type,omitempty"`
		// 行業ID
		IndustryID string `json:"industry_id,omitempty"`
		// 行業名稱
		IndustryName string `json:"industry_name,omitempty"`
		// 父系帳戶ID
		ParentAccountID string `json:"parent_account_id,omitempty"`
		// 父系帳戶名稱
		ParentAccountName string `json:"parent_account_name,omitempty"`
		// 業務員ID
		SalespersonID string `json:"salesperson_id,omitempty"`
		// 業務員名稱
		SalespersonName string `json:"salesperson_name,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"accounts"`
	// 分頁返回結構檔
	page.Total
}

// ListNoPagination is multiple return structure files without pagination
type ListNoPagination struct {
	// 多筆
	Accounts []*struct {
		// 帳戶ID
		AccountID string `json:"account_id,omitempty"`
		// 帳戶名稱
		Name string `json:"name,omitempty"`
	} `json:"accounts"`
}

// Single return structure file
type Single struct {
	// 帳戶ID
	AccountID string `json:"account_id,omitempty"`
	// 帳戶名稱
	Name string `json:"name,omitempty"`
	// 帳戶電話
	PhoneNumber string `json:"phone_number,omitempty"`
	// 帳戶類型
	Type []string `json:"type,omitempty"`
	// 行業ID
	IndustryID string `json:"industry_id,omitempty"`
	// 行業名稱
	IndustryName string `json:"industry_name,omitempty"`
	// 父系帳戶ID
	ParentAccountID string `json:"parent_account_id,omitempty"`
	// 父系帳戶名稱
	ParentAccountName string `json:"parent_account_name,omitempty"`
	// 業務員ID
	SalespersonID string `json:"salesperson_id,omitempty"`
	// 業務員名稱
	SalespersonName string `json:"salesperson_name,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
}

// SingleContacts return structure file containing contacts
type SingleContacts struct {
	// 帳戶ID
	AccountID string `json:"account_id,omitempty"`
	// 帳戶名稱
	Name string `json:"name,omitempty"`
	// 帳戶電話
	PhoneNumber string `json:"phone_number,omitempty"`
	// 帳戶類型
	Type []string `json:"type,omitempty"`
	// 行業ID
	IndustryID string `json:"industry_id,omitempty"`
	// 行業名稱
	IndustryName string `json:"industry_name,omitempty"`
	// 父系帳戶ID
	ParentAccountID string `json:"parent_account_id,omitempty"`
	// 父系帳戶名稱
	ParentAccountName string `json:"parent_account_name,omitempty"`
	// 業務員ID
	SalespersonID string `json:"salesperson_id,omitempty"`
	// 業務員名稱
	SalespersonName string `json:"salesperson_name,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
	// account_contacts data
	AccountContacts []account_contacts.AccountSingle `json:"contacts,omitempty"`
}

// Update struct is used to update achieves
type Update struct {
	// 帳戶ID
	AccountID string `json:"account_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 帳戶名稱
	Name *string `json:"name,omitempty"`
	// 帳戶電話
	PhoneNumber *string `json:"phone_number,omitempty"`
	// 帳戶類型
	Type *[]string `json:"type,omitempty"`
	// 行業ID
	IndustryID *string `json:"industry_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 父系帳戶ID
	ParentAccountID *string `json:"parent_account_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 業務員ID
	SalespersonID *string `json:"salesperson_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}
