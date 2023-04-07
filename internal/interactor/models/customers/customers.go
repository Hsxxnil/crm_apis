package customers

import (
	"app.eirc/internal/interactor/models/page"
	"app.eirc/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	//簡稱
	ShortName string `json:"short_name,omitempty" binding:"required" validate:"required"`
	//英文名稱
	EngName string `json:"eng_name,omitempty" validate:"required"`
	//中文名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	//郵遞區號
	ZipCode string `json:"zip_code,omitempty" binding:"required" validate:"required"`
	//地址
	Address string `json:"address,omitempty" binding:"required" validate:"required"`
	//電話
	Tel string `json:"tel,omitempty" binding:"required" validate:"required"`
	//傳真
	Fax string `json:"fax,omitempty" validate:"required"`
	//地圖
	Map string `json:"map,omitempty" validate:"required"`
	//聯絡人
	Liaison string `json:"liaison,omitempty" validate:"required"`
	//電子郵件
	Mail string `json:"mail,omitempty" validate:"required"`
	//聯絡人手機號碼
	LiaisonPhone string `json:"liaison_phone,omitempty" validate:"required"`
	//統編
	TaxIdNumber string `json:"tax_id_number,omitempty" binding:"required" validate:"required"`
	//備註
	Remark string `json:"remark,omitempty" validate:"required"`
	//創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required"`
}

// Field is structure file for search
type Field struct {
	//客戶編號
	CID string `json:"c_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	//簡稱
	ShortName *string `json:"short_name,omitempty" from:"short_name"`
	//英文名稱
	EngName *string `json:"eng_name,omitempty" from:"eng_name"`
	//中文名稱
	Name *string `json:"name,omitempty" from:"name"`
	//郵遞區號
	ZipCode *string `json:"zip_code,omitempty" from:"zip_code"`
	//地址
	Address *string `json:"address,omitempty" from:"address"`
	//電話
	Tel *string `json:"tel,omitempty" from:"tel"`
	//傳真
	Fax *string `json:"fax,omitempty" from:"fax"`
	//地圖
	Map *string `json:"map,omitempty" from:"map"`
	//聯絡人
	Liaison *string `json:"liaison,omitempty" from:"liaison"`
	//電子郵件
	Mail *string `json:"mail,omitempty" from:"mail"`
	//聯絡人手機號碼
	LiaisonPhone *string `json:"liaison_phone,omitempty" from:"liaison_phone"`
	//統編
	TaxIdNumber *string `json:"tax_id_number,omitempty" from:"tax_id_number"`
	//備註
	Remark *string `json:"remark,omitempty" from:"remark"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	//搜尋結構檔
	Field
	//分頁搜尋結構檔
	page.Pagination
}

// List is multiple return structure files
type List struct {
	//多筆
	Customers []*struct {
		//客戶編號
		CID string `json:"c_id,omitempty"`
		//簡稱
		ShortName string `json:"short_name,omitempty"`
		//英文名稱
		EngName string `json:"eng_name,omitempty"`
		//中文名稱
		Name string `json:"name,omitempty"`
		//郵遞區號
		ZipCode string `json:"zip_code,omitempty"`
		//地址
		Address string `json:"address,omitempty"`
		//電話
		Tel string `json:"tel,omitempty"`
		//傳真
		Fax string `json:"fax,omitempty"`
		//地圖
		Map string `json:"map,omitempty"`
		//聯絡人
		Liaison string `json:"liaison,omitempty"`
		//電子郵件
		Mail string `json:"mail,omitempty"`
		//聯絡人手機號碼
		LiaisonPhone string `json:"liaison_phone,omitempty"`
		//統編
		TaxIdNumber string `json:"tax_id_number,omitempty"`
		//備註
		Remark string `json:"remark,omitempty"`
		//創建者
		CreatedBy string `json:"created_by"`
		//時間戳記
		section.TimeAt
	} `json:"customers"`
	//分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	//客戶編號
	CID string `json:"c_id,omitempty"`
	//簡稱
	ShortName string `json:"short_name,omitempty"`
	//英文名稱
	EngName string `json:"eng_name,omitempty"`
	//中文名稱
	Name string `json:"name,omitempty"`
	//郵遞區號
	ZipCode string `json:"zip_code,omitempty"`
	//地址
	Address string `json:"address,omitempty"`
	//電話
	Tel string `json:"tel,omitempty"`
	//傳真
	Fax string `json:"fax,omitempty"`
	//地圖
	Map string `json:"map,omitempty"`
	//聯絡人
	Liaison string `json:"liaison,omitempty"`
	//電子郵件
	Mail string `json:"mail,omitempty"`
	//聯絡人手機號碼
	LiaisonPhone string `json:"liaison_phone,omitempty"`
	//統編
	TaxIdNumber string `json:"tax_id_number,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
	//創建者
	CreatedBy string `json:"created_by"`
	//時間戳記
	section.TimeAt
}

// Update struct is used to update achieves
type Update struct {
	//客戶編號
	CID string `json:"c_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	//簡稱
	ShortName string `json:"short_name,omitempty"`
	//英文名稱
	EngName string `json:"eng_name,omitempty"`
	//中文名稱
	Name string `json:"name,omitempty"`
	//郵遞區號
	ZipCode string `json:"zip_code,omitempty"`
	//地址
	Address string `json:"address,omitempty"`
	//電話
	Tel string `json:"tel,omitempty"`
	//傳真
	Fax string `json:"fax,omitempty"`
	//地圖
	Map string `json:"map,omitempty"`
	//聯絡人
	Liaison string `json:"liaison,omitempty"`
	//電子郵件
	Mail string `json:"mail,omitempty"`
	//聯絡人手機號碼
	LiaisonPhone string `json:"liaison_phone,omitempty"`
	//統編
	TaxIdNumber string `json:"tax_id_number,omitempty"`
	//備註
	Remark string `json:"remark,omitempty"`
}
