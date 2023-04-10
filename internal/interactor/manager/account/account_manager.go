package account

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	accountModel "app.eirc/internal/interactor/models/accounts"
	"app.eirc/internal/interactor/service/account"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *accountModel.Create) interface{}
	GetByList(input *accountModel.Fields) interface{}
	GetBySingle(input *accountModel.Field) interface{}
	Delete(input *accountModel.Update) interface{}
	Update(input *accountModel.Update) interface{}
}

type manager struct {
	AccountService account.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		AccountService: account.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *accountModel.Create) interface{} {
	defer trx.Rollback()

	accountBase, err := m.AccountService.WithTrx(trx).Create(input)
	if err != nil {
		if err.Error() == "account already exists" {
			return code.GetCodeMessage(code.BadRequest, err.Error())
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, accountBase.AccountID)
}

func (m *manager) GetByList(input *accountModel.Fields) interface{} {
	input.IsDeleted = util.PointerBool(false)
	output := &accountModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, accountBase, err := m.AccountService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	accountByte, err := json.Marshal(accountBase)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(accountByte, &output.Accounts)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *accountModel.Field) interface{} {
	input.IsDeleted = util.PointerBool(false)
	accountBase, err := m.AccountService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &accountModel.Single{}
	accountByte, _ := json.Marshal(accountBase)
	err = json.Unmarshal(accountByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *accountModel.Update) interface{} {
	_, err := m.AccountService.GetBySingle(&accountModel.Field{AccountID: input.AccountID, IsDeleted: util.PointerBool(false)})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.AccountService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *accountModel.Update) interface{} {
	accountBase, err := m.AccountService.GetBySingle(&accountModel.Field{AccountID: input.AccountID, IsDeleted: util.PointerBool(false)})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.AccountService.Update(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, accountBase.AccountID)
}
