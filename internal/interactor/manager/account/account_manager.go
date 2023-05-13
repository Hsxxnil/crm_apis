package account

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	accountModel "app.eirc/internal/interactor/models/accounts"
	accountService "app.eirc/internal/interactor/service/account"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *accountModel.Create) (int, interface{})
	GetByList(input *accountModel.Fields) (int, interface{})
	GetBySingle(input *accountModel.Field) (int, interface{})
	GetBySingleContacts(input *accountModel.Field) (int, interface{})
	Delete(input *accountModel.Field) (int, interface{})
	Update(input *accountModel.Update) (int, interface{})
}

type manager struct {
	AccountService accountService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		AccountService: accountService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *accountModel.Create) (int, interface{}) {
	defer trx.Rollback()

	accountBase, err := m.AccountService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, accountBase.AccountID)
}

func (m *manager) GetByList(input *accountModel.Fields) (int, interface{}) {
	output := &accountModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, accountBase, err := m.AccountService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	accountByte, err := json.Marshal(accountBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(accountByte, &output.Accounts)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, accounts := range output.Accounts {
		accounts.IndustryName = *accountBase[i].Industries.Name
		accounts.CreatedBy = *accountBase[i].CreatedByUsers.Name
		accounts.UpdatedBy = *accountBase[i].UpdatedByUsers.Name
		accounts.SalespersonName = *accountBase[i].Salespeople.Name
		if parentAccountsBase, err := m.AccountService.GetBySingle(&accountModel.Field{
			AccountID: accounts.ParentAccountID,
		}); err != nil {
			accounts.ParentAccountName = ""
		} else {
			accounts.ParentAccountName = *parentAccountsBase.Name
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *accountModel.Field) (int, interface{}) {
	accountBase, err := m.AccountService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &accountModel.Single{}
	accountByte, _ := json.Marshal(accountBase)
	err = json.Unmarshal(accountByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.IndustryName = *accountBase.Industries.Name
	output.CreatedBy = *accountBase.CreatedByUsers.Name
	output.UpdatedBy = *accountBase.UpdatedByUsers.Name
	output.SalespersonName = *accountBase.Salespeople.Name
	if parentAccountsBase, err := m.AccountService.GetBySingle(&accountModel.Field{
		AccountID: *accountBase.ParentAccountID,
	}); err != nil {
		output.ParentAccountName = ""
	} else {
		output.ParentAccountName = *parentAccountsBase.Name
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingleContacts(input *accountModel.Field) (int, interface{}) {
	accountBase, err := m.AccountService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &accountModel.SingleContacts{}
	accountByte, _ := json.Marshal(accountBase)
	err = json.Unmarshal(accountByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.IndustryName = *accountBase.Industries.Name
	output.CreatedBy = *accountBase.CreatedByUsers.Name
	output.UpdatedBy = *accountBase.UpdatedByUsers.Name
	output.SalespersonName = *accountBase.Salespeople.Name
	if parentAccountsBase, err := m.AccountService.GetBySingle(&accountModel.Field{
		AccountID: *accountBase.ParentAccountID,
	}); err != nil {
		output.ParentAccountName = ""
	} else {
		output.ParentAccountName = *parentAccountsBase.Name
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *accountModel.Field) (int, interface{}) {
	_, err := m.AccountService.GetBySingle(&accountModel.Field{
		AccountID: input.AccountID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.AccountService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *accountModel.Update) (int, interface{}) {
	accountBase, err := m.AccountService.GetBySingle(&accountModel.Field{
		AccountID: input.AccountID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.AccountService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, accountBase.AccountID)
}
