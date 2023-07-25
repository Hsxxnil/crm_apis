package event_user_attendee

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	eventUserAttendeeModel "app.eirc/internal/interactor/models/event_user_attendees"
	eventUserAttendeeService "app.eirc/internal/interactor/service/event_user_attendee"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *eventUserAttendeeModel.Create) (int, interface{})
	GetByList(input *eventUserAttendeeModel.Fields) (int, interface{})
	GetBySingle(input *eventUserAttendeeModel.Field) (int, interface{})
	Delete(input *eventUserAttendeeModel.Update) (int, interface{})
	Update(input *eventUserAttendeeModel.Update) (int, interface{})
}

type manager struct {
	EventUserAttendeeService eventUserAttendeeService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		EventUserAttendeeService: eventUserAttendeeService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *eventUserAttendeeModel.Create) (int, interface{}) {
	defer trx.Rollback()

	eventUserAttendeeBase, err := m.EventUserAttendeeService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, eventUserAttendeeBase.EventUserAttendeeID)
}

func (m *manager) GetByList(input *eventUserAttendeeModel.Fields) (int, interface{}) {
	input.IsDeleted = util.PointerBool(false)
	output := &eventUserAttendeeModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, eventUserAttendeeBase, err := m.EventUserAttendeeService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	eventUserAttendeeByte, err := json.Marshal(eventUserAttendeeBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(eventUserAttendeeByte, &output.EventUserAttendees)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *eventUserAttendeeModel.Field) (int, interface{}) {
	input.IsDeleted = util.PointerBool(false)
	eventUserAttendeeBase, err := m.EventUserAttendeeService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &eventUserAttendeeModel.Single{}
	eventUserAttendeeByte, _ := json.Marshal(eventUserAttendeeBase)
	err = json.Unmarshal(eventUserAttendeeByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *eventUserAttendeeModel.Update) (int, interface{}) {
	_, err := m.EventUserAttendeeService.GetBySingle(&eventUserAttendeeModel.Field{
		EventUserAttendeeID: input.EventUserAttendeeID,
		IsDeleted:           util.PointerBool(false),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.EventUserAttendeeService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *eventUserAttendeeModel.Update) (int, interface{}) {
	eventUserAttendeeBase, err := m.EventUserAttendeeService.GetBySingle(&eventUserAttendeeModel.Field{
		EventUserAttendeeID: input.EventUserAttendeeID,
		IsDeleted:           util.PointerBool(false),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.EventUserAttendeeService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, eventUserAttendeeBase.EventUserAttendeeID)
}
