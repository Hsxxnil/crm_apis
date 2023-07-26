package quote_product

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"app.eirc/internal/interactor/pkg/util"

	quoteProductModel "app.eirc/internal/interactor/models/quote_products"
	quoteModel "app.eirc/internal/interactor/models/quotes"
	quoteService "app.eirc/internal/interactor/service/quote"
	quoteProductService "app.eirc/internal/interactor/service/quote_product"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *quoteProductModel.CreateList) (int, interface{})
	GetByList(input *quoteProductModel.Fields) (int, interface{})
	GetBySingle(input *quoteProductModel.Field) (int, interface{})
	Delete(input *quoteProductModel.UpdateList) (int, interface{})
	Update(input *quoteProductModel.UpdateList) (int, interface{})
}

type manager struct {
	QuoteProductService quoteProductService.Service
	QuoteService        quoteService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		QuoteProductService: quoteProductService.Init(db),
		QuoteService:        quoteService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *quoteProductModel.CreateList) (int, interface{}) {
	defer trx.Rollback()

	var output []*string
	number := 0
	for i, inputBody := range input.QuoteProducts {
		inputBody.SubTotal = inputBody.UnitPrice * float64(inputBody.Quantity)
		inputBody.TotalPrice = inputBody.SubTotal * (1 - inputBody.Discount/100)
		// 取得報價號碼
		quoteBase, err := m.QuoteService.GetBySingle(&quoteModel.Field{
			QuoteID:   inputBody.QuoteID,
			IsDeleted: util.PointerBool(false),
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
			}
		}
		quoteCode := *quoteBase.Code
		// 判斷是否已存在同報價的報價產品
		quantity, _ := m.QuoteProductService.GetByQuantity(&quoteProductModel.Field{
			QuoteID:   util.PointerString(inputBody.QuoteID),
			IsDeleted: util.PointerBool(false),
		})
		if quantity != 0 {
			// 陣列中第一筆單號數字等於同報價的最後一筆報價產品單號數字+1
			if i == 0 {
				// 取得同報價的最後一筆單號
				quoteProductBase, _ := m.QuoteProductService.GetByLastCode(&quoteProductModel.Field{
					QuoteID:   util.PointerString(inputBody.QuoteID),
					IsDeleted: util.PointerBool(false),
				})
				// 將最後一筆單號的數字部分取出
				codeParts := strings.Split(*quoteProductBase.Code, "-")
				numericPart := codeParts[1]
				number, _ = strconv.Atoi(numericPart)
			}
		}
		inputBody.Code = fmt.Sprintf("%s-%d", quoteCode, number+1)
		quoteProductBase, err := m.QuoteProductService.WithTrx(trx).Create(inputBody)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
		output = append(output, quoteProductBase.QuoteProductID)
		// 陣列中第二筆後單號數字等於前次迴圈的單號數字+1
		number++
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetByList(input *quoteProductModel.Fields) (int, interface{}) {
	input.IsDeleted = util.PointerBool(false)
	output := &quoteProductModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, quoteProductBase, err := m.QuoteProductService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	quoteProductByte, err := json.Marshal(quoteProductBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(quoteProductByte, &output.QuoteProducts)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, quoteProducts := range output.QuoteProducts {
		quoteProducts.ProductName = *quoteProductBase[i].Products.Name
		quoteProducts.ProductPrice = *quoteProductBase[i].Products.Price
		quoteProducts.CreatedBy = *quoteProductBase[i].CreatedByUsers.Name
		quoteProducts.UpdatedBy = *quoteProductBase[i].UpdatedByUsers.Name
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *quoteProductModel.Field) (int, interface{}) {
	input.IsDeleted = util.PointerBool(false)
	quoteProductBase, err := m.QuoteProductService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &quoteProductModel.Single{}
	quoteProductByte, _ := json.Marshal(quoteProductBase)
	err = json.Unmarshal(quoteProductByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.ProductName = *quoteProductBase.Products.Name
	output.ProductPrice = *quoteProductBase.Products.Price
	output.CreatedBy = *quoteProductBase.CreatedByUsers.Name
	output.UpdatedBy = *quoteProductBase.UpdatedByUsers.Name

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *quoteProductModel.UpdateList) (int, interface{}) {
	for _, inputBody := range input.QuoteProducts {
		_, err := m.QuoteProductService.GetBySingle(&quoteProductModel.Field{
			QuoteProductID: inputBody.QuoteProductID,
			IsDeleted:      util.PointerBool(false),
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
			}

			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		err = m.QuoteProductService.Delete(inputBody)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *quoteProductModel.UpdateList) (int, interface{}) {
	var output []*string
	for _, inputBody := range input.QuoteProducts {
		quoteProductBase, err := m.QuoteProductService.GetBySingle(&quoteProductModel.Field{
			QuoteProductID: inputBody.QuoteProductID,
			IsDeleted:      util.PointerBool(false),
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
			}

			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		// 同步更新小計、總價
		unitPrice := quoteProductBase.UnitPrice
		quantity := quoteProductBase.Quantity
		discount := quoteProductBase.Discount
		if inputBody.UnitPrice != nil && inputBody.UnitPrice != quoteProductBase.UnitPrice {
			unitPrice = inputBody.UnitPrice
		}
		if inputBody.Quantity != nil && inputBody.Quantity != quoteProductBase.Quantity {
			quantity = inputBody.Quantity
		}
		if inputBody.Discount != nil && inputBody.Discount != quoteProductBase.Discount {
			discount = inputBody.Discount
		}
		inputBody.SubTotal = *unitPrice * float64(*quantity)
		inputBody.TotalPrice = inputBody.SubTotal * (1 - *discount/100)

		err = m.QuoteProductService.Update(inputBody)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		output = append(output, quoteProductBase.QuoteProductID)
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}
