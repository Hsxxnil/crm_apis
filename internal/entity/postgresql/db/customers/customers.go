package customers

import (
	"app.eirc/internal/interactor/models/special"
)

// Table struct is customers database table struct
type Table struct {
	//客戶編號
	CID string `gorm:"column:c_id;type:uuid;not null;primaryKey;" json:"c_id"`
	//簡稱
	ShortName string `gorm:"column:short_name;type:text;not null;" json:"short_name"`
	//英文名稱
	EngName string `gorm:"column:eng_name;type:text;" json:"eng_name,omitempty"`
	//中文名稱
	Name string `gorm:"column:name;type:text;not null;" json:"name"`
	//郵遞區號
	ZipCode string `gorm:"column:zip_code;type:text;not null;" json:"zip_code"`
	//地址
	Address string `gorm:"column:address;type:text;not null;" json:"address"`
	//電話
	Tel string `gorm:"column:tel;type:text;not null;" json:"tel"`
	//傳真
	Fax string `gorm:"column:fax;type:text;" json:"fax,omitempty"`
	//地圖
	Map string `gorm:"column:map;type:text;" json:"map,omitempty"`
	//聯絡人
	Liaison string `gorm:"column:liaison;type:text;" json:"liaison,omitempty"`
	//電子郵件
	Mail string `gorm:"column:mail;type:text;" json:"mail,omitempty"`
	//聯絡人手機號碼90
	LiaisonPhone string `gorm:"column:liaison_phone;type:text;" json:"liaison_phone,omitempty"`
	//統編
	TaxIdNumber string `gorm:"column:tax_id_number;type:text;not null;" json:"tax_id_number"`
	//備註
	Remark string `gorm:"column:remark;type:text;" json:"remark,omitempty"`
	special.UseTable
}

// Base struct is corresponding to customers table structure file
type Base struct {
	//客戶編號
	CID *string `json:"c_id,omitempty"`
	//簡稱
	ShortName *string `json:"short_name,omitempty"`
	//英文名稱
	EngName *string `json:"eng_name,omitempty"`
	//中文名稱
	Name *string `json:"name,omitempty"`
	//郵遞區號
	ZipCode *string `json:"zip_code,omitempty"`
	//地址
	Address *string `json:"address,omitempty"`
	//電話
	Tel *string `json:"tel,omitempty"`
	//傳真
	Fax *string `json:"fax,omitempty"`
	//地圖
	Map *string `json:"map,omitempty"`
	//聯絡人
	Liaison *string `json:"liaison,omitempty"`
	//電子郵件
	Mail *string `json:"mail,omitempty"`
	//聯絡人手機號碼
	LiaisonPhone *string `json:"liaison_phone,omitempty"`
	//統編
	TaxIdNumber *string `json:"tax_id_number,omitempty"`
	//備註
	Remark *string `json:"remark,omitempty"`
	special.UseBase
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "crm_customers"
}
