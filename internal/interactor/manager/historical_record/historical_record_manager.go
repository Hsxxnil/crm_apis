package historical_record

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	historicalRecordModel "app.eirc/internal/interactor/models/historical_records"
	historicalRecordService "app.eirc/internal/interactor/service/historical_record"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	GetByList(input *historicalRecordModel.Fields) (int, interface{})
	GetBySingle(input *historicalRecordModel.Field) (int, interface{})
}

type manager struct {
	HistoricalRecordService historicalRecordService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		HistoricalRecordService: historicalRecordService.Init(db),
	}
}

func (m *manager) GetByList(input *historicalRecordModel.Fields) (int, interface{}) {
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
		historicalRecords.Description = *historicalRecordBase[i].Action + *historicalRecordBase[i].Content
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *historicalRecordModel.Field) (int, interface{}) {
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
	output.Description = *historicalRecordBase.Action + *historicalRecordBase.Content

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}
