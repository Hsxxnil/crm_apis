package contract

import (
	"encoding/json"
	"errors"
	"strconv"

	"app.eirc/internal/interactor/helpers"

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
	Create(trx *gorm.DB, input *contractModel.Create) (int, any)
	GetByList(input *contractModel.Fields) (int, any)
	GetByListNoPagination(input *contractModel.FieldsNoPagination) (int, any)
	GetBySingle(input *contractModel.Field) (int, any)
	Delete(input *contractModel.Field) (int, any)
	Update(trx *gorm.DB, input *contractModel.Update) (int, any)
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

func (m *manager) Create(trx *gorm.DB, input *contractModel.Create) (int, any) {
	defer trx.Rollback()

	// 同步商機的account_id
	opportunityBase, _ := m.OpportunityService.GetBySingle(&opportunityModel.Field{
		OpportunityID: input.OpportunityID,
		IsDeleted:     util.PointerBool(false),
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
		SourceType: sourceType,
		ModifiedBy: *contractBase.CreatedBy,
	})
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, contractBase.ContractID)
}

func (m *manager) GetByList(input *contractModel.Fields) (int, any) {
	input.IsDeleted = util.PointerBool(false)
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

func (m *manager) GetByListNoPagination(input *contractModel.FieldsNoPagination) (int, any) {
	input.IsDeleted = util.PointerBool(false)
	output := &contractModel.ListNoPagination{}
	contractBase, err := m.ContractService.GetByListNoPagination(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
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

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *contractModel.Field) (int, any) {
	input.IsDeleted = util.PointerBool(false)
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

func (m *manager) Delete(input *contractModel.Field) (int, any) {
	_, err := m.ContractService.GetBySingle(&contractModel.Field{
		ContractID: input.ContractID,
		IsDeleted:  util.PointerBool(false),
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

func (m *manager) Update(trx *gorm.DB, input *contractModel.Update) (int, any) {
	defer trx.Rollback()

	contractBase, err := m.ContractService.GetBySingle(&contractModel.Field{
		ContractID: input.ContractID,
		IsDeleted:  util.PointerBool(false),
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
			IsDeleted:     util.PointerBool(false),
		})
		input.AccountID = opportunityBase.AccountID

		// 同步修改帳戶至訂單
		orders, err := m.OrderService.GetByListNoPagination(&orderModel.Field{
			ContractID: util.PointerString(input.ContractID),
			IsDeleted:  util.PointerBool(false),
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

	if input.OpportunityID != nil && *input.OpportunityID != *contractBase.OpportunityID {
		opportunityBase, _ := m.OpportunityService.GetBySingle(&opportunityModel.Field{
			OpportunityID: *input.OpportunityID,
			IsDeleted:     util.PointerBool(false),
		})
		helpers.AddHistoricalRecord(&records, "修改", "商機為", *opportunityBase.Name)

		if opportunityBase.AccountID != contractBase.AccountID {
			accountBase, _ := m.AccountService.GetBySingle(&accountModel.Field{
				AccountID: *opportunityBase.AccountID,
				IsDeleted: util.PointerBool(false),
			})
			helpers.AddHistoricalRecord(&records, "修改", "帳戶為", *accountBase.Name)
		}
	}

	if input.Status != nil && *input.Status != *contractBase.Status {
		helpers.AddHistoricalRecord(&records, "修改", "狀態為", *input.Status)
	}

	if input.StartDate != nil && *input.StartDate != *contractBase.StartDate {
		helpers.AddHistoricalRecord(&records, "修改", "開始日期為", input.StartDate.UTC().Format("2006-01-02T15:04:05.999999Z"))
	}

	if input.Term != nil && *input.Term != *contractBase.Term {
		helpers.AddHistoricalRecord(&records, "修改", "有效期限為", strconv.Itoa(*input.Term)+"個月")
	}

	if input.Description != nil {
		if *input.Description != *contractBase.Description {
			if *input.Description == "" {
				helpers.AddHistoricalRecord(&records, "清除", "描述", "")
			} else {
				helpers.AddHistoricalRecord(&records, "修改", "描述為", *input.Description)
			}
		}
	} else if *contractBase.Description != "" {
		helpers.AddHistoricalRecord(&records, "清除", "描述", "")
	}

	if input.SalespersonID != nil && *input.SalespersonID != *contractBase.SalespersonID {
		salespersonBase, _ := m.UserService.GetBySingle(&userModel.Field{
			UserID:    *input.SalespersonID,
			IsDeleted: util.PointerBool(false),
		})
		helpers.AddHistoricalRecord(&records, "修改", "業務員為", *salespersonBase.Name)
	}

	for _, record := range records {
		_, err = m.HistoricalRecordService.WithTrx(trx).Create(&historicalRecordModel.Create{
			SourceID:   *contractBase.ContractID,
			Action:     record.Actions,
			SourceType: sourceType,
			Field:      record.Fields,
			Value:      record.Values,
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
