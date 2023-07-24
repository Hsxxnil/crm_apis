package event

import (
	contactModel "app.eirc/internal/interactor/models/contacts"
	userModel "app.eirc/internal/interactor/models/users"
	"app.eirc/internal/interactor/pkg/util"
	contactService "app.eirc/internal/interactor/service/contact"
	userService "app.eirc/internal/interactor/service/user"

	"encoding/json"
	"errors"
	"strings"
	"time"

	eventModel "app.eirc/internal/interactor/models/events"
	eventService "app.eirc/internal/interactor/service/event"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *eventModel.Create) (int, interface{})
	GetByList(input *eventModel.Fields) (int, interface{})
	GetBySingle(input *eventModel.Field) (int, interface{})
	Delete(input *eventModel.Field) (int, interface{})
	Update(input *eventModel.Update) (int, interface{})
}

type manager struct {
	EventService   eventService.Service
	UserService    userService.Service
	ContactService contactService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		EventService:   eventService.Init(db),
		UserService:    userService.Init(db),
		ContactService: contactService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *eventModel.Create) (int, interface{}) {
	defer trx.Rollback()

	input.MainID = strings.Join(input.Main, ", ")
	input.AttendeeID = strings.Join(input.Attendee, ", ")
	input.ContactID = strings.Join(input.Contact, ", ")

	eventBase, err := m.EventService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, eventBase.EventID)
}

func (m *manager) GetByList(input *eventModel.Fields) (int, interface{}) {
	output := &eventModel.List{}

	// 將FilterStartDate轉為時間格式
	startDate, err := time.Parse(time.RFC3339, input.FilterStartDate)
	if err != nil {
		// 如果FilterStartDate轉換師被，設定為空字串
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

	// TODO 暫時先取得所有使用者ID
	// 取得所有使用者ID
	users, _ := m.UserService.GetByListNoQuantity(&userModel.Field{})
	// 建立使用者ID的映射表
	userMap := make(map[string]string)
	for _, user := range users {
		userMap[*user.UserID] = *user.Name
	}

	// TODO 暫時先取得所有聯絡人ID
	// 取得所有聯絡人ID
	contacts, err := m.ContactService.GetByListNoQuantity(&contactModel.Field{})
	// 建立聯絡人ID的映射表
	contactMap := make(map[string]string)
	for _, contact := range contacts {
		contactMap[*contact.ContactID] = *contact.Name
	}

	for i, events := range output.Events {
		events.AccountName = *eventBase[i].Accounts.Name
		events.CreatedBy = *eventBase[i].CreatedByUsers.Name
		events.UpdatedBy = *eventBase[i].UpdatedByUsers.Name
		// 將MainID拆分並賦值到events.Main
		mainIDs := strings.Split(*eventBase[i].MainID, ",")
		for _, main := range mainIDs {
			mainID := strings.TrimSpace(main)
			// 尋找MainID對應的中文名稱
			events.Main = append(events.Main, &eventModel.Main{
				MainID:   mainID,
				MainName: userMap[mainID],
			})
		}

		// 將AttendeeID拆分並賦值到events.Attendees
		attendeeIDs := strings.Split(*eventBase[i].AttendeeID, ",")
		for _, attendee := range attendeeIDs {
			attendeeID := strings.TrimSpace(attendee)
			// 尋找AttendeeID對應的中文名稱
			events.Attendees = append(events.Attendees, &eventModel.Attendees{
				AttendeeID:   attendeeID,
				AttendeeName: userMap[attendeeID],
			})
		}

		// 將ContactID拆分並賦值到events.Contacts
		contactIDs := strings.Split(*eventBase[i].ContactID, ",")
		for _, contact := range contactIDs {
			contactID := strings.TrimSpace(contact)
			// 尋找ContactID對應的中文名稱
			events.Contacts = append(events.Contacts, &eventModel.Contacts{
				ContactID:   contactID,
				ContactName: contactMap[contactID],
			})
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *eventModel.Field) (int, interface{}) {
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

	// TODO 暫時先取得所有使用者ID
	// 取得所有使用者ID
	users, _ := m.UserService.GetByListNoQuantity(&userModel.Field{})
	// 建立使用者ID的映射表
	userMap := make(map[string]string)
	for _, user := range users {
		userMap[*user.UserID] = *user.Name
	}

	// TODO 暫時先取得所有聯絡人ID
	// 取得所有聯絡人ID
	contacts, err := m.ContactService.GetByListNoQuantity(&contactModel.Field{})
	// 建立聯絡人ID的映射表
	contactMap := make(map[string]string)
	for _, contact := range contacts {
		contactMap[*contact.ContactID] = *contact.Name
	}

	output.AccountName = *eventBase.Accounts.Name
	output.CreatedBy = *eventBase.CreatedByUsers.Name
	output.UpdatedBy = *eventBase.UpdatedByUsers.Name
	// 將MainID拆分並賦值到output.Main
	mainIDs := strings.Split(*eventBase.MainID, ",")
	for _, main := range mainIDs {
		mainID := strings.TrimSpace(main)
		// 尋找MainID對應的中文名稱
		output.Main = append(output.Main, &eventModel.Main{
			MainID:   mainID,
			MainName: userMap[mainID],
		})
	}

	// 將AttendeeID拆分並賦值到output.Attendees
	attendeeIDs := strings.Split(*eventBase.AttendeeID, ",")
	for _, attendee := range attendeeIDs {
		attendeeID := strings.TrimSpace(attendee)
		// 尋找AttendeeID對應的中文名稱
		output.Attendees = append(output.Attendees, &eventModel.Attendees{
			AttendeeID:   attendeeID,
			AttendeeName: userMap[attendeeID],
		})
	}

	// 將ContactID拆分並賦值到output.Contacts
	contactIDs := strings.Split(*eventBase.ContactID, ",")
	for _, contact := range contactIDs {
		contactID := strings.TrimSpace(contact)
		// 尋找ContactID對應的中文名稱
		output.Contacts = append(output.Contacts, &eventModel.Contacts{
			ContactID:   contactID,
			ContactName: contactMap[contactID],
		})
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *eventModel.Field) (int, interface{}) {
	_, err := m.EventService.GetBySingle(&eventModel.Field{
		EventID: input.EventID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.EventService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *eventModel.Update) (int, interface{}) {
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

	mainID := util.PointerString(strings.Join(*input.Main, ", "))
	attendeeID := util.PointerString(strings.Join(*input.Attendee, ", "))
	contactID := util.PointerString(strings.Join(*input.Contact, ", "))

	if input.Main != nil && mainID != eventBase.MainID {
		input.MainID = mainID
	}
	if input.Attendee != nil && attendeeID != eventBase.AttendeeID {
		input.AttendeeID = attendeeID
	}
	if input.Contact != nil && contactID != eventBase.ContactID {
		input.ContactID = contactID
	}

	err = m.EventService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, eventBase.EventID)
}
