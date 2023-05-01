package order_product

import (
	"encoding/json"
	"errors"

	"github.com/shopspring/decimal"

	"app.eirc/internal/interactor/pkg/util"

	orderProductModel "app.eirc/internal/interactor/models/order_products"
	orderProductService "app.eirc/internal/interactor/service/order_product"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *orderProductModel.Create) interface{}
	GetByList(input *orderProductModel.Fields) interface{}
	GetBySingle(input *orderProductModel.Field) interface{}
	Delete(input *orderProductModel.Field) interface{}
	Update(input *orderProductModel.Update) interface{}
}

type manager struct {
	OrderProductService orderProductService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		OrderProductService: orderProductService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *orderProductModel.Create) interface{} {
	defer trx.Rollback()

	input.SubTotal = input.UnitPrice.Mul(decimal.NewFromInt(int64(input.Quantity)))
	orderProductBase, err := m.OrderProductService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, orderProductBase.OrderProductID)
}

func (m *manager) GetByList(input *orderProductModel.Fields) interface{} {
	output := &orderProductModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, orderProductBase, err := m.OrderProductService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	orderProductByte, err := json.Marshal(orderProductBase)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(orderProductByte, &output.OrderProducts)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, orderProducts := range output.OrderProducts {
		orderProducts.ProductName = *orderProductBase[i].Products.Name
		orderProducts.CreatedBy = *orderProductBase[i].CreatedByUsers.Name
		orderProducts.UpdatedBy = *orderProductBase[i].UpdatedByUsers.Name
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *orderProductModel.Field) interface{} {
	orderProductBase, err := m.OrderProductService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &orderProductModel.Single{}
	orderProductByte, _ := json.Marshal(orderProductBase)
	err = json.Unmarshal(orderProductByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output.ProductName = *orderProductBase.Products.Name
	output.CreatedBy = *orderProductBase.CreatedByUsers.Name
	output.UpdatedBy = *orderProductBase.UpdatedByUsers.Name

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *orderProductModel.Field) interface{} {
	_, err := m.OrderProductService.GetBySingle(&orderProductModel.Field{
		OrderProductID: input.OrderProductID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.OrderProductService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *orderProductModel.Update) interface{} {
	orderProductBase, err := m.OrderProductService.GetBySingle(&orderProductModel.Field{
		OrderProductID: input.OrderProductID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	if input.UnitPrice != nil {
		input.SubTotal = input.UnitPrice.Mul(decimal.NewFromInt(int64(*orderProductBase.Quantity)))
	} else if input.Quantity != nil {
		input.SubTotal = orderProductBase.UnitPrice.Mul(decimal.NewFromInt(int64(*input.Quantity)))
	} else {
		input.SubTotal = orderProductBase.UnitPrice.Mul(decimal.NewFromInt(int64(*orderProductBase.Quantity)))
	}

	err = m.OrderProductService.Update(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, orderProductBase.OrderProductID)
}
