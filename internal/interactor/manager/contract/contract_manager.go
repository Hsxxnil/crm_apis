package contract

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	contractModel "app.eirc/internal/interactor/models/contracts"
	contractService "app.eirc/internal/interactor/service/contract"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *contractModel.Create) interface{}
	GetByList(input *contractModel.Fields) interface{}
	GetBySingle(input *contractModel.Field) interface{}
	Delete(input *contractModel.Field) interface{}
	Update(input *contractModel.Update) interface{}
}

type manager struct {
	ContractService contractService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		ContractService: contractService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *contractModel.Create) interface{} {
	defer trx.Rollback()

	contractBase, err := m.ContractService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, contractBase.ContractID)
}

func (m *manager) GetByList(input *contractModel.Fields) interface{} {
	output := &contractModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, contractBase, err := m.ContractService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	output.Pages = util.Pagination(quantity, output.Limit)
	contractByte, err := json.Marshal(contractBase)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = json.Unmarshal(contractByte, &output.Contracts)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for _, contractsBase := range contractBase {
		for _, contracts := range output.Contracts {
			contracts.AccountName = *contractsBase.Accounts.Name
			contracts.CreatedBy = *contractsBase.CreatedByUsers.Name
			contracts.UpdatedBy = *contractsBase.UpdatedByUsers.Name
		}
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *contractModel.Field) interface{} {
	contractBase, err := m.ContractService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &contractModel.Single{}
	contractByte, _ := json.Marshal(contractBase)
	err = json.Unmarshal(contractByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output.AccountName = *contractBase.Accounts.Name
	output.CreatedBy = *contractBase.CreatedByUsers.Name
	output.UpdatedBy = *contractBase.UpdatedByUsers.Name

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *contractModel.Field) interface{} {
	_, err := m.ContractService.GetBySingle(&contractModel.Field{
		ContractID: input.ContractID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.ContractService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *contractModel.Update) interface{} {
	contractBase, err := m.ContractService.GetBySingle(&contractModel.Field{
		ContractID: input.ContractID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.ContractService.Update(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, contractBase.ContractID)
}
