package order

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/helpers"

	accountModel "app.eirc/internal/interactor/models/accounts"
	historicalRecordModel "app.eirc/internal/interactor/models/historical_records"
	accountService "app.eirc/internal/interactor/service/account"
	historicalRecordService "app.eirc/internal/interactor/service/historical_record"

	"app.eirc/internal/interactor/pkg/util"

	contractModel "app.eirc/internal/interactor/models/contracts"
	orderModel "app.eirc/internal/interactor/models/orders"
	contractService "app.eirc/internal/interactor/service/contract"
	orderService "app.eirc/internal/interactor/service/order"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *orderModel.Create) (int, any)
	GetByList(input *orderModel.Fields) (int, any)
	GetBySingle(input *orderModel.Field) (int, any)
	GetBySingleProducts(input *orderModel.Field) (int, any)
	Delete(input *orderModel.Field) (int, any)
	Update(trx *gorm.DB, input *orderModel.Update) (int, any)
}

type manager struct {
	OrderService            orderService.Service
	ContractService         contractService.Service
	HistoricalRecordService historicalRecordService.Service
	AccountService          accountService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		OrderService:            orderService.Init(db),
		ContractService:         contractService.Init(db),
		HistoricalRecordService: historicalRecordService.Init(db),
		AccountService:          accountService.Init(db),
	}
}

const sourceType = "訂單"

func (m *manager) Create(trx *gorm.DB, input *orderModel.Create) (int, any) {
	defer trx.Rollback()

	// 同步契約的account_id
	contractBase, _ := m.ContractService.GetBySingle(&contractModel.Field{
		ContractID: input.ContractID,
	})
	input.AccountID = *contractBase.AccountID

	orderBase, err := m.OrderService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步新增訂單歷程記錄
	_, err = m.HistoricalRecordService.WithTrx(trx).Create(&historicalRecordModel.Create{
		SourceID:   *orderBase.OrderID,
		Action:     "建立",
		SourceType: sourceType,
		ModifiedBy: *orderBase.CreatedBy,
	})
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, orderBase.OrderID)
}

func (m *manager) GetByList(input *orderModel.Fields) (int, any) {
	output := &orderModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, orderBase, err := m.OrderService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	orderByte, err := json.Marshal(orderBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(orderByte, &output.Orders)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, orders := range output.Orders {
		orders.AccountName = *orderBase[i].Accounts.Name
		orders.ContractCode = *orderBase[i].Contracts.Code
		orders.CreatedBy = *orderBase[i].CreatedByUsers.Name
		orders.UpdatedBy = *orderBase[i].UpdatedByUsers.Name
		orders.ActivatedBy = *orderBase[i].ActivatedByUsers.Name
		orders.ActivatedAt = orderBase[i].ActivatedAt
		for _, products := range orderBase[i].OrderProducts {
			// 計算訂單總計
			orders.GrandTotal += *products.SubTotal
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *orderModel.Field) (int, any) {
	orderBase, err := m.OrderService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &orderModel.Single{}
	orderByte, _ := json.Marshal(orderBase)
	err = json.Unmarshal(orderByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.AccountName = *orderBase.Accounts.Name
	output.ContractCode = *orderBase.Contracts.Code
	output.CreatedBy = *orderBase.CreatedByUsers.Name
	output.UpdatedBy = *orderBase.UpdatedByUsers.Name
	output.ActivatedBy = *orderBase.ActivatedByUsers.Name
	output.ActivatedAt = orderBase.ActivatedAt
	for _, products := range orderBase.OrderProducts {
		output.GrandTotal += *products.SubTotal
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingleProducts(input *orderModel.Field) (int, any) {
	orderBase, err := m.OrderService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &orderModel.SingleProducts{}
	orderByte, _ := json.Marshal(orderBase)
	err = json.Unmarshal(orderByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.AccountName = *orderBase.Accounts.Name
	output.ContractCode = *orderBase.Contracts.Code
	output.CreatedBy = *orderBase.CreatedByUsers.Name
	output.UpdatedBy = *orderBase.UpdatedByUsers.Name
	output.ActivatedBy = *orderBase.ActivatedByUsers.Name
	output.ActivatedAt = orderBase.ActivatedAt
	for i, products := range orderBase.OrderProducts {
		output.OrderProducts[i].ProductName = *products.Products.Name
		output.OrderProducts[i].ProductPrice = *products.Products.Price
		output.OrderProducts[i].CreatedBy = *products.Products.CreatedByUsers.Name
		output.OrderProducts[i].UpdatedBy = *products.Products.UpdatedByUsers.Name
		output.GrandTotal += *products.SubTotal
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *orderModel.Field) (int, any) {
	_, err := m.OrderService.GetBySingle(&orderModel.Field{
		OrderID: input.OrderID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.OrderService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(trx *gorm.DB, input *orderModel.Update) (int, any) {
	defer trx.Rollback()

	orderBase, err := m.OrderService.GetBySingle(&orderModel.Field{
		OrderID: input.OrderID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 判斷該訂單是否啟用
	if *orderBase.Status != *input.Status {
		if *input.Status == "啟動中" {
			input.ActivatedBy = input.UpdatedBy
		} else {
			input.ActivatedBy = nil
		}
	}

	// 同步更新契約的account_id至該訂單
	if input.ContractID != nil && *input.ContractID != *orderBase.ContractID {
		contractBase, _ := m.ContractService.GetBySingle(&contractModel.Field{
			ContractID: *input.ContractID,
		})
		input.AccountID = contractBase.AccountID
	}

	err = m.OrderService.WithTrx(trx).Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步新增契約歷程記錄
	var records []historicalRecordModel.AddHistoricalRecord

	if input.Status != nil && *input.Status != *orderBase.Status {
		helpers.AddHistoricalRecord(&records, "修改", "狀態為", *input.Status)
	}

	if input.StartDate != nil && *input.StartDate != *orderBase.StartDate {
		helpers.AddHistoricalRecord(&records, "修改", "開始日期為", input.StartDate.UTC().Format("2006-01-02T15:04:05.999999Z"))
	}

	if input.ContractID != nil && *input.ContractID != *orderBase.ContractID {
		contractBase, _ := m.ContractService.GetBySingle(&contractModel.Field{
			ContractID: *input.ContractID,
		})
		helpers.AddHistoricalRecord(&records, "修改", "契約號碼為", *contractBase.Code)

		if contractBase.AccountID != orderBase.AccountID {
			accountBase, _ := m.AccountService.GetBySingle(&accountModel.Field{
				AccountID: *contractBase.AccountID,
			})
			helpers.AddHistoricalRecord(&records, "修改", "帳戶為", *accountBase.Name)
		}
	}

	if input.Description != nil {
		if *input.Description != *orderBase.Description {
			if *input.Description == "" {
				helpers.AddHistoricalRecord(&records, "清除", "描述", "")
			} else {
				helpers.AddHistoricalRecord(&records, "修改", "描述為", *input.Description)
			}
		}
	} else if *orderBase.Description != "" {
		helpers.AddHistoricalRecord(&records, "清除", "描述", "")
	}

	for _, record := range records {
		_, err = m.HistoricalRecordService.WithTrx(trx).Create(&historicalRecordModel.Create{
			SourceID:   *orderBase.ContractID,
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
	return code.Successful, code.GetCodeMessage(code.Successful, orderBase.OrderID)
}
