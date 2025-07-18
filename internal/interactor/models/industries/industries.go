package industries

// Create struct is used to create achieves
type Create struct {
	// 行業名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
}

// Field is structure file for search
type Field struct {
	// 行業ID
	IndustryID string `json:"industry_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 行業名稱
	Name *string `json:"name,omitempty" form:"name"`
}

// List is multiple return structure files
type List struct {
	// 多筆
	Industries []*struct {
		// 行業ID
		IndustryID string `json:"industry_id,omitempty"`
		// 行業名稱
		Name string `json:"name,omitempty"`
	} `json:"industries"`
}

// Single return structure file
type Single struct {
	// 行業ID
	IndustryID string `json:"industry_id,omitempty"`
	// 行業名稱
	Name string `json:"name,omitempty"`
}

// Update struct is used to update achieves
type Update struct {
	// 行業ID
	IndustryID string `json:"industry_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 行業名稱
	Name *string `json:"name,omitempty"`
}
