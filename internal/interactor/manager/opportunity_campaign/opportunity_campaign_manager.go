package opportunity_campaign

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	opportunityCampaignModel "app.eirc/internal/interactor/models/opportunity_campaigns"
	opportunityCampaignService "app.eirc/internal/interactor/service/opportunity_campaign"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *opportunityCampaignModel.Create) (int, interface{})
	GetByList(input *opportunityCampaignModel.Fields) (int, interface{})
	GetBySingle(input *opportunityCampaignModel.Field) (int, interface{})
	Delete(input *opportunityCampaignModel.Field) (int, interface{})
	Update(input *opportunityCampaignModel.Update) (int, interface{})
}

type manager struct {
	OpportunityCampaignService opportunityCampaignService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		OpportunityCampaignService: opportunityCampaignService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *opportunityCampaignModel.Create) (int, interface{}) {
	defer trx.Rollback()

	opportunityCampaignBase, err := m.OpportunityCampaignService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, opportunityCampaignBase.OpportunityCampaignID)
}

func (m *manager) GetByList(input *opportunityCampaignModel.Fields) (int, interface{}) {
	input.IsDeleted = util.PointerBool(false)
	output := &opportunityCampaignModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, opportunityCampaignBase, err := m.OpportunityCampaignService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	opportunityCampaignByte, err := json.Marshal(opportunityCampaignBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(opportunityCampaignByte, &output.OpportunityCampaigns)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, OpportunityCampaigns := range output.OpportunityCampaigns {
		OpportunityCampaigns.CreatedBy = *opportunityCampaignBase[i].CreatedByUsers.Name
		OpportunityCampaigns.UpdatedBy = *opportunityCampaignBase[i].UpdatedByUsers.Name
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *opportunityCampaignModel.Field) (int, interface{}) {
	input.IsDeleted = util.PointerBool(false)
	opportunityCampaignBase, err := m.OpportunityCampaignService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &opportunityCampaignModel.Single{}
	opportunityCampaignByte, _ := json.Marshal(opportunityCampaignBase)
	err = json.Unmarshal(opportunityCampaignByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.CreatedBy = *opportunityCampaignBase.CreatedByUsers.Name
	output.UpdatedBy = *opportunityCampaignBase.UpdatedByUsers.Name

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *opportunityCampaignModel.Field) (int, interface{}) {
	_, err := m.OpportunityCampaignService.GetBySingle(&opportunityCampaignModel.Field{
		OpportunityCampaignID: input.OpportunityCampaignID,
		IsDeleted:             util.PointerBool(false),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.OpportunityCampaignService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *opportunityCampaignModel.Update) (int, interface{}) {
	opportunityCampaignBase, err := m.OpportunityCampaignService.GetBySingle(&opportunityCampaignModel.Field{
		OpportunityCampaignID: input.OpportunityCampaignID,
		IsDeleted:             util.PointerBool(false),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.OpportunityCampaignService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, opportunityCampaignBase.OpportunityCampaignID)
}
