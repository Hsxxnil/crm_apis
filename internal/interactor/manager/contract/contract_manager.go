package contract

import (
	"encoding/json"
	"errors"
	"strconv"

	userModel "app.eirc/internal/interactor/models/users"
	userService "app.eirc/internal/interactor/service/user"

	orderModel "app.eirc/internal/interactor/models/orders"
	"app.eirc/internal/interactor/pkg/util"

	accountModel "app.eirc/internal/interactor/models/accounts"
	contractModel "app.eirc/internal/interactor/models/contracts"
	historicalRecordModel "app.eirc/internal/interactor/models/historical_records"
	opportunityModel "app.eirc/internal/interactor/models/opportunities"
	accountService "app.eirc/internal/interactor/service/account"
	contractService "app.eirc/internal/interactor/service/contract"
	historicalRecordService "app.eirc/internal/interactor/service/historical_record"
	opportunityService "app.eirc/internal/interactor/service/opportunity"
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
	Update(trx *gorm.DB, input *contractModel.Update) (int, interface{})
}

type manager struct {
	ContractService         contractService.Service
	OrderService            orderService.Service
	HistoricalRecordService historicalRecordService.Service
	AccountService          accountService.Service
	OpportunityService      opportunityService.Service
	UserService             userService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		ContractService:         contractService.Init(db),
		OrderService:            orderService.Init(db),
		HistoricalRecordService: historicalRecordService.Init(db),
		AccountService:          accountService.Init(db),
		OpportunityService:      opportunityService.Init(db),
		UserService:             userService.Init(db),
	}
}

const sourceType = "契約"

func (m *manager) Create(trx *gorm.DB, input *contractModel.Create) (int, interface{}) {
	defer trx.Rollback()

	// 同步商機的account_id
	opportunityBase, _ := m.OpportunityService.GetBySingle(&opportunityModel.Field{
		OpportunityID: input.OpportunityID,
	})
	input.AccountID = *opportunityBase.AccountID
	input.EndDate = input.StartDate.AddDate(0, input.Term, 0)
	contractBase, err := m.ContractService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步新增契約歷程記錄
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
		contracts.OpportunityName = *contractBase[i].Opportunities.Name
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
	output.OpportunityName = *contractBase.Opportunities.Name

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

func (m *manager) Update(trx *gorm.DB, input *contractModel.Update) (int, interface{}) {
	defer trx.Rollback()

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

	// 計算契約結束日期
	startDate := contractBase.StartDate
	term := contractBase.Term
	if input.StartDate != nil && input.StartDate != contractBase.StartDate {
		startDate = input.StartDate
	}
	if input.Term != nil && input.Term != contractBase.Term {
		term = input.Term
	}
	input.EndDate = startDate.AddDate(0, *term, 0)

	// 同步更新商機的account_id至該契約
	if input.OpportunityID != nil && *input.OpportunityID != *contractBase.OpportunityID {
		opportunityBase, _ := m.OpportunityService.GetBySingle(&opportunityModel.Field{
			OpportunityID: *input.OpportunityID,
		})
		input.AccountID = opportunityBase.AccountID

		// 同步修改帳戶至orders
		orders, err := m.OrderService.GetByListNoQuantity(&orderModel.Field{
			ContractID: util.PointerString(input.ContractID),
		})
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		for _, orderBase := range orders {
			err = m.OrderService.WithTrx(trx).Update(&orderModel.Update{
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

	err = m.ContractService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步新增契約歷程記錄
	var records []historicalRecordModel.AddHistoricalRecord

	if *input.OpportunityID != *contractBase.OpportunityID {
		opportunityBase, _ := m.OpportunityService.GetBySingle(&opportunityModel.Field{
			OpportunityID: *input.OpportunityID,
		})
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "商機",
			Values: "為" + *opportunityBase.Name,
		})

		if opportunityBase.AccountID != contractBase.AccountID {
			accountBase, _ := m.AccountService.GetBySingle(&accountModel.Field{
				AccountID: *opportunityBase.AccountID,
			})
			records = append(records, historicalRecordModel.AddHistoricalRecord{
				Fields: "帳戶",
				Values: "為" + *accountBase.Name,
			})
		}
	}

	if *input.Status != *contractBase.Status {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "狀態",
			Values: "為" + *input.Status,
		})
	}

	if *input.StartDate != *contractBase.StartDate {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "開始日期",
			Values: "為" + input.StartDate.Format("2006-01-02"),
		})
	}

	if *input.Term != *contractBase.Term {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "有效期限",
			Values: "為" + strconv.Itoa(*input.Term) + "個月",
		})
	}

	if *input.Description != *contractBase.Description {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "描述",
			Values: "為" + *input.Description,
		})
	}

	if *input.SalespersonID != *contractBase.SalespersonID {
		salespersonBase, _ := m.UserService.GetBySingle(&userModel.Field{
			UserID: *input.SalespersonID,
		})
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "業務員",
			Values: "為" + *salespersonBase.Name,
		})
	}

	for _, record := range records {
		_, err = m.HistoricalRecordService.WithTrx(trx).Create(&historicalRecordModel.Create{
			SourceID:   *contractBase.ContractID,
			Action:     "修改",
			Content:    sourceType + record.Fields + record.Values,
			ModifiedBy: *input.UpdatedBy,
		})
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, contractBase.ContractID)
}
