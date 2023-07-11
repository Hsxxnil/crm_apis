package quote

import (
	"encoding/json"
	"errors"
	"math"
	"strconv"

	accountModel "app.eirc/internal/interactor/models/accounts"

	historicalRecordModel "app.eirc/internal/interactor/models/historical_records"
	opportunityModel "app.eirc/internal/interactor/models/opportunities"
	quoteModel "app.eirc/internal/interactor/models/quotes"
	accountService "app.eirc/internal/interactor/service/account"
	historicalRecordService "app.eirc/internal/interactor/service/historical_record"
	opportunityService "app.eirc/internal/interactor/service/opportunity"

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
	Update(trx *gorm.DB, input *quoteModel.Update) (int, interface{})
}

type manager struct {
	QuoteService            quoteService.Service
	HistoricalRecordService historicalRecordService.Service
	OpportunityService      opportunityService.Service
	AccountService          accountService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		QuoteService:            quoteService.Init(db),
		HistoricalRecordService: historicalRecordService.Init(db),
		OpportunityService:      opportunityService.Init(db),
		AccountService:          accountService.Init(db),
	}
}

const sourceType = "報價"

func (m *manager) Create(trx *gorm.DB, input *quoteModel.Create) (int, interface{}) {
	defer trx.Rollback()

	// 同步商機的account_id
	opportunityBase, _ := m.OpportunityService.GetBySingle(&opportunityModel.Field{
		OpportunityID: input.OpportunityID,
	})
	input.AccountID = *opportunityBase.AccountID

	quoteBase, err := m.QuoteService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步新增報價歷程記錄
	_, err = m.HistoricalRecordService.WithTrx(trx).Create(&historicalRecordModel.Create{
		SourceID:   *quoteBase.QuoteID,
		Action:     "建立",
		Content:    sourceType,
		ModifiedBy: *quoteBase.CreatedBy,
	})
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
			// 計算報價小計
			quotes.SubTotal += *products.SubTotal
			// 計算報價總價
			quotes.TotalPrice += *products.TotalPrice
		}
		// 計算報價折扣
		if quotes.SubTotal != 0 {
			// 四捨五入至小數點後第二位
			quotes.Discount = math.Round((1-quotes.TotalPrice/quotes.SubTotal)*100*100) / 100
		}
		// 計算報價總計
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
		output.QuoteProducts[i].CreatedBy = *products.Products.CreatedByUsers.Name
		output.QuoteProducts[i].UpdatedBy = *products.Products.UpdatedByUsers.Name
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

func (m *manager) Update(trx *gorm.DB, input *quoteModel.Update) (int, interface{}) {
	defer trx.Rollback()

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

	// 同步更新商機的account_id至該報價
	if input.OpportunityID != nil && *input.OpportunityID != *quoteBase.OpportunityID {
		opportunityBase, _ := m.OpportunityService.GetBySingle(&opportunityModel.Field{
			OpportunityID: *input.OpportunityID,
		})
		input.AccountID = opportunityBase.AccountID
	}

	err = m.QuoteService.WithTrx(trx).Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步新增商機歷程記錄
	var records []historicalRecordModel.AddHistoricalRecord
	action := "修改"

	if *input.Name != *quoteBase.Name {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "名稱",
			Values: "為" + *input.Name,
		})
	}

	if *input.Status != *quoteBase.Status {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "狀態",
			Values: "為" + *input.Status,
		})
	}

	if *input.IsSyncing != *quoteBase.IsSyncing {
		if *input.IsSyncing == true {
			action = "確認"

		} else {
			action = "取消"
		}
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "同步化",
			Values: "此報價至商機",
		})
	}

	if *input.IsFinal != *quoteBase.IsFinal {
		if *input.IsFinal == true {
			action = "確認"

		} else {
			action = "取消"
		}
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Values: "此報價為最終版",
		})
	}

	if *input.OpportunityID != *quoteBase.OpportunityID {
		opportunityBase, _ := m.OpportunityService.GetBySingle(&opportunityModel.Field{
			OpportunityID: *input.OpportunityID,
		})
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "商機",
			Values: "為" + *opportunityBase.Name,
		})

		if opportunityBase.AccountID != quoteBase.AccountID {
			accountBase, _ := m.AccountService.GetBySingle(&accountModel.Field{
				AccountID: *opportunityBase.AccountID,
			})
			records = append(records, historicalRecordModel.AddHistoricalRecord{
				Fields: "帳戶",
				Values: "為" + *accountBase.Name,
			})
		}
	}

	if *input.ExpirationDate != *quoteBase.ExpirationDate {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "到期日期",
			Values: "為" + input.ExpirationDate.Format("2006-01-02"),
		})
	}

	if *input.Description != *quoteBase.Description {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "描述",
			Values: "為" + *input.Description,
		})
	}

	if *input.Tax != *quoteBase.Tax {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "稅額",
			Values: "為" + strconv.FormatFloat(*input.Tax, 'f', -1, 64),
		})
	}

	if *input.ShippingAndHandling != *quoteBase.ShippingAndHandling {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "運費及其他費用",
			Values: "為" + strconv.FormatFloat(*input.ShippingAndHandling, 'f', -1, 64),
		})
	}

	for _, record := range records {
		_, err = m.HistoricalRecordService.WithTrx(trx).Create(&historicalRecordModel.Create{
			SourceID:   *quoteBase.QuoteID,
			Action:     action,
			Content:    sourceType + record.Fields + record.Values,
			ModifiedBy: *input.UpdatedBy,
		})
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, quoteBase.QuoteID)
}
