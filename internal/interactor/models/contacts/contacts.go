package contacts

import (
	"app.eirc/internal/interactor/models/page"
	"app.eirc/internal/interactor/models/section"
	"app.eirc/internal/interactor/models/sort"
)

// Create struct is used to create achieves
type Create struct {
	// 聯絡人名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	// 聯絡人職稱
	Title string `json:"title,omitempty"`
	// 聯絡人電話
	PhoneNumber string `json:"phone_number,omitempty" binding:"required" validate:"required"`
	// 聯絡人行動電話
	CellPhone string `json:"cell_phone,omitempty"`
	// 聯絡人電子郵件
	Email string `json:"email,omitempty"`
	// 聯絡人稱謂
	Salutation string `json:"salutation,omitempty"`
	// 聯絡人部門
	Department string `json:"department,omitempty"`
	// 聯絡人直屬上司ID
	SupervisorID string `json:"supervisor_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 帳戶ID
	AccountID string `json:"account_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}

// Field is structure file for search
type Field struct {
	// 聯絡人ID
	ContactID string `json:"contact_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 聯絡人名稱
	Name *string `json:"name,omitempty" form:"name"`
	// 聯絡人職稱
	Title *string `json:"title,omitempty" form:"title"`
	// 聯絡人電話
	PhoneNumber *string `json:"phone_number,omitempty" form:"phone_number"`
	// 聯絡人行動電話
	CellPhone *string `json:"cell_phone,omitempty" form:"cell_phone"`
	// 聯絡人電子郵件
	Email *string `json:"email,omitempty" form:"email"`
	// 聯絡人稱謂
	Salutation *string `json:"salutation,omitempty" form:"salutation"`
	// 聯絡人部門
	Department *string `json:"department,omitempty" form:"department"`
	// 聯絡人直屬上司ID
	SupervisorID *string `json:"supervisor_id,omitempty" form:"supervisor_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4" `
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty" form:"account_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4" `
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

// Filter struct is used to store the search field
type Filter struct {
	// 聯絡人名稱
	FilterName *string `json:"name,omitempty"`
	// TODO 帳戶名稱
	// 聯絡人行動電話
	FilterCellPhone *string `json:"cell_phone,omitempty"`
	// 聯絡人電子郵件
	FilterEmail *string `json:"email,omitempty"`
	// TODO 業務員名稱
}

// List is multiple return structure files
type List struct {
	// 多筆
	Contacts []*struct {
		// 聯絡人ID
		ContactID string `json:"contact_id,omitempty"`
		// 聯絡人名稱
		Name string `json:"name,omitempty"`
		// 聯絡人職稱
		Title string `json:"title,omitempty"`
		// 聯絡人電話
		PhoneNumber string `json:"phone_number,omitempty"`
		// 聯絡人行動電話
		CellPhone string `json:"cell_phone,omitempty"`
		// 聯絡人電子郵件
		Email string `json:"email,omitempty"`
		// 聯絡人稱謂
		Salutation string `json:"salutation,omitempty"`
		// 聯絡人部門
		Department string `json:"department,omitempty"`
		// 聯絡人直屬上司ID
		SupervisorID string `json:"supervisor_id,omitempty"`
		// 聯絡人直屬上司名稱
		SupervisorName string `json:"supervisor_name,omitempty"`
		// 帳戶ID
		AccountID string `json:"account_id,omitempty"`
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
	} `json:"contacts"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 聯絡人ID
	ContactID string `json:"contact_id,omitempty"`
	// 聯絡人名稱
	Name string `json:"name,omitempty"`
	// 聯絡人職稱
	Title string `json:"title,omitempty"`
	// 聯絡人電話
	PhoneNumber string `json:"phone_number,omitempty"`
	// 聯絡人行動電話
	CellPhone string `json:"cell_phone,omitempty"`
	// 聯絡人電子郵件
	Email string `json:"email,omitempty"`
	// 聯絡人稱謂
	Salutation string `json:"salutation,omitempty"`
	// 聯絡人部門
	Department string `json:"department,omitempty"`
	// 聯絡人直屬上司ID
	SupervisorID string `json:"supervisor_id,omitempty"`
	// 聯絡人直屬上司名稱
	SupervisorName string `json:"supervisor_name,omitempty"`
	// 帳戶ID
	AccountID string `json:"account_id,omitempty"`
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

// Update struct is used to update achieves
type Update struct {
	// 聯絡人ID
	ContactID string `json:"contact_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 聯絡人名稱
	Name *string `json:"name,omitempty"`
	// 聯絡人職稱
	Title *string `json:"title,omitempty"`
	// 聯絡人電話
	PhoneNumber *string `json:"phone_number,omitempty"`
	// 聯絡人行動電話
	CellPhone *string `json:"cell_phone,omitempty"`
	// 聯絡人電子郵件
	Email *string `json:"email,omitempty"`
	// 聯絡人稱謂
	Salutation *string `json:"salutation,omitempty"`
	// 聯絡人部門
	Department *string `json:"department,omitempty"`
	// 聯絡人直屬上司ID
	SupervisorID *string `json:"supervisor_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 帳戶ID
	AccountID *string `json:"account_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 業務員ID
	SalespersonID *string `json:"salesperson_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}

// AccountSingle return structure file for accounts
type AccountSingle struct {
	// 聯絡人ID
	ContactID string `json:"contact_id,omitempty"`
	// 聯絡人名稱
	Name string `json:"name,omitempty"`
	// 聯絡人職稱
	Title string `json:"title,omitempty"`
	// 聯絡人電話
	PhoneNumber string `json:"phone_number,omitempty"`
	// 聯絡人行動電話
	CellPhone string `json:"cell_phone,omitempty"`
	// 聯絡人電子郵件
	Email string `json:"email,omitempty"`
	// 聯絡人稱謂
	Salutation string `json:"salutation,omitempty"`
	// 聯絡人部門
	Department string `json:"department,omitempty"`
}
