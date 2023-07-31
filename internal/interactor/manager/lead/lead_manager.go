package lead

import (
	"encoding/json"
	"errors"

	historicalRecordModel "app.eirc/internal/interactor/models/historical_records"
	userModel "app.eirc/internal/interactor/models/users"
	historicalRecordService "app.eirc/internal/interactor/service/historical_record"
	userService "app.eirc/internal/interactor/service/user"

	"app.eirc/internal/interactor/pkg/util"

	leadModel "app.eirc/internal/interactor/models/leads"
	leadService "app.eirc/internal/interactor/service/lead"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *leadModel.Create) (int, interface{})
	GetByList(input *leadModel.Fields) (int, interface{})
	GetByListNoPagination(input *leadModel.Field) (int, interface{})
	GetBySingle(input *leadModel.Field) (int, interface{})
	Delete(input *leadModel.Field) (int, interface{})
	Update(trx *gorm.DB, input *leadModel.Update) (int, interface{})
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

func (m *manager) Create(trx *gorm.DB, input *leadModel.Create) (int, interface{}) {
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

func (m *manager) GetByList(input *leadModel.Fields) (int, interface{}) {
	input.IsDeleted = util.PointerBool(false)
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

func (m *manager) GetByListNoPagination(input *leadModel.Field) (int, interface{}) {
	input.IsDeleted = util.PointerBool(false)
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

func (m *manager) GetBySingle(input *leadModel.Field) (int, interface{}) {
	input.IsDeleted = util.PointerBool(false)
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

func (m *manager) Delete(input *leadModel.Field) (int, interface{}) {
	_, err := m.LeadService.GetBySingle(&leadModel.Field{
		LeadID:    input.LeadID,
		IsDeleted: util.PointerBool(false),
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

func (m *manager) Update(trx *gorm.DB, input *leadModel.Update) (int, interface{}) {
	defer trx.Rollback()

	leadBase, err := m.LeadService.GetBySingle(&leadModel.Field{
		LeadID:    input.LeadID,
		IsDeleted: util.PointerBool(false),
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
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "狀態為",
			Values: *input.Status,
		})
	}

	if input.Description != nil && *input.Description != *leadBase.Description {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "描述為",
			Values: *input.Description,
		})
	}

	if input.Source != nil && *input.Source != *leadBase.Source {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "來源為",
			Values: *input.Source,
		})
	}

	if input.Rating != nil && *input.Rating != *leadBase.Rating {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "分級為",
			Values: *input.Rating,
		})
	}

	if input.SalespersonID != nil && *input.SalespersonID != *leadBase.SalespersonID {
		salespersonBase, _ := m.UserService.GetBySingle(&userModel.Field{
			UserID:    *input.SalespersonID,
			IsDeleted: util.PointerBool(false),
		})
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "業務員為",
			Values: *salespersonBase.Name,
		})
	}

	for _, record := range records {
		_, err = m.HistoricalRecordService.WithTrx(trx).Create(&historicalRecordModel.Create{
			SourceID:   *leadBase.LeadID,
			Action:     "修改",
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
