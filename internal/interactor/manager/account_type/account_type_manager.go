package account_type

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	accountTypeModel "app.eirc/internal/interactor/models/account_types"
	accountTypeService "app.eirc/internal/interactor/service/account_type"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *accountTypeModel.Create) interface{}
	GetByList(input *accountTypeModel.Fields) interface{}
	GetBySingle(input *accountTypeModel.Field) interface{}
	Delete(input *accountTypeModel.Field) interface{}
	Update(input *accountTypeModel.Update) interface{}
}

type manager struct {
	AccountTypeService accountTypeService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		AccountTypeService: accountTypeService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *accountTypeModel.Create) interface{} {
	defer trx.Rollback()

	accountTypeBase, err := m.AccountTypeService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, accountTypeBase.AccountTypeID)
}

func (m *manager) GetByList(input *accountTypeModel.Fields) interface{} {
	output := &accountTypeModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, accountTypeBase, err := m.AccountTypeService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	accountTypeByte, err := json.Marshal(accountTypeBase)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(accountTypeByte, &output.Industries)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *accountTypeModel.Field) interface{} {
	accountTypeBase, err := m.AccountTypeService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &accountTypeModel.Single{}
	accountTypeByte, _ := json.Marshal(accountTypeBase)
	err = json.Unmarshal(accountTypeByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *accountTypeModel.Field) interface{} {
	_, err := m.AccountTypeService.GetBySingle(&accountTypeModel.Field{
		AccountTypeID: input.AccountTypeID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.AccountTypeService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *accountTypeModel.Update) interface{} {
	accountTypeBase, err := m.AccountTypeService.GetBySingle(&accountTypeModel.Field{
		AccountTypeID: input.AccountTypeID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.AccountTypeService.Update(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, accountTypeBase.AccountTypeID)
}
