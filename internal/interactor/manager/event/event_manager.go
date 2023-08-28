package event

import (
	"encoding/json"
	"errors"
	"time"

	eventContactModel "app.eirc/internal/interactor/models/event_contacts"
	eventUserAttendeeModel "app.eirc/internal/interactor/models/event_user_attendees"
	eventUserMainModel "app.eirc/internal/interactor/models/event_user_mains"
	"app.eirc/internal/interactor/pkg/util"
	contactService "app.eirc/internal/interactor/service/contact"
	eventContactService "app.eirc/internal/interactor/service/event_contact"
	eventUserAttendeeService "app.eirc/internal/interactor/service/event_user_attendee"
	eventUserMainService "app.eirc/internal/interactor/service/event_user_main"

	eventModel "app.eirc/internal/interactor/models/events"
	eventService "app.eirc/internal/interactor/service/event"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *eventModel.Create) (int, any)
	GetByList(input *eventModel.Fields) (int, any)
	GetBySingle(input *eventModel.Field) (int, any)
	Delete(trx *gorm.DB, input *eventModel.Update) (int, any)
	Update(trx *gorm.DB, input *eventModel.Update) (int, any)
}

type manager struct {
	EventService             eventService.Service
	ContactService           contactService.Service
	EventUserMainService     eventUserMainService.Service
	EventUserAttendeeService eventUserAttendeeService.Service
	EventContactService      eventContactService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		EventService:             eventService.Init(db),
		ContactService:           contactService.Init(db),
		EventUserMainService:     eventUserMainService.Init(db),
		EventUserAttendeeService: eventUserAttendeeService.Init(db),
		EventContactService:      eventContactService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *eventModel.Create) (int, any) {
	defer trx.Rollback()

	eventBase, err := m.EventService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步新增事件主要人員關聯
	if input.Main != nil {
		for _, mainID := range input.Main {
			_, err = m.EventUserMainService.WithTrx(trx).Create(&eventUserMainModel.Create{
				EventID:   *eventBase.EventID,
				MainID:    mainID,
				CreatedBy: *eventBase.CreatedBy,
			})
			if err != nil {
				log.Error(err)
				return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
			}
		}
	}

	// 同步新增事件參與人員關聯
	if input.Attendee != nil {
		for _, attendeeID := range input.Attendee {
			_, err = m.EventUserAttendeeService.WithTrx(trx).Create(&eventUserAttendeeModel.Create{
				EventID:    *eventBase.EventID,
				AttendeeID: attendeeID,
				CreatedBy:  *eventBase.CreatedBy,
			})
			if err != nil {
				log.Error(err)
				return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
			}
		}
	}

	// 同步新增事件聯絡人關聯
	if input.Contact != nil {
		for _, contactID := range input.Contact {
			_, err = m.EventContactService.WithTrx(trx).Create(&eventContactModel.Create{
				EventID:   *eventBase.EventID,
				ContactID: contactID,
				CreatedBy: *eventBase.CreatedBy,
			})
			if err != nil {
				log.Error(err)
				return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
			}
		}
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, eventBase.EventID)
}

