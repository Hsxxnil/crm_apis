package industry

import (
	"encoding/json"
	"errors"

	industryModel "crm/internal/interactor/models/industries"
	industryService "crm/internal/interactor/service/industry"

	"gorm.io/gorm"

	"crm/internal/interactor/pkg/util/code"
	"crm/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *industryModel.Create) (int, any)
	GetByList(input *industryModel.Field) (int, any)
	GetBySingle(input *industryModel.Field) (int, any)
	Delete(input *industryModel.Field) (int, any)
	Update(input *industryModel.Update) (int, any)
}

type manager struct {
	IndustryService industryService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		IndustryService: industryService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *industryModel.Create) (int, any) {
	defer trx.Rollback()

	industryBase, err := m.IndustryService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, industryBase.IndustryID)
}

func (m *manager) GetByList(input *industryModel.Field) (int, any) {
	output := &industryModel.List{}
	industryBase, err := m.IndustryService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	industryByte, err := json.Marshal(industryBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	err = json.Unmarshal(industryByte, &output.Industries)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *industryModel.Field) (int, any) {
	industryBase, err := m.IndustryService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &industryModel.Single{}
	industryByte, _ := json.Marshal(industryBase)
	err = json.Unmarshal(industryByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *industryModel.Field) (int, any) {
	_, err := m.IndustryService.GetBySingle(&industryModel.Field{
		IndustryID: input.IndustryID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.IndustryService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *industryModel.Update) (int, any) {
	industryBase, err := m.IndustryService.GetBySingle(&industryModel.Field{
		IndustryID: input.IndustryID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.IndustryService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, industryBase.IndustryID)
}
