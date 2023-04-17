package industry

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	industryModel "app.eirc/internal/interactor/models/industries"
	industryService "app.eirc/internal/interactor/service/industry"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *industryModel.Create) interface{}
	GetByList(input *industryModel.Fields) interface{}
	GetBySingle(input *industryModel.Field) interface{}
	Delete(input *industryModel.Field) interface{}
	Update(input *industryModel.Update) interface{}
}

type manager struct {
	IndustryService industryService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		IndustryService: industryService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *industryModel.Create) interface{} {
	defer trx.Rollback()

	industryBase, err := m.IndustryService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, industryBase.IndustryID)
}

func (m *manager) GetByList(input *industryModel.Fields) interface{} {
	output := &industryModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, industryBase, err := m.IndustryService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	industryByte, err := json.Marshal(industryBase)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(industryByte, &output.Industries)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *industryModel.Field) interface{} {
	industryBase, err := m.IndustryService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &industryModel.Single{}
	industryByte, _ := json.Marshal(industryBase)
	err = json.Unmarshal(industryByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *industryModel.Field) interface{} {
	_, err := m.IndustryService.GetBySingle(&industryModel.Field{
		IndustryID: input.IndustryID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.IndustryService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *industryModel.Update) interface{} {
	industryBase, err := m.IndustryService.GetBySingle(&industryModel.Field{
		IndustryID: input.IndustryID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.IndustryService.Update(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, industryBase.IndustryID)
}
