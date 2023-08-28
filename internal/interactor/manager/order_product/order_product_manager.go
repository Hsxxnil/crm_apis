package order_product

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	orderModel "app.eirc/internal/interactor/models/orders"
	orderService "app.eirc/internal/interactor/service/order"

	"app.eirc/internal/interactor/pkg/util"

	orderProductModel "app.eirc/internal/interactor/models/order_products"
	orderProductService "app.eirc/internal/interactor/service/order_product"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *orderProductModel.CreateList) (int, any)
	GetByList(input *orderProductModel.Fields) (int, any)
	GetBySingle(input *orderProductModel.Field) (int, any)
	Delete(input *orderProductModel.UpdateList) (int, any)
	Update(input *orderProductModel.UpdateList) (int, any)
}

type manager struct {
	OrderProductService orderProductService.Service
	OrderService        orderService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		OrderProductService: orderProductService.Init(db),
		OrderService:        orderService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *orderProductModel.CreateList) (int, any) {
	defer trx.Rollback()

	var output []*string
	number := 0
	for i, inputBody := range input.OrderProducts {
		// 計算小計
		inputBody.SubTotal = inputBody.UnitPrice * float64(inputBody.Quantity)
		// 取得訂單號碼
		orderBase, err := m.OrderService.GetBySingle(&orderModel.Field{
			OrderID: inputBody.OrderID,
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
			}
		}
		orderCode := *orderBase.Code
		// 判斷是否已存在同訂單的訂單產品
		quantity, _ := m.OrderProductService.GetByQuantity(&orderProductModel.Field{
			OrderID: util.PointerString(inputBody.OrderID),
		})
		if quantity != 0 {
			// 陣列中第一筆單號數字等於同訂單的最後一筆訂單產品單號數字+1
			if i == 0 {
				// 取得同訂單的最後一筆單號
				orderProductBase, _ := m.OrderProductService.GetByLastCode(&orderProductModel.Field{
					OrderID: util.PointerString(inputBody.OrderID),
				})
				// 將最後一筆單號的數字部分取出
				codeParts := strings.Split(*orderProductBase.Code, "-")
				numericPart := codeParts[1]
				number, _ = strconv.Atoi(numericPart)
			}
		}
		inputBody.Code = fmt.Sprintf("%s-%d", orderCode, number+1)
		orderProductBase, err := m.OrderProductService.WithTrx(trx).Create(inputBody)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
		output = append(output, orderProductBase.OrderProductID)
		// 陣列中第二筆後單號數字等於前次迴圈的單號數字+1
		number++
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetByList(input *orderProductModel.Fields) (int, any) {
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

func (m *manager) GetBySingle(input *orderProductModel.Field) (int, any) {
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

func (m *manager) Delete(input *orderProductModel.UpdateList) (int, any) {
	for _, inputBody := range input.OrderProducts {
		_, err := m.OrderProductService.GetBySingle(&orderProductModel.Field{
			OrderProductID: inputBody.OrderProductID,
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
			}

			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		err = m.OrderProductService.Delete(inputBody)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *orderProductModel.UpdateList) (int, any) {
	var output []*string
	for _, inputBody := range input.OrderProducts {
		orderProductBase, err := m.OrderProductService.GetBySingle(&orderProductModel.Field{
			OrderProductID: inputBody.OrderProductID,
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
			}

			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		// 同步更新小計
		unitPrice := inputBody.UnitPrice
		quantity := inputBody.Quantity
		if inputBody.UnitPrice == orderProductBase.UnitPrice {
			unitPrice = orderProductBase.UnitPrice
		}
		if inputBody.Quantity == orderProductBase.Quantity {
			quantity = orderProductBase.Quantity
		}
		inputBody.SubTotal = *unitPrice * float64(*quantity)

		err = m.OrderProductService.Update(inputBody)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		output = append(output, orderProductBase.OrderProductID)
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}
