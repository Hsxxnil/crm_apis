package historical_record

import (
	"encoding/json"
	"errors"

	"crm/internal/interactor/pkg/util"

	historicalRecordModel "crm/internal/interactor/models/historical_records"
	historicalRecordService "crm/internal/interactor/service/historical_record"

	"gorm.io/gorm"

	"crm/internal/interactor/pkg/util/code"
	"crm/internal/interactor/pkg/util/log"
)

type Manager interface {
	GetByList(input *historicalRecordModel.Fields) (int, any)
	GetBySingle(input *historicalRecordModel.Field) (int, any)
}

type manager struct {
	HistoricalRecordService historicalRecordService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		HistoricalRecordService: historicalRecordService.Init(db),
	}
}

func (m *manager) GetByList(input *historicalRecordModel.Fields) (int, any) {
	output := &historicalRecordModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, historicalRecordBase, err := m.HistoricalRecordService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	historicalRecordByte, err := json.Marshal(historicalRecordBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(historicalRecordByte, &output.HistoricalRecords)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, historicalRecords := range output.HistoricalRecords {
		historicalRecords.ModifiedBy = *historicalRecordBase[i].ModifiedByUsers.Name
		historicalRecords.Content = *historicalRecordBase[i].Action + *historicalRecordBase[i].SourceType + *historicalRecordBase[i].Field
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *historicalRecordModel.Field) (int, any) {
	historicalRecordBase, err := m.HistoricalRecordService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &historicalRecordModel.Single{}
	historicalRecordByte, _ := json.Marshal(historicalRecordBase)
	err = json.Unmarshal(historicalRecordByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.ModifiedBy = *historicalRecordBase.ModifiedByUsers.Name
	output.Description = *historicalRecordBase.Action + *historicalRecordBase.SourceType + *historicalRecordBase.Field

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}
