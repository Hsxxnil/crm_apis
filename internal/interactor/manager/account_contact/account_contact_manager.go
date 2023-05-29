package account_contact

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	accountContactModel "app.eirc/internal/interactor/models/account_contacts"
	accountContactService "app.eirc/internal/interactor/service/account_contact"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *accountContactModel.Create) (int, interface{})
	GetByList(input *accountContactModel.Fields) (int, interface{})
	GetBySingle(input *accountContactModel.Field) (int, interface{})
	Delete(input *accountContactModel.Field) (int, interface{})
	Update(input *accountContactModel.Update) (int, interface{})
}

type manager struct {
	AccountContactService accountContactService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		AccountContactService: accountContactService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *accountContactModel.Create) (int, interface{}) {
	defer trx.Rollback()

	accountContactBase, err := m.AccountContactService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, accountContactBase.AccountContactID)
}

func (m *manager) GetByList(input *accountContactModel.Fields) (int, interface{}) {
	output := &accountContactModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, accountContactBase, err := m.AccountContactService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	accountContactByte, err := json.Marshal(accountContactBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(accountContactByte, &output.AccountContacts)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, AccountContacts := range output.AccountContacts {
		AccountContacts.CreatedBy = *accountContactBase[i].CreatedByUsers.Name
		AccountContacts.UpdatedBy = *accountContactBase[i].UpdatedByUsers.Name
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *accountContactModel.Field) (int, interface{}) {
	accountContactBase, err := m.AccountContactService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &accountContactModel.Single{}
	accountContactByte, _ := json.Marshal(accountContactBase)
	err = json.Unmarshal(accountContactByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.CreatedBy = *accountContactBase.CreatedByUsers.Name
	output.UpdatedBy = *accountContactBase.UpdatedByUsers.Name

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *accountContactModel.Field) (int, interface{}) {
	_, err := m.AccountContactService.GetBySingle(&accountContactModel.Field{
		AccountContactID: input.AccountContactID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.AccountContactService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *accountContactModel.Update) (int, interface{}) {
	accountContactBase, err := m.AccountContactService.GetBySingle(&accountContactModel.Field{
		AccountContactID: input.AccountContactID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.AccountContactService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, accountContactBase.AccountContactID)
}
