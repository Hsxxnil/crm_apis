package industries

import "app.eirc/internal/interactor/models/page"

// Table struct is industries database table struct
type Table struct {
	// 行業ID
	IndustryID string `gorm:"<-:create;column:industry_id;type:uuid;not null;primaryKey;" json:"industry_id"`
	// 行業名稱
	Name string `gorm:"column:name;type:text;not null;" json:"name"`
}

// Base struct is corresponding to industries table structure file
type Base struct {
	// 行業ID
	IndustryID *string `json:"industry_id,omitempty"`
	// 行業名稱
	Name *string `json:"name,omitempty"`
	// 引入page
	page.Pagination
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "crm_industries"
}
