package event_contact

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	eventContactModel "app.eirc/internal/interactor/models/event_contacts"
	eventContactService "app.eirc/internal/interactor/service/event_contact"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *eventContactModel.Create) (int, interface{})
	GetByList(input *eventContactModel.Fields) (int, interface{})
	GetBySingle(input *eventContactModel.Field) (int, interface{})
	Delete(input *eventContactModel.Update) (int, interface{})
	Update(input *eventContactModel.Update) (int, interface{})
}

type manager struct {
	EventContactService eventContactService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		EventContactService: eventContactService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *eventContactModel.Create) (int, interface{}) {
	defer trx.Rollback()

	eventContactBase, err := m.EventContactService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, eventContactBase.EventContactID)
}

func (m *manager) GetByList(input *eventContactModel.Fields) (int, interface{}) {
	input.IsDeleted = util.PointerBool(false)
	output := &eventContactModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, eventContactBase, err := m.EventContactService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	eventContactByte, err := json.Marshal(eventContactBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(eventContactByte, &output.EventContacts)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *eventContactModel.Field) (int, interface{}) {
	input.IsDeleted = util.PointerBool(false)
	eventContactBase, err := m.EventContactService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &eventContactModel.Single{}
	eventContactByte, _ := json.Marshal(eventContactBase)
	err = json.Unmarshal(eventContactByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *eventContactModel.Update) (int, interface{}) {
	_, err := m.EventContactService.GetBySingle(&eventContactModel.Field{
		EventContactID: input.EventContactID,
		IsDeleted:      util.PointerBool(false),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.EventContactService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *eventContactModel.Update) (int, interface{}) {
	eventContactBase, err := m.EventContactService.GetBySingle(&eventContactModel.Field{
		EventContactID: input.EventContactID,
		IsDeleted:      util.PointerBool(false),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.EventContactService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, eventContactBase.EventContactID)
}
