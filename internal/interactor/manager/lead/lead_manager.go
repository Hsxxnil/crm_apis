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
	GetBySingle(input *leadModel.Field) (int, interface{})
	Delete(input *leadModel.Field) (int, interface{})
	Update(input *leadModel.Update) (int, interface{})
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
		Content:    sourceType,
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

func (m *manager) GetBySingle(input *leadModel.Field) (int, interface{}) {
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

func (m *manager) Update(input *leadModel.Update) (int, interface{}) {
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

	err = m.LeadService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步新增線索歷程記錄
	var records []historicalRecordModel.AddHistoricalRecord

	if *input.Status != *leadBase.Status {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "狀態",
			Values: "為" + *input.Status,
		})
	}

	if *input.Description != *leadBase.Description {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "描述",
			Values: "為" + *input.Description,
		})
	}

	if *input.Source != *leadBase.Source {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "來源",
			Values: "為" + *input.Source,
		})
	}

	if *input.Rating != *leadBase.Rating {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "分級",
			Values: "為" + *input.Rating,
		})
	}

	if *input.SalespersonID != *leadBase.SalespersonID {
		salespersonBase, _ := m.UserService.GetBySingle(&userModel.Field{
			UserID: *input.SalespersonID,
		})
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "業務員",
			Values: "為" + *salespersonBase.Name,
		})
	}

	for _, record := range records {
		_, err = m.HistoricalRecordService.Create(&historicalRecordModel.Create{
			SourceID:   *leadBase.LeadID,
			Action:     "修改",
			Content:    sourceType + record.Fields + record.Values,
			ModifiedBy: *input.UpdatedBy,
		})
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, leadBase.LeadID)
}
