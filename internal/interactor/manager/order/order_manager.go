package order

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	orderModel "app.eirc/internal/interactor/models/orders"
	orderService "app.eirc/internal/interactor/service/order"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *orderModel.Create) interface{}
	GetByList(input *orderModel.Fields) interface{}
	GetBySingle(input *orderModel.Field) interface{}
	Delete(input *orderModel.Field) interface{}
	Update(input *orderModel.Update) interface{}
}

type manager struct {
	OrderService orderService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		OrderService: orderService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *orderModel.Create) interface{} {
	defer trx.Rollback()

	orderBase, err := m.OrderService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, orderBase.OrderID)
}

func (m *manager) GetByList(input *orderModel.Fields) interface{} {
	output := &orderModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, orderBase, err := m.OrderService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	orderByte, err := json.Marshal(orderBase)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(orderByte, &output.Orders)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *orderModel.Field) interface{} {
	orderBase, err := m.OrderService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &orderModel.Single{}
	orderByte, _ := json.Marshal(orderBase)
	err = json.Unmarshal(orderByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *orderModel.Field) interface{} {
	_, err := m.OrderService.GetBySingle(&orderModel.Field{
		OrderID: input.OrderID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.OrderService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *orderModel.Update) interface{} {
	orderBase, err := m.OrderService.GetBySingle(&orderModel.Field{
		OrderID: input.OrderID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.OrderService.Update(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, orderBase.OrderID)
}
