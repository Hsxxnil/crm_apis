package logins

type Login struct {
	// 公司ID
	CompanyID string `json:"company_id,omitempty" binding:"required" validate:"required"`
	// 使用者名稱
	UserName string `json:"user_name,omitempty" binding:"required" validate:"required"`
	// 密碼
	Password string `json:"password,omitempty" binding:"required" validate:"required"`
}
