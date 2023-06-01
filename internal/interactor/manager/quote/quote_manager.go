package quote

import (
	"encoding/json"
	"errors"
	"math"

	quoteModel "app.eirc/internal/interactor/models/quotes"

	"app.eirc/internal/interactor/pkg/util"

	quoteService "app.eirc/internal/interactor/service/quote"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *quoteModel.Create) (int, interface{})
	GetByList(input *quoteModel.Fields) (int, interface{})
	GetBySingle(input *quoteModel.Field) (int, interface{})
	GetBySingleProducts(input *quoteModel.Field) (int, interface{})
	Delete(input *quoteModel.Field) (int, interface{})
	Update(input *quoteModel.Update) (int, interface{})
}

type manager struct {
	QuoteService quoteService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		QuoteService: quoteService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *quoteModel.Create) (int, interface{}) {
	defer trx.Rollback()

	quoteBase, err := m.QuoteService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, quoteBase.QuoteID)
}

func (m *manager) GetByList(input *quoteModel.Fields) (int, interface{}) {
	output := &quoteModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, quoteBase, err := m.QuoteService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	quoteByte, err := json.Marshal(quoteBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(quoteByte, &output.Quotes)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, quotes := range output.Quotes {
		quotes.OpportunityName = *quoteBase[i].Opportunities.Name
		quotes.CreatedBy = *quoteBase[i].CreatedByUsers.Name
		quotes.UpdatedBy = *quoteBase[i].UpdatedByUsers.Name
		for _, products := range quoteBase[i].QuoteProducts {
			quotes.SubTotal += *products.SubTotal
			quotes.TotalPrice += *products.TotalPrice
		}
		if quotes.SubTotal != 0 {
			// 四捨五入至小數點後第二位
			quotes.Discount = math.Round((1-quotes.TotalPrice/quotes.SubTotal)*100*100) / 100
		}
		quotes.GrandTotal = quotes.TotalPrice + *quoteBase[i].ShippingAndHandling + *quoteBase[i].Tax
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *quoteModel.Field) (int, interface{}) {
	quoteBase, err := m.QuoteService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &quoteModel.Single{}
	quoteByte, _ := json.Marshal(quoteBase)
	err = json.Unmarshal(quoteByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.OpportunityName = *quoteBase.Opportunities.Name
	output.CreatedBy = *quoteBase.CreatedByUsers.Name
	output.UpdatedBy = *quoteBase.UpdatedByUsers.Name
	for _, products := range quoteBase.QuoteProducts {
		output.SubTotal += *products.SubTotal
		output.TotalPrice += *products.TotalPrice
	}
	if output.SubTotal != 0 {
		// 四捨五入至小數點後第二位
		output.Discount = math.Round((1-output.TotalPrice/output.SubTotal)*100*100) / 100
	}
	output.GrandTotal = output.TotalPrice + *quoteBase.ShippingAndHandling + *quoteBase.Tax

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingleProducts(input *quoteModel.Field) (int, interface{}) {
	quoteBase, err := m.QuoteService.GetBySingle(input)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &quoteModel.SingleProducts{}
	quoteByte, _ := json.Marshal(quoteBase)
	err = json.Unmarshal(quoteByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.OpportunityName = *quoteBase.Opportunities.Name
	output.CreatedBy = *quoteBase.CreatedByUsers.Name
	output.UpdatedBy = *quoteBase.UpdatedByUsers.Name
	for i, products := range quoteBase.QuoteProducts {
		output.QuoteProducts[i].ProductName = *products.Products.Name
		output.QuoteProducts[i].ProductPrice = *products.Products.Price
		output.SubTotal += *products.SubTotal
		output.TotalPrice += *products.TotalPrice
	}
	if output.SubTotal != 0 {
		// 四捨五入至小數點後第二位
		output.Discount = math.Round((1-output.TotalPrice/output.SubTotal)*100*100) / 100
	}
	output.GrandTotal = output.TotalPrice + *quoteBase.ShippingAndHandling + *quoteBase.Tax

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *quoteModel.Field) (int, interface{}) {
	_, err := m.QuoteService.GetBySingle(&quoteModel.Field{
		QuoteID: input.QuoteID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.QuoteService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *quoteModel.Update) (int, interface{}) {
	quoteBase, err := m.QuoteService.GetBySingle(&quoteModel.Field{
		QuoteID: input.QuoteID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.QuoteService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, quoteBase.QuoteID)
}