func (m *manager) GetByList(input *eventModel.Fields) (int, any) {
	output := &eventModel.List{}

	// 將FilterStartDate轉為時間格式
	startDate, err := time.Parse(time.RFC3339, input.FilterStartDate)
	if err != nil {
		// 如果FilterStartDate轉換失敗，設定為空字串
		input.FilterStartDate = ""
		log.Error(err)
	} else {
		// 如果FilterStartDate轉換成功，設定篩選區間為31天
		input.FilterEndDate = startDate.AddDate(0, 0, 31)
	}

	eventBase, err := m.EventService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	eventByte, err := json.Marshal(eventBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = json.Unmarshal(eventByte, &output.Events)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, events := range output.Events {
		events.AccountName = *eventBase[i].Accounts.Name
		events.CreatedBy = *eventBase[i].CreatedByUsers.Name
		events.UpdatedBy = *eventBase[i].UpdatedByUsers.Name
		for j, mains := range eventBase[i].EventUserMains {
			events.EventUserMains[j].MainName = *mains.Mains.Name
		}
		for k, attendees := range eventBase[i].EventUserAttendees {
			events.EventUserAttendees[k].AttendeeName = *attendees.Attendees.Name
		}
		for z, contacts := range eventBase[i].EventContacts {
			events.EventContacts[z].ContactName = *contacts.Contacts.Name
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *eventModel.Field) (int, any) {
	eventBase, err := m.EventService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &eventModel.Single{}
	eventByte, _ := json.Marshal(eventBase)
	err = json.Unmarshal(eventByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.AccountName = *eventBase.Accounts.Name
	output.CreatedBy = *eventBase.CreatedByUsers.Name
	output.UpdatedBy = *eventBase.UpdatedByUsers.Name
	for i, mains := range eventBase.EventUserMains {
		output.EventUserMains[i].MainName = *mains.Mains.Name
	}
	for j, attendees := range eventBase.EventUserAttendees {
		output.EventUserAttendees[j].AttendeeName = *attendees.Attendees.Name
	}
	for k, contacts := range eventBase.EventContacts {
		output.EventContacts[k].ContactName = *contacts.Contacts.Name
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(trx *gorm.DB, input *eventModel.Update) (int, any) {
	defer trx.Rollback()

	eventBase, err := m.EventService.GetBySingle(&eventModel.Field{
		EventID: input.EventID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.EventService.WithTrx(trx).Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 將舊事件主要人員關聯資料改為刪除
	eventUserMainBase, err := m.EventUserMainService.GetByListNoPagination(&eventUserMainModel.Field{
		EventID: eventBase.EventID,
	})
	for _, main := range eventUserMainBase {
		err = m.EventUserMainService.WithTrx(trx).Delete(&eventUserMainModel.Update{
			EventUserMainID: *main.EventUserMainID,
			UpdatedBy:       input.UpdatedBy,
		})
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	// 將舊事件參與人員關聯資料改為刪除
	eventUserAttendeeBase, err := m.EventUserAttendeeService.GetByListNoPagination(&eventUserAttendeeModel.Field{
		EventID: eventBase.EventID,
	})
	for _, attendee := range eventUserAttendeeBase {
		err = m.EventUserAttendeeService.WithTrx(trx).Delete(&eventUserAttendeeModel.Update{
			EventUserAttendeeID: *attendee.EventUserAttendeeID,
			UpdatedBy:           input.UpdatedBy,
		})
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	// 將舊事件聯絡人關聯資料改為刪除
	eventContactBase, err := m.EventContactService.GetByListNoPagination(&eventContactModel.Field{
		EventID: eventBase.EventID,
	})
	for _, contact := range eventContactBase {
		err = m.EventContactService.WithTrx(trx).Delete(&eventContactModel.Update{
			EventContactID: *contact.EventContactID,
			UpdatedBy:      input.UpdatedBy,
		})
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(trx *gorm.DB, input *eventModel.Update) (int, any) {
	defer trx.Rollback()

	eventBase, err := m.EventService.GetBySingle(&eventModel.Field{
		EventID: input.EventID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 設定修改期限為開始日期前一天
	deadline := eventBase.StartDate.AddDate(0, 0, -1)
	// 若現在時間在修改期限之後回傳400
	if util.NowToUTC().After(deadline) {
		log.Info("Exceeded modification deadline. Deadline: ", deadline)
		return code.BadRequest, code.GetCodeMessage(code.BadRequest, "Exceeded modification deadline.")
	}

	// 如果輸入的開始日期和原来的不相等，再次判斷修改期限
	if input.StartDate != nil && input.StartDate != eventBase.StartDate {
		// 設定修改期限為輸入的開始日期前一天
		deadline = input.StartDate.AddDate(0, 0, -1)
		// 若現在時間在修改期限之後回傳400
		if util.NowToUTC().After(deadline) {
			log.Info("Exceeded modification deadline. Deadline: ", deadline)
			return code.BadRequest, code.GetCodeMessage(code.BadRequest, "Exceeded modification deadline.")
		}
	}

	err = m.EventService.WithTrx(trx).Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步修改事件主要人員關聯
	if input.Main != nil {
		// 將舊關聯資料改為刪除
		eventUserMainBase, err := m.EventUserMainService.GetByListNoPagination(&eventUserMainModel.Field{
			EventID: eventBase.EventID,
		})
		for _, main := range eventUserMainBase {
			err = m.EventUserMainService.WithTrx(trx).Delete(&eventUserMainModel.Update{
				EventUserMainID: *main.EventUserMainID,
				UpdatedBy:       input.UpdatedBy,
			})
			if err != nil {
				log.Error(err)
				return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
			}
		}
		// 新增事件主要人員關聯
		for _, mainID := range *input.Main {
			_, err = m.EventUserMainService.WithTrx(trx).Create(&eventUserMainModel.Create{
				EventID:   *eventBase.EventID,
				MainID:    mainID,
				CreatedBy: *eventBase.CreatedBy,
			})
			if err != nil {
				log.Error(err)
				return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
			}
		}
	}

	// 同步修改事件參與人員關聯
	if input.Attendee != nil {
		// 將舊關聯資料改為刪除
		eventUserAttendeeBase, err := m.EventUserAttendeeService.GetByListNoPagination(&eventUserAttendeeModel.Field{
			EventID: eventBase.EventID,
		})
		for _, attendee := range eventUserAttendeeBase {
			err = m.EventUserAttendeeService.WithTrx(trx).Delete(&eventUserAttendeeModel.Update{
				EventUserAttendeeID: *attendee.EventUserAttendeeID,
				UpdatedBy:           input.UpdatedBy,
			})
			if err != nil {
				log.Error(err)
				return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
			}
		}
		// 新增事件參與人員關聯
		for _, attendeeID := range *input.Attendee {
			_, err = m.EventUserAttendeeService.WithTrx(trx).Create(&eventUserAttendeeModel.Create{
				EventID:    *eventBase.EventID,
				AttendeeID: attendeeID,
				CreatedBy:  *eventBase.CreatedBy,
			})
			if err != nil {
				log.Error(err)
				return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
			}
		}
	}

	// 同步修改事件聯絡人關聯
	if input.Contact != nil {
		// 將舊關聯資料改為刪除
		eventContactBase, err := m.EventContactService.GetByListNoPagination(&eventContactModel.Field{
			EventID: eventBase.EventID,
		})
		for _, contact := range eventContactBase {
			err = m.EventContactService.WithTrx(trx).Delete(&eventContactModel.Update{
				EventContactID: *contact.EventContactID,
				UpdatedBy:      input.UpdatedBy,
			})
			if err != nil {
				log.Error(err)
				return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
			}
		}
		// 新增事件聯絡人關聯
		for _, contactID := range *input.Contact {
			_, err = m.EventContactService.WithTrx(trx).Create(&eventContactModel.Create{
				EventID:   *eventBase.EventID,
				ContactID: contactID,
				CreatedBy: *eventBase.CreatedBy,
			})
			if err != nil {
				log.Error(err)
				return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
			}
		}
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, eventBase.EventID)
}
