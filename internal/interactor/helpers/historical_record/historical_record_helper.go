package historical_record

import (
	historicalRecordModel "app.eirc/internal/interactor/models/historical_records"
)

func AddHistoricalRecord(input *[]historicalRecordModel.AddHistoricalRecord, action, field, value string) {
	*input = append(*input, historicalRecordModel.AddHistoricalRecord{
		Actions: action,
		Fields:  field,
		Values:  value,
	})
}
