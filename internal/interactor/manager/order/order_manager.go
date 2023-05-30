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
	Create(trx *gorm.DB, input *orderModel.Create) (int, interface{})
	GetByList(input *orderModel.Fields) (int, interface{})
	GetBySingle(input *orderModel.Field) (int, interface{})
	GetBySingleProducts(input *orderModel.Field) (int, interface{})
	Delete(input *orderModel.Field) (int, interface{})
	Update(input *orderModel.Update) (int, interface{})
}

type manager struct {
	OrderService orderService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		OrderService: orderService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *orderModel.Create) (int, interface{}) {
	defer trx.Rollback()

	orderBase, err := m.OrderService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, orderBase.OrderID)
}

func (m *manager) GetByList(input *orderModel.Fields) (int, interface{}) {
	output := &orderModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, orderBase, err := m.OrderService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	orderByte, err := json.Marshal(orderBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(orderByte, &output.Orders)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, orders := range output.Orders {
		orders.AccountName = *orderBase[i].Accounts.Name
		orders.ContractCode = *orderBase[i].Contracts.Code
		orders.CreatedBy = *orderBase[i].CreatedByUsers.Name
		orders.UpdatedBy = *orderBase[i].UpdatedByUsers.Name
		if *orderBase[i].Status == "啟動中" {
			orders.ActivatedBy = *orderBase[i].ActivatedByUsers.Name
			orders.ActivatedAt = orderBase[i].ActivatedAt
		} else {
			orders.ActivatedAt = nil
			orders.ActivatedBy = ""
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *orderModel.Field) (int, interface{}) {
	orderBase, err := m.OrderService.GetBySingle(input)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &orderModel.Single{}
	orderByte, _ := json.Marshal(orderBase)
	err = json.Unmarshal(orderByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.AccountName = *orderBase.Accounts.Name
	output.ContractCode = *orderBase.Contracts.Code
	output.CreatedBy = *orderBase.CreatedByUsers.Name
	output.UpdatedBy = *orderBase.UpdatedByUsers.Name
	if *orderBase.Status == "啟動中" {
		output.ActivatedBy = *orderBase.ActivatedByUsers.Name
		output.ActivatedAt = orderBase.ActivatedAt
	} else {
		output.ActivatedAt = nil
		output.ActivatedBy = ""
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingleProducts(input *orderModel.Field) (int, interface{}) {
	orderBase, err := m.OrderService.GetBySingle(input)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &orderModel.SingleProducts{}
	orderByte, _ := json.Marshal(orderBase)
	err = json.Unmarshal(orderByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.AccountName = *orderBase.Accounts.Name
	output.ContractCode = *orderBase.Contracts.Code
	output.CreatedBy = *orderBase.CreatedByUsers.Name
	output.UpdatedBy = *orderBase.UpdatedByUsers.Name
	if *orderBase.Status == "啟動中" {
		output.ActivatedBy = *orderBase.ActivatedByUsers.Name
		output.ActivatedAt = orderBase.ActivatedAt
	} else {
		output.ActivatedAt = nil
		output.ActivatedBy = ""
	}
	for i, products := range orderBase.OrderProducts {
		output.OrderProducts[i].ProductName = *products.Products.Name
		output.OrderProducts[i].ProductPrice = *products.Products.Price
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *orderModel.Field) (int, interface{}) {
	_, err := m.OrderService.GetBySingle(&orderModel.Field{
		OrderID: input.OrderID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.OrderService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *orderModel.Update) (int, interface{}) {
	orderBase, err := m.OrderService.GetBySingle(&orderModel.Field{
		OrderID: input.OrderID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	if *orderBase.Status != *input.Status {
		if *input.Status == "啟動中" {
			input.ActivatedBy = input.UpdatedBy
		}
	}

	err = m.OrderService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, orderBase.OrderID)
}
