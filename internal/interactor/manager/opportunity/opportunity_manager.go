package opportunity

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	opportunityModel "app.eirc/internal/interactor/models/opportunities"
	opportunityService "app.eirc/internal/interactor/service/opportunity"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *opportunityModel.Create) interface{}
	GetByList(input *opportunityModel.Fields) interface{}
	GetBySingle(input *opportunityModel.Field) interface{}
	Delete(input *opportunityModel.Field) interface{}
	Update(input *opportunityModel.Update) interface{}
}

type manager struct {
	OpportunityService opportunityService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		OpportunityService: opportunityService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *opportunityModel.Create) interface{} {
	defer trx.Rollback()

	opportunityBase, err := m.OpportunityService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, opportunityBase.OpportunityID)
}

func (m *manager) GetByList(input *opportunityModel.Fields) interface{} {
	output := &opportunityModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, opportunityBase, err := m.OpportunityService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	opportunityByte, err := json.Marshal(opportunityBase)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(opportunityByte, &output.Opportunities)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for _, opportunities := range output.Opportunities {
		opportunities.CreatedBy = *opportunities.CreatedByUsers.Name
		opportunities.CreatedByUsers = nil
		opportunities.UpdatedBy = *opportunities.UpdatedByUsers.Name
		opportunities.UpdatedByUsers = nil
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *opportunityModel.Field) interface{} {
	opportunityBase, err := m.OpportunityService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &opportunityModel.Single{}
	opportunityByte, _ := json.Marshal(opportunityBase)
	err = json.Unmarshal(opportunityByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output.CreatedBy = *output.CreatedByUsers.Name
	output.CreatedByUsers = nil
	output.UpdatedBy = *output.UpdatedByUsers.Name
	output.UpdatedByUsers = nil

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *opportunityModel.Field) interface{} {
	_, err := m.OpportunityService.GetBySingle(&opportunityModel.Field{
		OpportunityID: input.OpportunityID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.OpportunityService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *opportunityModel.Update) interface{} {
	opportunityBase, err := m.OpportunityService.GetBySingle(&opportunityModel.Field{
		OpportunityID: input.OpportunityID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.OpportunityService.Update(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, opportunityBase.OpportunityID)
}
