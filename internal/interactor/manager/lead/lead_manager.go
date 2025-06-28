package lead

import (
	"encoding/json"
	"errors"

	"crm/internal/interactor/helpers"

	historicalRecordModel "crm/internal/interactor/models/historical_records"
	userModel "crm/internal/interactor/models/users"
	historicalRecordService "crm/internal/interactor/service/historical_record"
	userService "crm/internal/interactor/service/user"

	"crm/internal/interactor/pkg/util"

	leadModel "crm/internal/interactor/models/leads"
	leadService "crm/internal/interactor/service/lead"

	"gorm.io/gorm"

	"crm/internal/interactor/pkg/util/code"
	"crm/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *leadModel.Create) (int, any)
	GetByList(input *leadModel.Fields) (int, any)
	GetByListNoPagination(input *leadModel.FieldsNoPagination) (int, any)
	GetBySingle(input *leadModel.Field) (int, any)
	Delete(input *leadModel.Field) (int, any)
	Update(trx *gorm.DB, input *leadModel.Update) (int, any)
}

type manager struct {
	LeadService             leadService.Service
	HistoricalRecordService historicalRecordService.Service
	UserService             userService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		LeadService:             leadService.Init(db),
		HistoricalRecordService: historicalRecordService.Init(db),
		UserService:             userService.Init(db),
	}
}

const sourceType = "線索"

func (m *manager) Create(trx *gorm.DB, input *leadModel.Create) (int, any) {
	defer trx.Rollback()

	leadBase, err := m.LeadService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步新增線索歷程記錄
	_, err = m.HistoricalRecordService.WithTrx(trx).Create(&historicalRecordModel.Create{
		SourceID:   *leadBase.LeadID,
		Action:     "建立",
		SourceType: sourceType,
		ModifiedBy: *leadBase.CreatedBy,
	})
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, leadBase.LeadID)
}

func (m *manager) GetByList(input *leadModel.Fields) (int, any) {
	output := &leadModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, leadBase, err := m.LeadService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	leadByte, err := json.Marshal(leadBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(leadByte, &output.Leads)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, leads := range output.Leads {
		leads.AccountName = *leadBase[i].Accounts.Name
		leads.CreatedBy = *leadBase[i].CreatedByUsers.Name
		leads.UpdatedBy = *leadBase[i].UpdatedByUsers.Name
		leads.SalespersonName = *leadBase[i].Salespeople.Name
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetByListNoPagination(input *leadModel.FieldsNoPagination) (int, any) {
	output := &leadModel.ListNoPagination{}
	leadBase, err := m.LeadService.GetByListNoPagination(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	leadByte, err := json.Marshal(leadBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	err = json.Unmarshal(leadByte, &output.Leads)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *leadModel.Field) (int, any) {
	leadBase, err := m.LeadService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &leadModel.Single{}
	leadByte, _ := json.Marshal(leadBase)
	err = json.Unmarshal(leadByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.AccountName = *leadBase.Accounts.Name
	output.CreatedBy = *leadBase.CreatedByUsers.Name
	output.UpdatedBy = *leadBase.UpdatedByUsers.Name
	output.SalespersonName = *leadBase.Salespeople.Name

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *leadModel.Field) (int, any) {
	_, err := m.LeadService.GetBySingle(&leadModel.Field{
		LeadID: input.LeadID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.LeadService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(trx *gorm.DB, input *leadModel.Update) (int, any) {
	defer trx.Rollback()

	leadBase, err := m.LeadService.GetBySingle(&leadModel.Field{
		LeadID: input.LeadID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.LeadService.WithTrx(trx).Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步新增線索歷程記錄
	var records []historicalRecordModel.AddHistoricalRecord

	if input.Status != nil && *input.Status != *leadBase.Status {
		helpers.AddHistoricalRecord(&records, "修改", "狀態為", *input.Status)
	}

	if input.Description != nil && *input.Description != *leadBase.Description {
		helpers.AddHistoricalRecord(&records, "修改", "描述為", *input.Description)
	}

	if input.Source != nil {
		if *input.Source != *leadBase.Source {
			if *input.Source == "" {
				helpers.AddHistoricalRecord(&records, "移除", "來源", "")
			} else {
				helpers.AddHistoricalRecord(&records, "修改", "來源為", *input.Source)
			}
		}
	} else if *leadBase.Source != "" {
		helpers.AddHistoricalRecord(&records, "移除", "來源", "")
	}

	if input.Rating != nil {
		if *input.Rating != *leadBase.Rating {
			if *input.Rating == "" {
				helpers.AddHistoricalRecord(&records, "清除", "分級", "")
			} else {
				helpers.AddHistoricalRecord(&records, "修改", "分級為", *input.Rating)
			}
		}
	} else if *leadBase.Rating != "" {
		helpers.AddHistoricalRecord(&records, "清除", "分級", "")
	}

	if input.SalespersonID != nil && *input.SalespersonID != *leadBase.SalespersonID {
		salespersonBase, _ := m.UserService.GetBySingle(&userModel.Field{
			UserID: *input.SalespersonID,
		})
		helpers.AddHistoricalRecord(&records, "修改", "業務員為", *salespersonBase.Name)
	}

	for _, record := range records {
		_, err = m.HistoricalRecordService.WithTrx(trx).Create(&historicalRecordModel.Create{
			SourceID:   *leadBase.LeadID,
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
	return code.Successful, code.GetCodeMessage(code.Successful, leadBase.LeadID)
}
