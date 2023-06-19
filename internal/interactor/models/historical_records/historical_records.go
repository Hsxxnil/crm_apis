package historical_records

import (
	"app.eirc/internal/interactor/models/page"
	"app.eirc/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 來源ID
	SourceID string `json:"source_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 歷程記錄內容
	Content string `json:"content,omitempty" binding:"required" validate:"required"`
	// 歷程記錄動作
	Action string `json:"action,omitempty" binding:"required" validate:"required"`
	// 異動者
	ModifiedBy string `json:"modified_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 歷程記錄ID
	HistoricalRecordID string `json:"historical_record_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 歷程記錄內容
	Content *string `json:"content,omitempty" form:"content"`
	// 歷程記錄動作
	Action *string `json:"action,omitempty" form:"action"`
	// 來源ID
	SourceID *string `json:"source_id,omitempty" form:"source_id"`
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
	HistoricalRecords []*struct {
		// 歷程記錄ID
		HistoricalRecordID string `json:"historical_record_id,omitempty"`
		// 來源ID
		SourceID string `json:"source_id,omitempty"`
		// 歷程記錄描述
		Description string `json:"description,omitempty"`
		// 異動者
		ModifiedBy string `json:"modified_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"historical_records"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 歷程記錄ID
	HistoricalRecordID string `json:"historical_record_id,omitempty"`
	// 來源ID
	SourceID string `json:"source_id,omitempty"`
	// 歷程記錄描述
	Description string `json:"description,omitempty"`
	// 異動者
	ModifiedBy string `json:"modified_by,omitempty"`
	// 時間戳記
	section.TimeAt
}
