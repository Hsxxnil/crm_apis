package opportunity

import (
	"encoding/json"
	"errors"
	"strconv"

	"crm/internal/interactor/helpers"
	historicalRecordModel "crm/internal/interactor/models/historical_records"
	userModel "crm/internal/interactor/models/users"
	historicalRecordService "crm/internal/interactor/service/historical_record"
	userService "crm/internal/interactor/service/user"

	campaignModel "crm/internal/interactor/models/campaigns"
	leadModel "crm/internal/interactor/models/leads"
	opportunityModel "crm/internal/interactor/models/opportunities"
	"crm/internal/interactor/pkg/util"
	campaignService "crm/internal/interactor/service/campaign"
	leadService "crm/internal/interactor/service/lead"
	opportunityService "crm/internal/interactor/service/opportunity"

	"gorm.io/gorm"

	"crm/internal/interactor/pkg/util/code"
	"crm/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *opportunityModel.Create) (int, any)
	GetByList(input *opportunityModel.Fields) (int, any)
	GetByListNoPagination(input *opportunityModel.FieldsNoPagination) (int, any)
	GetBySingle(input *opportunityModel.Field) (int, any)
	GetBySingleCampaigns(input *opportunityModel.Field) (int, any)
	Delete(trx *gorm.DB, input *opportunityModel.Update) (int, any)
	Update(trx *gorm.DB, input *opportunityModel.Update) (int, any)
}

type manager struct {
	OpportunityService      opportunityService.Service
	CampaignService         campaignService.Service
	LeadService             leadService.Service
	HistoricalRecordService historicalRecordService.Service
	UserService             userService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		OpportunityService:      opportunityService.Init(db),
		CampaignService:         campaignService.Init(db),
		LeadService:             leadService.Init(db),
		HistoricalRecordService: historicalRecordService.Init(db),
		UserService:             userService.Init(db),
	}
}

const sourceType = "商機"

func (m *manager) Create(trx *gorm.DB, input *opportunityModel.Create) (int, any) {
	defer trx.Rollback()

	// 若由線索轉換
	if input.LeadID != "" {
		// 同步將線索狀態改為「已轉換」
		err := m.LeadService.WithTrx(trx).Update(&leadModel.Update{
			LeadID:    input.LeadID,
			Status:    util.PointerString("已轉換"),
			UpdatedBy: util.PointerString(input.CreatedBy),
		})
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		// 同步新增線索歷程記錄
		_, err = m.HistoricalRecordService.WithTrx(trx).Create(&historicalRecordModel.Create{
			SourceID:   input.LeadID,
			Action:     "修改",
			SourceType: "線索",
			Field:      "狀態為",
			Value:      "已轉換",
			ModifiedBy: input.CreatedBy,
		})
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		// 同步線索的account_id
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

	// 同步新增商機歷程記錄
	_, err = m.HistoricalRecordService.WithTrx(trx).Create(&historicalRecordModel.Create{
		SourceID:   *opportunityBase.OpportunityID,
		Action:     "建立",
		SourceType: sourceType,
		ModifiedBy: *opportunityBase.CreatedBy,
	})
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, opportunityBase.OpportunityID)
}

func (m *manager) GetByList(input *opportunityModel.Fields) (int, any) {
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

func (m *manager) GetByListNoPagination(input *opportunityModel.FieldsNoPagination) (int, any) {
	output := &opportunityModel.ListNoPagination{}
	opportunityBase, err := m.OpportunityService.GetByListNoPagination(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	opportunityByte, err := json.Marshal(opportunityBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	err = json.Unmarshal(opportunityByte, &output.Opportunities)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *opportunityModel.Field) (int, any) {
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

func (m *manager) GetBySingleCampaigns(input *opportunityModel.Field) (int, any) {
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

func (m *manager) Delete(trx *gorm.DB, input *opportunityModel.Update) (int, any) {
	defer trx.Rollback()

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

	// 若是由線索轉換為商機，同步將線索狀態改為「發展中」
	if opportunityBase.LeadID != nil {
		err = m.LeadService.WithTrx(trx).Update(&leadModel.Update{
			LeadID:    *opportunityBase.LeadID,
			Status:    util.PointerString("發展中"),
			UpdatedBy: input.UpdatedBy,
		})
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		// 同步新增線索歷程記錄
		_, err = m.HistoricalRecordService.WithTrx(trx).Create(&historicalRecordModel.Create{
			SourceID:   *opportunityBase.LeadID,
			Action:     "修改",
			SourceType: "線索",
			Field:      "狀態為",
			Value:      "發展中",
			ModifiedBy: *input.UpdatedBy,
		})
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	err = m.OpportunityService.WithTrx(trx).Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(trx *gorm.DB, input *opportunityModel.Update) (int, any) {
	defer trx.Rollback()

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

	err = m.OpportunityService.WithTrx(trx).Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步新增商機歷程記錄
	var records []historicalRecordModel.AddHistoricalRecord

	if input.Name != nil && *input.Name != *opportunityBase.Name {
		helpers.AddHistoricalRecord(&records, "修改", "名稱為", *input.Name)
	}

	if input.Stage != nil && *input.Stage != *opportunityBase.Stage {
		helpers.AddHistoricalRecord(&records, "修改", "階段為", *input.Stage)
	}

	if input.ForecastCategory != nil && *input.ForecastCategory != *opportunityBase.ForecastCategory {
		helpers.AddHistoricalRecord(&records, "修改", "預測種類為", *input.ForecastCategory)
	}

	if input.CloseDate != nil && *input.CloseDate != *opportunityBase.CloseDate {
		helpers.AddHistoricalRecord(&records, "修改", "結束日期為", input.CloseDate.UTC().Format("2006-01-02T15:04:05.999999Z"))
	}

	if input.Amount != nil && *input.Amount != *opportunityBase.Amount {
		helpers.AddHistoricalRecord(&records, "修改", "金額為", strconv.FormatFloat(*input.Amount, 'f', -1, 64))
	} else if *opportunityBase.Amount != 0 {
		helpers.AddHistoricalRecord(&records, "清除", "金額", "")
	}

	if input.SalespersonID != nil && *input.SalespersonID != *opportunityBase.SalespersonID {
		salespersonBase, _ := m.UserService.GetBySingle(&userModel.Field{
			UserID: *input.SalespersonID,
		})
		helpers.AddHistoricalRecord(&records, "修改", "業務員為", *salespersonBase.Name)
	}

	for _, record := range records {
		_, err = m.HistoricalRecordService.WithTrx(trx).Create(&historicalRecordModel.Create{
			SourceID:   *opportunityBase.OpportunityID,
			Action:     record.Actions,
			SourceType: sourceType,
			Field:      record.Fields,
			Value:      record.Values,
			ModifiedBy: *input.UpdatedBy,
		})
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, opportunityBase.OpportunityID)
}
