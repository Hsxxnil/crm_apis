package lead

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	leadModel "app.eirc/internal/interactor/models/leads"
	leadService "app.eirc/internal/interactor/service/lead"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *leadModel.Create) interface{}
	GetByList(input *leadModel.Fields) interface{}
	GetBySingle(input *leadModel.Field) interface{}
	Delete(input *leadModel.Field) interface{}
	Update(input *leadModel.Update) interface{}
}

type manager struct {
	LeadService leadService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		LeadService: leadService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *leadModel.Create) interface{} {
	defer trx.Rollback()

	leadBase, err := m.LeadService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, leadBase.LeadID)
}

func (m *manager) GetByList(input *leadModel.Fields) interface{} {
	output := &leadModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, leadBase, err := m.LeadService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	leadByte, err := json.Marshal(leadBase)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(leadByte, &output.Leads)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, leads := range output.Leads {
		leads.AccountName = *leadBase[i].Accounts.Name
		leads.CreatedBy = *leadBase[i].CreatedByUsers.Name
		leads.UpdatedBy = *leadBase[i].UpdatedByUsers.Name
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *leadModel.Field) interface{} {
	leadBase, err := m.LeadService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &leadModel.Single{}
	leadByte, _ := json.Marshal(leadBase)
	err = json.Unmarshal(leadByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output.AccountName = *leadBase.Accounts.Name
	output.CreatedBy = *leadBase.CreatedByUsers.Name
	output.UpdatedBy = *leadBase.UpdatedByUsers.Name

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *leadModel.Field) interface{} {
	_, err := m.LeadService.GetBySingle(&leadModel.Field{
		LeadID: input.LeadID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.LeadService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *leadModel.Update) interface{} {
	leadBase, err := m.LeadService.GetBySingle(&leadModel.Field{
		LeadID: input.LeadID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.LeadService.Update(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, leadBase.LeadID)
}
