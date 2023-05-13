package contact

import (
	"encoding/json"
	"errors"

	accountModel "app.eirc/internal/interactor/models/accounts"
	accountService "app.eirc/internal/interactor/service/account"

	"app.eirc/internal/interactor/pkg/util"

	contactModel "app.eirc/internal/interactor/models/contacts"
	contactService "app.eirc/internal/interactor/service/contact"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *contactModel.Create) (int, interface{})
	GetByList(input *contactModel.Fields) (int, interface{})
	GetBySingle(input *contactModel.Field) (int, interface{})
	Delete(input *contactModel.Field) (int, interface{})
	Update(input *contactModel.Update) (int, interface{})
}

type manager struct {
	ContactService contactService.Service
	AccountService accountService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		ContactService: contactService.Init(db),
		AccountService: accountService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *contactModel.Create) (int, interface{}) {
	defer trx.Rollback()

	contactBase, err := m.ContactService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, contactBase.ContactID)
}

func (m *manager) GetByList(input *contactModel.Fields) (int, interface{}) {
	output := &contactModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, contactBase, err := m.ContactService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	contactByte, err := json.Marshal(contactBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(contactByte, &output.Contacts)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, contacts := range output.Contacts {
		contacts.CreatedBy = *contactBase[i].CreatedByUsers.Name
		contacts.UpdatedBy = *contactBase[i].UpdatedByUsers.Name
		contacts.SalespersonName = *contactBase[i].Salespeople.Name
		if supervisorBase, err := m.ContactService.GetBySingle(&contactModel.Field{
			ContactID: contacts.SupervisorID,
		}); err != nil {
			contacts.SupervisorName = ""
		} else {
			contacts.SupervisorName = *supervisorBase.Name
		}
		if accountBase, err := m.AccountService.GetBySingle(&accountModel.Field{
			AccountID: contacts.AccountID,
		}); err != nil {
			contacts.AccountName = ""
		} else {
			contacts.AccountName = *accountBase.Name
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *contactModel.Field) (int, interface{}) {
	contactBase, err := m.ContactService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &contactModel.Single{}
	contactByte, _ := json.Marshal(contactBase)
	err = json.Unmarshal(contactByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.CreatedBy = *contactBase.CreatedByUsers.Name
	output.UpdatedBy = *contactBase.UpdatedByUsers.Name
	output.SalespersonName = *contactBase.Salespeople.Name
	if supervisorBase, err := m.ContactService.GetBySingle(&contactModel.Field{
		ContactID: *contactBase.SupervisorID,
	}); err != nil {
		output.SupervisorName = ""
	} else {
		output.SupervisorName = *supervisorBase.Name
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *contactModel.Field) (int, interface{}) {
	_, err := m.ContactService.GetBySingle(&contactModel.Field{
		ContactID: input.ContactID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.ContactService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *contactModel.Update) (int, interface{}) {
	contactBase, err := m.ContactService.GetBySingle(&contactModel.Field{
		ContactID: input.ContactID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.ContactService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, contactBase.ContactID)
}
