package opportunity_campaigns

import (
	"app.eirc/internal/interactor/models/page"
	"app.eirc/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 商機ID
	OpportunityID string `json:"opportunity_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 行銷活動ID
	CampaignID string `json:"campaign_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}

// Field is structure file for search
type Field struct {
	// 商機行銷活動ID
	OpportunityCampaignID string `json:"opportunity_campaign_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 商機ID
	OpportunityID *string `json:"opportunity_id,omitempty" form:"opportunity_id"`
	// 行銷活動ID
	CampaignID *string `json:"campaign_id,omitempty" form:"campaign_id"`
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
	OpportunityCampaigns []*struct {
		// 商機行銷活動ID
		OpportunityCampaignID string `json:"opportunity_campaign_id,omitempty"`
		// 商機ID
		OpportunityID string `json:"opportunity_id,omitempty"`
		// 行銷活動ID
		CampaignID string `json:"campaign_id,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"opportunity_campaigns"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 商機行銷活動ID
	OpportunityCampaignID string `json:"opportunity_campaign_id,omitempty"`
	// 商機ID
	OpportunityID string `json:"opportunity_id,omitempty"`
	// 行銷活動ID
	CampaignID string `json:"campaign_id,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
}

// Update struct is used to update achieves
type Update struct {
	// 商機行銷活動ID
	OpportunityCampaignID string `json:"opportunity_campaign_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 商機ID
	OpportunityID *string `json:"opportunity_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 行銷活動ID
	CampaignID *string `json:"campaign_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
}

type OpportunitySingle struct {
	// 行銷活動ID
	CampaignID string `json:"campaign_id,omitempty"`
	// 行銷活動名稱
	CampaignName string `json:"campaign_name,omitempty"`
}
