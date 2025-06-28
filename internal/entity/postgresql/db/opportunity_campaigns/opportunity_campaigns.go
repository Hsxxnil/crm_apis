package opportunity_campaigns

import (
	"crm/internal/entity/postgresql/db/users"

	"crm/internal/interactor/models/special"
)

// Table struct is opportunity_campaigns database table struct
type Table struct {
	// 商機行銷活動ID
	OpportunityCampaignID string `gorm:"<-:create;column:opportunity_campaign_id;type:uuid;not null;primaryKey;" json:"opportunity_campaign_id"`
	// 商機ID
	OpportunityID string `gorm:"column:opportunity_id;type:uuid;not null;" json:"opportunity_id"`
	// 行銷活動ID
	CampaignID string `gorm:"column:campaign_id;type:uuid;not null;" json:"campaign_id"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:CreatedBy;references:UserID" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:UpdatedBy;references:UserID" json:"updated_by_users,omitempty"`
	special.Table
}

// Base struct is corresponding to opportunity_campaigns table structure file
type Base struct {
	// 商機行銷活動ID
	OpportunityCampaignID *string `json:"opportunity_campaign_id,omitempty"`
	// 商機ID
	OpportunityID *string `json:"opportunity_id,omitempty"`
	// 行銷活動ID
	CampaignID *string `json:"campaign_id,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.Base
}

// TableName sets the insert table opportunity_id for this struct type
func (t *Table) TableName() string {
	return "opportunity_campaigns"
}
