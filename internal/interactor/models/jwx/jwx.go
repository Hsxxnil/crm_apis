package jwx

// JWX struct is used to create token
type JWX struct {
	// 公司ID
	CompanyID *string `json:"company_id,omitempty"`
	// 中文名稱
	Name *string `json:"name,omitempty"`
	// 使用者ID
	UserID *string `json:"user_id,omitempty"`
	// 角色ID
	RoleID *string `json:"role_id,omitempty"`
}

// Token return structure file
type Token struct {
	// 授權令牌
	AccessToken string `json:"access_token,omitempty"`
	// 刷新令牌
	RefreshToken string `json:"refresh_token,omitempty"`
}

// Refresh struct is used to refresh token
type Refresh struct {
	// 刷新令牌
	RefreshToken string `json:"refresh_token,omitempty" binding:"required" validate:"required"`
}
