package contract

import (
	"encoding/json"
	"errors"
	"time"

	"app.eirc/internal/interactor/models/page"

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
	Create(trx *gorm.DB, input *contractModel.Create) (int, interface{})
	GetByList(input *contractModel.Fields) (int, interface{})
	GetBySingle(input *contractModel.Field) (int, interface{})
	Delete(input *contractModel.Field) (int, interface{})
	Update(input *contractModel.Update) (int, interface{})
}

type manager struct {
	ContractService contractService.Service
	OrderService    orderService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		ContractService: contractService.Init(db),
		OrderService:    orderService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *contractModel.Create) (int, interface{}) {
	defer trx.Rollback()

	input.EndDate = input.StartDate.AddDate(0, input.Term, 0)
	contractBase, err := m.ContractService.WithTrx(trx).Create(input)
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

	startDate := input.StartDate
	term := input.Term
	baseStartDate, _ := time.Parse("2006-01-02", *contractBase.StartDate)
	if input.StartDate == &baseStartDate {
		startDate = &baseStartDate
	}
	if input.Term == contractBase.Term {
		term = contractBase.Term
	}
	input.EndDate = startDate.AddDate(0, *term, 0)

	err = m.ContractService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	if input.AccountID != nil && input.AccountID != contractBase.AccountID {
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

	return code.Successful, code.GetCodeMessage(code.Successful, contractBase.ContractID)
}
