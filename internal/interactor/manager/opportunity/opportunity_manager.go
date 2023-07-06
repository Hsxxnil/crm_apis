package opportunity

import (
	"encoding/json"
	"errors"

	campaignModel "app.eirc/internal/interactor/models/campaigns"
	leadModel "app.eirc/internal/interactor/models/leads"
	opportunityModel "app.eirc/internal/interactor/models/opportunities"
	"app.eirc/internal/interactor/pkg/util"
	campaignService "app.eirc/internal/interactor/service/campaign"
	leadService "app.eirc/internal/interactor/service/lead"
	opportunityService "app.eirc/internal/interactor/service/opportunity"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *opportunityModel.Create) (int, interface{})
	GetByList(input *opportunityModel.Fields) (int, interface{})
	GetBySingle(input *opportunityModel.Field) (int, interface{})
	GetBySingleCampaigns(input *opportunityModel.Field) (int, interface{})
	Delete(input *opportunityModel.Field) (int, interface{})
	Update(input *opportunityModel.Update) (int, interface{})
}

type manager struct {
	OpportunityService opportunityService.Service
	CampaignService    campaignService.Service
	LeadService        leadService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		OpportunityService: opportunityService.Init(db),
		CampaignService:    campaignService.Init(db),
		LeadService:        leadService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *opportunityModel.Create) (int, interface{}) {
	defer trx.Rollback()

	// 若由線索轉換則同步線索的account_id
	if input.LeadID != "" {
		leadBase, _ := m.LeadService.GetBySingle(&leadModel.Field{
			LeadID: input.LeadID,
		})
		input.AccountID = *leadBase.AccountID
	}

	opportunityBase, err := m.OpportunityService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, opportunityBase.OpportunityID)
}

func (m *manager) GetByList(input *opportunityModel.Fields) (int, interface{}) {
	output := &opportunityModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, opportunityBase, err := m.OpportunityService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	opportunityByte, err := json.Marshal(opportunityBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(opportunityByte, &output.Opportunities)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, opportunities := range output.Opportunities {
		opportunities.AccountName = *opportunityBase[i].Accounts.Name
		opportunities.CreatedBy = *opportunityBase[i].CreatedByUsers.Name
		opportunities.UpdatedBy = *opportunityBase[i].UpdatedByUsers.Name
		opportunities.SalespersonName = *opportunityBase[i].Salespeople.Name
		opportunities.LeadDescription = *opportunityBase[i].Leads.Description
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *opportunityModel.Field) (int, interface{}) {
	opportunityBase, err := m.OpportunityService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &opportunityModel.Single{}
	opportunityByte, _ := json.Marshal(opportunityBase)
	err = json.Unmarshal(opportunityByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.AccountName = *opportunityBase.Accounts.Name
	output.CreatedBy = *opportunityBase.CreatedByUsers.Name
	output.UpdatedBy = *opportunityBase.UpdatedByUsers.Name
	output.SalespersonName = *opportunityBase.Salespeople.Name
	output.LeadDescription = *opportunityBase.Leads.Description

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingleCampaigns(input *opportunityModel.Field) (int, interface{}) {
	opportunityBase, err := m.OpportunityService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &opportunityModel.SingleCampaigns{}
	opportunityByte, _ := json.Marshal(opportunityBase)
	err = json.Unmarshal(opportunityByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.AccountName = *opportunityBase.Accounts.Name
	output.CreatedBy = *opportunityBase.CreatedByUsers.Name
	output.UpdatedBy = *opportunityBase.UpdatedByUsers.Name
	output.SalespersonName = *opportunityBase.Salespeople.Name
	output.LeadDescription = *opportunityBase.Leads.Description
	for i, campaigns := range opportunityBase.OpportunityCampaigns {
		campaignBase, _ := m.CampaignService.GetBySingle(&campaignModel.Field{
			CampaignID: *campaigns.CampaignID,
		})
		output.OpportunityCampaigns[i].CampaignName = *campaignBase.Name
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *opportunityModel.Field) (int, interface{}) {
	_, err := m.OpportunityService.GetBySingle(&opportunityModel.Field{
		OpportunityID: input.OpportunityID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.OpportunityService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *opportunityModel.Update) (int, interface{}) {
	opportunityBase, err := m.OpportunityService.GetBySingle(&opportunityModel.Field{
		OpportunityID: input.OpportunityID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.OpportunityService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, opportunityBase.OpportunityID)
}
