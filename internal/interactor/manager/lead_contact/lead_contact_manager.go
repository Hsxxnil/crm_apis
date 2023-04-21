package lead_contact

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	leadContactModel "app.eirc/internal/interactor/models/lead_contacts"
	leadContactService "app.eirc/internal/interactor/service/lead_contact"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *leadContactModel.Create) interface{}
	GetByList(input *leadContactModel.Fields) interface{}
	GetBySingle(input *leadContactModel.Field) interface{}
	Delete(input *leadContactModel.Field) interface{}
	Update(input *leadContactModel.Update) interface{}
}

type manager struct {
	LeadContactService leadContactService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		LeadContactService: leadContactService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *leadContactModel.Create) interface{} {
	defer trx.Rollback()

	leadContactBase, err := m.LeadContactService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, leadContactBase.LeadContactID)
}

func (m *manager) GetByList(input *leadContactModel.Fields) interface{} {
	output := &leadContactModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, leadContactBase, err := m.LeadContactService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	leadContactByte, err := json.Marshal(leadContactBase)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(leadContactByte, &output.LeadContacts)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *leadContactModel.Field) interface{} {
	leadContactBase, err := m.LeadContactService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &leadContactModel.Single{}
	leadContactByte, _ := json.Marshal(leadContactBase)
	err = json.Unmarshal(leadContactByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *leadContactModel.Field) interface{} {
	_, err := m.LeadContactService.GetBySingle(&leadContactModel.Field{
		LeadContactID: input.LeadContactID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.LeadContactService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *leadContactModel.Update) interface{} {
	leadContactBase, err := m.LeadContactService.GetBySingle(&leadContactModel.Field{
		LeadContactID: input.LeadContactID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.LeadContactService.Update(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, leadContactBase.LeadContactID)
}
