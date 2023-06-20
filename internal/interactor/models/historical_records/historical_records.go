package historical_records

import (
	"time"
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
		// 異動時間
		ModifiedAt *time.Time `json:"modified_at,omitempty"`
	} `json:"historical_records"`
	// 總筆數
	Total int64 `json:"total"`
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
	// 異動時間
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
}
