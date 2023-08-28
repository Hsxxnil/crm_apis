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
	Create(trx *gorm.DB, input *campaignModel.Create) (int, any)
	GetByList(input *campaignModel.Fields) (int, any)
	GetByListNoPagination(input *campaignModel.Field) (int, any)
	GetBySingle(input *campaignModel.Field) (int, any)
	GetBySingleOpportunities(input *campaignModel.Field) (int, any)
	Delete(input *campaignModel.Field) (int, any)
	Update(input *campaignModel.Update) (int, any)
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

func (m *manager) Create(trx *gorm.DB, input *campaignModel.Create) (int, any) {
	defer trx.Rollback()

	campaignBase, err := m.CampaignService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, campaignBase.CampaignID)
}

func (m *manager) GetByList(input *campaignModel.Fields) (int, any) {
	output := &campaignModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, campaignBase, err := m.CampaignService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	campaignByte, err := json.Marshal(campaignBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(campaignByte, &output.Campaigns)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, campaigns := range output.Campaigns {
		campaigns.CreatedBy = *campaignBase[i].CreatedByUsers.Name
		campaigns.UpdatedBy = *campaignBase[i].UpdatedByUsers.Name
		campaigns.SalespersonName = *campaignBase[i].Salespeople.Name
		if campaigns.ParentCampaignID != "" {
			parentCampaignsBase, err := m.CampaignService.GetBySingle(&campaignModel.Field{
				CampaignID: campaigns.ParentCampaignID,
			})
			if err != nil {
				campaigns.ParentCampaignName = ""
			} else {
				campaigns.ParentCampaignName = *parentCampaignsBase.Name
			}
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetByListNoPagination(input *campaignModel.Field) (int, any) {
	output := &campaignModel.ListNoPagination{}
	campaignBase, err := m.CampaignService.GetByListNoPagination(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	campaignByte, err := json.Marshal(campaignBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	err = json.Unmarshal(campaignByte, &output.Campaigns)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *campaignModel.Field) (int, any) {
	campaignBase, err := m.CampaignService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &campaignModel.Single{}
	campaignByte, _ := json.Marshal(campaignBase)
	err = json.Unmarshal(campaignByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.CreatedBy = *campaignBase.CreatedByUsers.Name
	output.UpdatedBy = *campaignBase.UpdatedByUsers.Name
	output.SalespersonName = *campaignBase.Salespeople.Name
	if campaignBase.ParentCampaignID != nil {
		parentCampaignsBase, err := m.CampaignService.GetBySingle(&campaignModel.Field{
			CampaignID: *campaignBase.ParentCampaignID,
		})
		if err != nil {
			output.ParentCampaignName = ""
		} else {
			output.ParentCampaignName = *parentCampaignsBase.Name
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingleOpportunities(input *campaignModel.Field) (int, any) {
	campaignBase, err := m.CampaignService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &campaignModel.SingleOpportunities{}
	campaignByte, _ := json.Marshal(campaignBase)
	err = json.Unmarshal(campaignByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.CreatedBy = *campaignBase.CreatedByUsers.Name
	output.UpdatedBy = *campaignBase.UpdatedByUsers.Name
	output.SalespersonName = *campaignBase.Salespeople.Name
	if campaignBase.ParentCampaignID != nil {
		parentCampaignsBase, err := m.CampaignService.GetBySingle(&campaignModel.Field{
			CampaignID: *campaignBase.ParentCampaignID,
		})
		if err != nil {
			output.ParentCampaignName = ""
		} else {
			output.ParentCampaignName = *parentCampaignsBase.Name
		}
	}
	for i, opportunities := range campaignBase.OpportunityCampaigns {
		opportunityBase, _ := m.OpportunityService.GetBySingle(&opportunityModel.Field{
			OpportunityID: *opportunities.OpportunityID,
		})
		output.OpportunityCampaigns[i].OpportunityName = *opportunityBase.Name
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *campaignModel.Field) (int, any) {
	_, err := m.CampaignService.GetBySingle(&campaignModel.Field{
		CampaignID: input.CampaignID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.CampaignService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *campaignModel.Update) (int, any) {
	campaignBase, err := m.CampaignService.GetBySingle(&campaignModel.Field{
		CampaignID: input.CampaignID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.CampaignService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, campaignBase.CampaignID)
}
