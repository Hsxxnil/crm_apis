package contact

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	contactModel "app.eirc/internal/interactor/models/contacts"
	contactService "app.eirc/internal/interactor/service/contact"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *contactModel.Create) interface{}
	GetByList(input *contactModel.Fields) interface{}
	GetBySingle(input *contactModel.Field) interface{}
	Delete(input *contactModel.Field) interface{}
	Update(input *contactModel.Update) interface{}
}

type manager struct {
	ContactService contactService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		ContactService: contactService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *contactModel.Create) interface{} {
	defer trx.Rollback()

	contactBase, err := m.ContactService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, contactBase.ContactID)
}

func (m *manager) GetByList(input *contactModel.Fields) interface{} {
	output := &contactModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, contactBase, err := m.ContactService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	contactByte, err := json.Marshal(contactBase)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(contactByte, &output.Contacts)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *contactModel.Field) interface{} {
	contactBase, err := m.ContactService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &contactModel.Single{}
	contactByte, _ := json.Marshal(contactBase)
	err = json.Unmarshal(contactByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *contactModel.Field) interface{} {
	_, err := m.ContactService.GetBySingle(&contactModel.Field{
		ContactID: input.ContactID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.ContactService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *contactModel.Update) interface{} {
	contactBase, err := m.ContactService.GetBySingle(&contactModel.Field{
		ContactID: input.ContactID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.ContactService.Update(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, contactBase.ContactID)
}
