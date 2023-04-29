package campaign

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	campaignModel "app.eirc/internal/interactor/models/campaigns"
	campaignService "app.eirc/internal/interactor/service/campaign"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *campaignModel.Create) interface{}
	GetByList(input *campaignModel.Fields) interface{}
	GetBySingle(input *campaignModel.Field) interface{}
	Delete(input *campaignModel.Field) interface{}
	Update(input *campaignModel.Update) interface{}
}

type manager struct {
	CampaignService campaignService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		CampaignService: campaignService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *campaignModel.Create) interface{} {
	defer trx.Rollback()

	campaignBase, err := m.CampaignService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, campaignBase.CampaignID)
}

func (m *manager) GetByList(input *campaignModel.Fields) interface{} {
	output := &campaignModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, campaignBase, err := m.CampaignService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	campaignByte, err := json.Marshal(campaignBase)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(campaignByte, &output.Campaigns)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *campaignModel.Field) interface{} {
	campaignBase, err := m.CampaignService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &campaignModel.Single{}
	campaignByte, _ := json.Marshal(campaignBase)
	err = json.Unmarshal(campaignByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *campaignModel.Field) interface{} {
	_, err := m.CampaignService.GetBySingle(&campaignModel.Field{
		CampaignID: input.CampaignID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.CampaignService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *campaignModel.Update) interface{} {
	campaignBase, err := m.CampaignService.GetBySingle(&campaignModel.Field{
		CampaignID: input.CampaignID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.CampaignService.Update(input)
	if err != nil {
		if err.Error() == "code already exists" {
			return code.GetCodeMessage(code.BadRequest, err.Error())
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, campaignBase.CampaignID)
}
