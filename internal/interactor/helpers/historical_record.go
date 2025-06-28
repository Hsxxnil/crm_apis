package helpers

import (
	historicalRecordModel "crm/internal/interactor/models/historical_records"
)

// AddHistoricalRecord is Helper function to create historical_record.
func AddHistoricalRecord(input *[]historicalRecordModel.AddHistoricalRecord, action, field, value string) {
	*input = append(*input, historicalRecordModel.AddHistoricalRecord{
		Actions: action,
		Fields:  field,
		Values:  value,
	})
}
