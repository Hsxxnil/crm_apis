package campaign

import (
	"encoding/json"
	"errors"

	opportunityModel "app.eirc/internal/interactor/models/opportunities"
	opportunityService "app.eirc/internal/interactor/service/opportunity"

	"app.eirc/internal/interactor/pkg/util"

	campaignModel "app.eirc/internal/interactor/models/campaigns"
	campaignService "app.eirc/internal/interactor/service/campaign"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *campaignModel.Create) interface{}
	GetByList(input *campaignModel.Fields) interface{}
	GetByListOpportunities(input *campaignModel.Fields) interface{}
	GetBySingle(input *campaignModel.Field) interface{}
	GetBySingleOpportunities(input *campaignModel.Field) interface{}
	Delete(input *campaignModel.Field) interface{}
	Update(input *campaignModel.Update) interface{}
}

type manager struct {
	CampaignService    campaignService.Service
	OpportunityService opportunityService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		CampaignService:    campaignService.Init(db),
		OpportunityService: opportunityService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *campaignModel.Create) interface{} {
	defer trx.Rollback()

	campaignBase, err := m.CampaignService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, campaignBase.CampaignID)
}

func (m *manager) GetByList(input *campaignModel.Fields) interface{} {
	output := &campaignModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, campaignBase, err := m.CampaignService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	campaignByte, err := json.Marshal(campaignBase)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(campaignByte, &output.Campaigns)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, campaigns := range output.Campaigns {
		campaigns.CreatedBy = *campaignBase[i].CreatedByUsers.Name
		campaigns.UpdatedBy = *campaignBase[i].UpdatedByUsers.Name
		if parentCampaignsBase, err := m.CampaignService.GetBySingle(&campaignModel.Field{
			CampaignID: campaigns.ParentCampaignID,
		}); err != nil {
			campaigns.ParentCampaignName = ""
		} else {
			campaigns.ParentCampaignName = *parentCampaignsBase.Name
		}
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetByListOpportunities(input *campaignModel.Fields) interface{} {
	output := &campaignModel.ListOpportunities{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, campaignBase, err := m.CampaignService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	campaignByte, err := json.Marshal(campaignBase)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(campaignByte, &output.Campaigns)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, campaigns := range output.Campaigns {
		campaigns.CreatedBy = *campaignBase[i].CreatedByUsers.Name
		campaigns.UpdatedBy = *campaignBase[i].UpdatedByUsers.Name
		if parentCampaignsBase, err := m.CampaignService.GetBySingle(&campaignModel.Field{
			CampaignID: campaigns.ParentCampaignID,
		}); err != nil {
			campaigns.ParentCampaignName = ""
		} else {
			campaigns.ParentCampaignName = *parentCampaignsBase.Name
		}
		for j, opportunities := range campaignBase[i].OpportunityCampaigns {
			opportunityBase, _ := m.OpportunityService.GetBySingle(&opportunityModel.Field{
				OpportunityID: *opportunities.OpportunityID,
			})
			campaigns.OpportunityCampaigns[j].OpportunityName = *opportunityBase.Name
		}
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *campaignModel.Field) interface{} {
	campaignBase, err := m.CampaignService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &campaignModel.Single{}
	campaignByte, _ := json.Marshal(campaignBase)
	err = json.Unmarshal(campaignByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output.CreatedBy = *campaignBase.CreatedByUsers.Name
	output.UpdatedBy = *campaignBase.UpdatedByUsers.Name
	if parentCampaignsBase, err := m.CampaignService.GetBySingle(&campaignModel.Field{
		CampaignID: *campaignBase.ParentCampaignID,
	}); err != nil {
		output.ParentCampaignName = ""
	} else {
		output.ParentCampaignName = *parentCampaignsBase.Name
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingleOpportunities(input *campaignModel.Field) interface{} {
	campaignBase, err := m.CampaignService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &campaignModel.SingleOpportunities{}
	campaignByte, _ := json.Marshal(campaignBase)
	err = json.Unmarshal(campaignByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output.CreatedBy = *campaignBase.CreatedByUsers.Name
	output.UpdatedBy = *campaignBase.UpdatedByUsers.Name
	if parentCampaignsBase, err := m.CampaignService.GetBySingle(&campaignModel.Field{
		CampaignID: *campaignBase.ParentCampaignID,
	}); err != nil {
		output.ParentCampaignName = ""
	} else {
		output.ParentCampaignName = *parentCampaignsBase.Name
	}
	for i, opportunities := range campaignBase.OpportunityCampaigns {
		opportunityBase, _ := m.OpportunityService.GetBySingle(&opportunityModel.Field{
			OpportunityID: *opportunities.OpportunityID,
		})
		output.OpportunityCampaigns[i].OpportunityName = *opportunityBase.Name
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *campaignModel.Field) interface{} {
	_, err := m.CampaignService.GetBySingle(&campaignModel.Field{
		CampaignID: input.CampaignID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.CampaignService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *campaignModel.Update) interface{} {
	campaignBase, err := m.CampaignService.GetBySingle(&campaignModel.Field{
		CampaignID: input.CampaignID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.CampaignService.Update(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, campaignBase.CampaignID)
}
