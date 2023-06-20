package order_product

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	orderProductModel "app.eirc/internal/interactor/models/order_products"
	orderProductService "app.eirc/internal/interactor/service/order_product"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *orderProductModel.Create) (int, interface{})
	GetByList(input *orderProductModel.Fields) (int, interface{})
	GetBySingle(input *orderProductModel.Field) (int, interface{})
	Delete(input *orderProductModel.Field) (int, interface{})
	Update(input *orderProductModel.Update) (int, interface{})
}

type manager struct {
	OrderProductService orderProductService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		OrderProductService: orderProductService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *orderProductModel.Create) (int, interface{}) {
	defer trx.Rollback()

	// 計算小計
	input.SubTotal = input.UnitPrice * float64(input.Quantity)
	orderProductBase, err := m.OrderProductService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, orderProductBase.OrderProductID)
}

func (m *manager) GetByList(input *orderProductModel.Fields) (int, interface{}) {
	output := &orderProductModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, orderProductBase, err := m.OrderProductService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	orderProductByte, err := json.Marshal(orderProductBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(orderProductByte, &output.OrderProducts)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, orderProducts := range output.OrderProducts {
		orderProducts.ProductName = *orderProductBase[i].Products.Name
		orderProducts.CreatedBy = *orderProductBase[i].CreatedByUsers.Name
		orderProducts.UpdatedBy = *orderProductBase[i].UpdatedByUsers.Name
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *orderProductModel.Field) (int, interface{}) {
	orderProductBase, err := m.OrderProductService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &orderProductModel.Single{}
	orderProductByte, _ := json.Marshal(orderProductBase)
	err = json.Unmarshal(orderProductByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.ProductName = *orderProductBase.Products.Name
	output.CreatedBy = *orderProductBase.CreatedByUsers.Name
	output.UpdatedBy = *orderProductBase.UpdatedByUsers.Name

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *orderProductModel.Field) (int, interface{}) {
	_, err := m.OrderProductService.GetBySingle(&orderProductModel.Field{
		OrderProductID: input.OrderProductID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.OrderProductService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *orderProductModel.Update) (int, interface{}) {
	orderProductBase, err := m.OrderProductService.GetBySingle(&orderProductModel.Field{
		OrderProductID: input.OrderProductID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步更新小計
	unitPrice := input.UnitPrice
	quantity := input.Quantity
	if input.UnitPrice == orderProductBase.UnitPrice {
		unitPrice = orderProductBase.UnitPrice
	}
	if input.Quantity == orderProductBase.Quantity {
		quantity = orderProductBase.Quantity
	}
	input.SubTotal = *unitPrice * float64(*quantity)

	err = m.OrderProductService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, orderProductBase.OrderProductID)
}
