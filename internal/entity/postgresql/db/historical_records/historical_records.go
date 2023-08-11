package historical_records

import (
	"time"

	"app.eirc/internal/entity/postgresql/db/users"
	"app.eirc/internal/interactor/models/special"
)

// Table struct is historical_records database table struct
type Table struct {
	// 歷程記錄ID
	HistoricalRecordID string `gorm:"<-:create;column:historical_record_id;type:uuid;not null;primaryKey;" json:"historical_record_id"`
	// 來源ID
	SourceID string `gorm:"column:source_id;type:uuid;not null;" json:"source_id"`
	// 歷程記錄來源類型
	SourceType string `gorm:"column:source_type;type:text;not null;" json:"source_type"`
	// 歷程記錄欄位
	Field string `gorm:"column:field;type:text;not null;" json:"field"`
	// 歷程記錄異動值
	Value string `gorm:"column:value;type:text;not null;" json:"value"`
	// 歷程記錄動作
	Action string `gorm:"column:action;type:text;not null;" json:"action"`
	// 異動時間
	ModifiedAt time.Time `gorm:"<-:create;column:modified_at;type:timestamp;not null;" json:"modified_at"`
	// 異動者
	ModifiedBy string `gorm:"<-:create;column:modified_by;type:uuid;not null;" json:"modified_by"`
	// modify_users data
	ModifiedByUsers users.Table `gorm:"foreignKey:ModifiedBy;references:UserID" json:"modified_by_users,omitempty"`
}

// Base struct is corresponding to historical_records table structure file
type Base struct {
	// 歷程記錄ID
	HistoricalRecordID *string `json:"historical_record_id,omitempty"`
	// 來源ID
	SourceID *string `json:"source_id,omitempty"`
	// 歷程記錄來源類型
	SourceType *string `json:"source_type,omitempty"`
	// 歷程記錄欄位
	Field *string `json:"field,omitempty"`
	// 歷程記錄異動值
	Value *string `json:"value,omitempty"`
	// 歷程記錄動作
	Action *string `json:"action,omitempty"`
	// modify_users data
	ModifiedByUsers users.Base `json:"modified_by_users,omitempty"`
	special.Base
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "historical_records"
}
