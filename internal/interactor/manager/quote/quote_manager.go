package quote

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	quoteModel "app.eirc/internal/interactor/models/quotes"
	quoteService "app.eirc/internal/interactor/service/quote"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *quoteModel.Create) interface{}
	GetByList(input *quoteModel.Fields) interface{}
	GetBySingle(input *quoteModel.Field) interface{}
	Delete(input *quoteModel.Field) interface{}
	Update(input *quoteModel.Update) interface{}
}

type manager struct {
	QuoteService quoteService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		QuoteService: quoteService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *quoteModel.Create) interface{} {
	defer trx.Rollback()

	quoteBase, err := m.QuoteService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, quoteBase.QuoteID)
}

func (m *manager) GetByList(input *quoteModel.Fields) interface{} {
	output := &quoteModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, quoteBase, err := m.QuoteService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	quoteByte, err := json.Marshal(quoteBase)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(quoteByte, &output.Quotes)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for _, quotesBase := range quoteBase {
		for _, quotes := range output.Quotes {
			quotes.OpportunityName = *quotesBase.Opportunities.Name
			quotes.CreatedBy = *quotesBase.CreatedByUsers.Name
			quotes.UpdatedBy = *quotesBase.UpdatedByUsers.Name
		}
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *quoteModel.Field) interface{} {
	quoteBase, err := m.QuoteService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &quoteModel.Single{}
	quoteByte, _ := json.Marshal(quoteBase)
	err = json.Unmarshal(quoteByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output.OpportunityName = *quoteBase.Opportunities.Name
	output.CreatedBy = *quoteBase.CreatedByUsers.Name
	output.UpdatedBy = *quoteBase.UpdatedByUsers.Name

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *quoteModel.Field) interface{} {
	_, err := m.QuoteService.GetBySingle(&quoteModel.Field{
		QuoteID: input.QuoteID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.QuoteService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *quoteModel.Update) interface{} {
	quoteBase, err := m.QuoteService.GetBySingle(&quoteModel.Field{
		QuoteID: input.QuoteID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.QuoteService.Update(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, quoteBase.QuoteID)
}
