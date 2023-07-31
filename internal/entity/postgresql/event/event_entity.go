package event

import (
	"encoding/json"

	model "app.eirc/internal/entity/postgresql/db/events"
	"app.eirc/internal/interactor/pkg/util/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Entity interface {
	WithTrx(trx *gorm.DB) Entity
	Create(input *model.Base) (err error)
	GetByList(input *model.Base) (output []*model.Table, err error)
	GetBySingle(input *model.Base) (output *model.Table, err error)
	GetByQuantity(input *model.Base) (quantity int64, err error)
	Delete(input *model.Base) (err error)
	Update(input *model.Base) (err error)
}

type storage struct {
	db *gorm.DB
}

func Init(db *gorm.DB) Entity {
	return &storage{
		db: db,
	}
}

func (s *storage) WithTrx(trx *gorm.DB) Entity {
	return &storage{
		db: trx,
	}
}

func (s *storage) Create(input *model.Base) (err error) {
	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		return err
	}

	data := &model.Table{}
	err = json.Unmarshal(marshal, data)
	if err != nil {
		log.Error(err)
		return err
	}

	err = s.db.Model(&model.Table{}).Omit(clause.Associations).Create(&data).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *storage) GetByList(input *model.Base) (output []*model.Table, err error) {
	query := s.db.Model(&model.Table{}).
		Preload("EventUserMains", "is_deleted = ?", false).
		Preload("EventUserMains.Mains").
		Preload("EventUserAttendees", "is_deleted = ?", false).
		Preload("EventUserAttendees.Attendees").
		Preload("EventContacts", "is_deleted = ?", false).
		Preload("EventContacts.Contacts").
		Preload(clause.Associations)

	if input.EventID != nil {
		query.Where("event_id = ?", input.EventID)
	}

	if input.IsDeleted != nil {
		query.Where("events.is_deleted = ?", input.IsDeleted)
	}

	// filter
	isFiltered := false
	filter := s.db.Model(&model.Table{})
	if input.FilterSubject != "" {
		filter.Where("events.subject like ?", "%"+input.FilterSubject+"%")
		isFiltered = true
	}

	if input.FilterMainID != "" {
		if isFiltered {
			filter.Or("events.main_id like ?", "%"+input.FilterMainID+"%")
		} else {
			filter.Where("events.main_id like ?", "%"+input.FilterMainID+"%")
		}
	}

	if input.FilterAttendeeID != "" {
		if isFiltered {
			filter.Or("events.attendee_id like ?", "%"+input.FilterAttendeeID+"%")
		} else {
			filter.Where("events.attendee_id like ?", "%"+input.FilterAttendeeID+"%")
		}
	}

	if input.FilterType != "" {
		if isFiltered {
			filter.Or("events.type like ?", "%"+input.FilterType+"%")
		} else {
			filter.Where("events.type like ?", "%"+input.FilterType+"%")
		}
	}

	if input.FilterStartDate != "" && input.FilterEndDate.String() != "0001-01-01 00:00:00 +0000 UTC" {
		if isFiltered {
			filter.Or("start_date >= ? and end_date <= ?", input.FilterStartDate, input.FilterEndDate)
		} else {
			filter.Where("start_date >= ? and end_date <= ?", input.FilterStartDate, input.FilterEndDate)
		}
	}

	query.Where(filter)

	err = query.Order("created_at desc").Find(&output).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}

func (s *storage) GetBySingle(input *model.Base) (output *model.Table, err error) {
	query := s.db.Model(&model.Table{}).
		Preload("EventUserMains", "is_deleted = ?", false).
		Preload("EventUserMains.Mains").
		Preload("EventUserAttendees", "is_deleted = ?", false).
		Preload("EventUserAttendees.Attendees").
		Preload("EventContacts", "is_deleted = ?", false).
		Preload("EventContacts.Contacts").
		Preload(clause.Associations)

	if input.EventID != nil {
		query.Where("event_id = ?", input.EventID)
	}

	if input.IsDeleted != nil {
		query.Where("events.is_deleted = ?", input.IsDeleted)
	}

	err = query.First(&output).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}

func (s *storage) GetByQuantity(input *model.Base) (quantity int64, err error) {
	query := s.db.Model(&model.Table{})
	if input.EventID != nil {
		query.Where("event_id = ?", input.EventID)
	}

	if input.IsDeleted != nil {
		query.Where("events.is_deleted = ?", input.IsDeleted)
	}

	err = query.Count(&quantity).Select("*").Error
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return quantity, nil
}

func (s *storage) Update(input *model.Base) (err error) {
	query := s.db.Model(&model.Table{}).Omit(clause.Associations)
	data := map[string]any{}

	if input.Subject != nil {
		data["subject"] = input.Subject
	}

	if input.IsWhole != nil {
		data["is_whole"] = input.IsWhole
	}

	if input.StartDate != nil {
		data["start_date"] = input.StartDate
	}

	if input.EndDate != nil {
		data["end_date"] = input.EndDate
	}

	if input.AccountID != nil {
		data["account_id"] = input.AccountID
	}

	if input.Type != nil {
		data["type"] = input.Type
	}

	if input.Location != nil {
		data["location"] = input.Location
	}

	if input.Description != nil {
		data["description"] = input.Description
	}

	if input.IsDeleted != nil {
		data["is_deleted"] = input.IsDeleted
	}

	if input.UpdatedBy != nil {
		data["updated_by"] = input.UpdatedBy
	}

	if input.EventID != nil {
		query.Where("event_id = ?", input.EventID)
	}

	err = query.Select("*").Updates(data).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *storage) Delete(input *model.Base) (err error) {
	query := s.db.Model(&model.Table{}).Omit(clause.Associations)
	if input.EventID != nil {
		query.Where("event_id = ?", input.EventID)
	}

	err = query.Delete(&model.Table{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
