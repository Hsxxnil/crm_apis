package contract

import (
	"encoding/json"
	"errors"
	"strconv"

	accountModel "app.eirc/internal/interactor/models/accounts"

	"app.eirc/internal/interactor/models/page"

	"app.eirc/internal/interactor/pkg/util"

	contractModel "app.eirc/internal/interactor/models/contracts"
	historicalRecordModel "app.eirc/internal/interactor/models/historical_records"
	orderModel "app.eirc/internal/interactor/models/orders"
	accountService "app.eirc/internal/interactor/service/account"
	contractService "app.eirc/internal/interactor/service/contract"
	historicalRecordService "app.eirc/internal/interactor/service/historical_record"
	orderService "app.eirc/internal/interactor/service/order"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *contractModel.Create) (int, interface{})
	GetByList(input *contractModel.Fields) (int, interface{})
	GetBySingle(input *contractModel.Field) (int, interface{})
	Delete(input *contractModel.Field) (int, interface{})
	Update(input *contractModel.Update) (int, interface{})
}

type manager struct {
	ContractService         contractService.Service
	OrderService            orderService.Service
	HistoricalRecordService historicalRecordService.Service
	AccountService          accountService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		ContractService:         contractService.Init(db),
		OrderService:            orderService.Init(db),
		HistoricalRecordService: historicalRecordService.Init(db),
		AccountService:          accountService.Init(db),
	}
}

const sourceType = "契約"

func (m *manager) Create(trx *gorm.DB, input *contractModel.Create) (int, interface{}) {
	defer trx.Rollback()

	input.EndDate = input.StartDate.AddDate(0, input.Term, 0)
	contractBase, err := m.ContractService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 新增契約歷程記錄
	_, err = m.HistoricalRecordService.WithTrx(trx).Create(&historicalRecordModel.Create{
		SourceID:   *contractBase.ContractID,
		Action:     "建立",
		Content:    sourceType,
		ModifiedBy: *contractBase.CreatedBy,
	})
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, contractBase.ContractID)
}

func (m *manager) GetByList(input *contractModel.Fields) (int, interface{}) {
	output := &contractModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, contractBase, err := m.ContractService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	output.Pages = util.Pagination(quantity, output.Limit)
	contractByte, err := json.Marshal(contractBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = json.Unmarshal(contractByte, &output.Contracts)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, contracts := range output.Contracts {
		contracts.AccountName = *contractBase[i].Accounts.Name
		contracts.CreatedBy = *contractBase[i].CreatedByUsers.Name
		contracts.UpdatedBy = *contractBase[i].UpdatedByUsers.Name
		contracts.SalespersonName = *contractBase[i].Salespeople.Name
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *contractModel.Field) (int, interface{}) {
	contractBase, err := m.ContractService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &contractModel.Single{}
	contractByte, _ := json.Marshal(contractBase)
	err = json.Unmarshal(contractByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.AccountName = *contractBase.Accounts.Name
	output.CreatedBy = *contractBase.CreatedByUsers.Name
	output.UpdatedBy = *contractBase.UpdatedByUsers.Name
	output.SalespersonName = *contractBase.Salespeople.Name

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *contractModel.Field) (int, interface{}) {
	_, err := m.ContractService.GetBySingle(&contractModel.Field{
		ContractID: input.ContractID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.ContractService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *contractModel.Update) (int, interface{}) {
	contractBase, err := m.ContractService.GetBySingle(&contractModel.Field{
		ContractID: input.ContractID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	startDate := contractBase.StartDate
	term := contractBase.Term
	if input.StartDate != nil && input.StartDate != contractBase.StartDate {
		startDate = input.StartDate
	}
	if input.Term != nil && input.Term != contractBase.Term {
		term = input.Term
	}
	input.EndDate = startDate.AddDate(0, *term, 0)

	err = m.ContractService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步帳戶至orders
	if input.AccountID != nil && *input.AccountID != *contractBase.AccountID {
		_, orders, err := m.OrderService.GetByList(&orderModel.Fields{
			Field: orderModel.Field{
				ContractID: util.PointerString(input.ContractID),
			},
			Pagination: page.Pagination{
				Page:  1,
				Limit: 20,
			},
		})
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		for _, orderBase := range orders {
			err = m.OrderService.Update(&orderModel.Update{
				OrderID:   *orderBase.OrderID,
				AccountID: input.AccountID,
				UpdatedBy: input.UpdatedBy,
			})
			if err != nil {
				log.Error(err)
				return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
			}
		}
	}

	// 新增契約歷程記錄
	var (
		colum string
		value string
	)

	switch {
	case *input.AccountID != *contractBase.AccountID:
		accountBase, _ := m.AccountService.GetBySingle(&accountModel.Field{
			AccountID: *input.AccountID,
		})
		value = "為" + *accountBase.Name
		colum = "帳戶"
	case *input.Status != *contractBase.Status:
		value = "為" + *input.Status
		colum = "狀態"
	case *input.StartDate != *contractBase.StartDate:
		value = "為" + input.StartDate.Format("2006-01-02")
		colum = "開始日期"
	case *input.Term != *contractBase.Term:
		value = "為" + strconv.Itoa(*input.Term) + "個月"
		colum = "有效期限"
	case *input.Description != *contractBase.Description:
		colum = "描述"
	}

	if colum != "" {
		_, err = m.HistoricalRecordService.Create(&historicalRecordModel.Create{
			SourceID:   *contractBase.ContractID,
			Action:     "修改",
			Content:    sourceType + colum + value,
			ModifiedBy: *input.UpdatedBy,
		})
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, contractBase.ContractID)
}
