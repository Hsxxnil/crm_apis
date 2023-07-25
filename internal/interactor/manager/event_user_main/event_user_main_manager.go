package event_user_main

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	eventUserMainModel "app.eirc/internal/interactor/models/event_user_mains"
	eventUserMainService "app.eirc/internal/interactor/service/event_user_main"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *eventUserMainModel.Create) (int, interface{})
	GetByList(input *eventUserMainModel.Fields) (int, interface{})
	GetBySingle(input *eventUserMainModel.Field) (int, interface{})
	Delete(input *eventUserMainModel.Update) (int, interface{})
	Update(input *eventUserMainModel.Update) (int, interface{})
}

type manager struct {
	EventUserMainService eventUserMainService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		EventUserMainService: eventUserMainService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *eventUserMainModel.Create) (int, interface{}) {
	defer trx.Rollback()

	eventUserMainBase, err := m.EventUserMainService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, eventUserMainBase.EventUserMainID)
}

func (m *manager) GetByList(input *eventUserMainModel.Fields) (int, interface{}) {
	input.IsDeleted = util.PointerBool(false)
	output := &eventUserMainModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, eventUserMainBase, err := m.EventUserMainService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	eventUserMainByte, err := json.Marshal(eventUserMainBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(eventUserMainByte, &output.EventUserMains)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *eventUserMainModel.Field) (int, interface{}) {
	input.IsDeleted = util.PointerBool(false)
	eventUserMainBase, err := m.EventUserMainService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &eventUserMainModel.Single{}
	eventUserMainByte, _ := json.Marshal(eventUserMainBase)
	err = json.Unmarshal(eventUserMainByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *eventUserMainModel.Update) (int, interface{}) {
	_, err := m.EventUserMainService.GetBySingle(&eventUserMainModel.Field{
		EventUserMainID: input.EventUserMainID,
		IsDeleted:       util.PointerBool(false),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.EventUserMainService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *eventUserMainModel.Update) (int, interface{}) {
	eventUserMainBase, err := m.EventUserMainService.GetBySingle(&eventUserMainModel.Field{
		EventUserMainID: input.EventUserMainID,
		IsDeleted:       util.PointerBool(false),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.EventUserMainService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, eventUserMainBase.EventUserMainID)
}
