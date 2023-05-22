package quote_product

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	quoteProductModel "app.eirc/internal/interactor/models/quote_products"
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
}

func Init(db *gorm.DB) Manager {
	return &manager{
		QuoteProductService: quoteProductService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *quoteProductModel.CreateList) (int, interface{}) {
	defer trx.Rollback()

	var output []*string
	for _, inputBody := range input.QuoteProducts {
		inputBody.SubTotal = inputBody.UnitPrice * float64(inputBody.Quantity) * inputBody.Discount / 100
		quoteProductBase, err := m.QuoteProductService.WithTrx(trx).Create(inputBody)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
		output = append(output, quoteProductBase.QuoteProductID)
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetByList(input *quoteProductModel.Fields) (int, interface{}) {
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
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
			}

			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		unitPrice := inputBody.UnitPrice
		quantity := inputBody.Quantity
		discount := inputBody.Discount
		if inputBody.UnitPrice == quoteProductBase.UnitPrice {
			unitPrice = quoteProductBase.UnitPrice
		}
		if inputBody.Quantity == quoteProductBase.Quantity {
			quantity = quoteProductBase.Quantity
		}
		if inputBody.Discount == quoteProductBase.Discount {
			discount = quoteProductBase.Discount
		}
		inputBody.SubTotal = *unitPrice * float64(*quantity) * *discount / 100

		err = m.QuoteProductService.Update(inputBody)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		output = append(output, quoteProductBase.QuoteProductID)
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}
