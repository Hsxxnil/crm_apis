package customer

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	customerModel "app.eirc/internal/interactor/models/customers"
	"app.eirc/internal/interactor/service/customer"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *customerModel.Create) interface{}
	GetByList(input *customerModel.Fields) interface{}
	GetBySingle(input *customerModel.Field) interface{}
	Delete(input *customerModel.Field) interface{}
	Update(input *customerModel.Update) interface{}
}

type manager struct {
	CustomerService customer.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		CustomerService: customer.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *customerModel.Create) interface{} {
	defer trx.Rollback()

	customer, err := m.CustomerService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, customer.CID)
}

func (m *manager) GetByList(input *customerModel.Fields) interface{} {
	output := &customerModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, customer, err := m.CustomerService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Total.Total = quantity
	customerByte, err := json.Marshal(customer)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(customerByte, &output.Customers)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *customerModel.Field) interface{} {
	customer, err := m.CustomerService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &customerModel.Single{}
	customerByte, _ := json.Marshal(customer)
	err = json.Unmarshal(customerByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *customerModel.Field) interface{} {
	_, err := m.CustomerService.GetBySingle(&customerModel.Field{CID: input.CID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.CustomerService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *customerModel.Update) interface{} {
	customer, err := m.CustomerService.GetBySingle(&customerModel.Field{CID: input.CID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.CustomerService.Update(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, customer.CID)
}
