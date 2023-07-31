package lead

import (
	"encoding/json"

	model "app.eirc/internal/entity/postgresql/db/leads"
	"app.eirc/internal/interactor/pkg/util/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Entity interface {
	WithTrx(trx *gorm.DB) Entity
	Create(input *model.Base) (err error)
	GetByList(input *model.Base) (quantity int64, output []*model.Table, err error)
	GetByListNoPagination(input *model.Base) (output []*model.Table, err error)
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

func (s *storage) GetByList(input *model.Base) (quantity int64, output []*model.Table, err error) {
	query := s.db.Model(&model.Table{}).Count(&quantity).Joins("Salespeople").Joins("Accounts").Preload(clause.Associations)
	if input.LeadID != nil {
		query.Where("lead_id = ?", input.LeadID)
	}

	if input.IsDeleted != nil {
		query.Where("leads.is_deleted = ?", input.IsDeleted)
	}

	if input.Sort.Field != "" && input.Sort.Direction != "" {
		if input.Sort.Field == "salesperson_name" && input.Sort.Direction != "" {
			query.Order(`"Salespeople".name` + " " + input.Sort.Direction)
		} else if input.Sort.Field == "account_name" && input.Sort.Direction != "" {
			query.Order(`"Accounts".name` + " " + input.Sort.Direction)
		} else {
			query.Order(input.Sort.Field + " " + input.Sort.Direction)
		}
	}

	// filter
	isFiltered := false
	filter := s.db.Model(&model.Table{})
	if input.FilterDescription != "" {
		filter.Where("leads.description like ?", "%"+input.FilterDescription+"%")
		isFiltered = true
	}

	if input.FilterAccountName != "" {
		if isFiltered {
			filter.Or(`"Accounts".name like ?`, "%"+input.FilterAccountName+"%")
		} else {
			filter.Where(`"Accounts".name like ?`, "%"+input.FilterAccountName+"%")
		}
	}

	if input.FilterRating != "" {
		if isFiltered {
			filter.Or("leads.rating like ?", "%"+input.FilterRating+"%")
		} else {
			filter.Where("leads.rating like ?", "%"+input.FilterRating+"%")
		}
	}

	if input.FilterSource != "" {
		if isFiltered {
			filter.Or("leads.source like ?", "%"+input.FilterSource+"%")
		} else {
			filter.Where("leads.source like ?", "%"+input.FilterSource+"%")
		}
	}

	if input.FilterStatus != "" {
		if isFiltered {
			filter.Or("leads.status = ?", input.FilterStatus)
		} else {
			filter.Where("leads.status = ?", input.FilterStatus)
		}
	}

	if input.FilterSalespersonName != "" {
		if isFiltered {
			filter.Or(`"Salespeople".name like ?`, "%"+input.FilterSalespersonName+"%")
		} else {
			filter.Where(`"Salespeople".name like ?`, "%"+input.FilterSalespersonName+"%")
		}
	}

	query.Where(filter)

	err = query.Count(&quantity).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("created_at desc").Find(&output).Error
	if err != nil {
		log.Error(err)
		return 0, nil, err
	}

	return quantity, output, nil
}

func (s *storage) GetByListNoPagination(input *model.Base) (output []*model.Table, err error) {
	query := s.db.Model(&model.Table{}).Preload(clause.Associations)
	if input.LeadID != nil {
		query.Where("lead_id = ?", input.LeadID)
	}

	if input.IsDeleted != nil {
		query.Where("is_deleted = ?", input.IsDeleted)
	}

	err = query.Order("created_at desc").Find(&output).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}

func (s *storage) GetBySingle(input *model.Base) (output *model.Table, err error) {
	query := s.db.Model(&model.Table{}).Preload(clause.Associations)
	if input.LeadID != nil {
		query.Where("lead_id = ?", input.LeadID)
	}

	if input.IsDeleted != nil {
		query.Where("is_deleted = ?", input.IsDeleted)
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
	if input.LeadID != nil {
		query.Where("lead_id = ?", input.LeadID)
	}

	if input.IsDeleted != nil {
		query.Where("is_deleted = ?", input.IsDeleted)
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

	if input.Status != nil {
		data["status"] = input.Status
	}

	if input.Description != nil {
		data["description"] = input.Description
	}

	if input.Source != nil {
		data["source"] = input.Source
	}

	if input.Rating != nil {
		data["rating"] = input.Rating
	}

	if input.SalespersonID != nil {
		data["salesperson_id"] = input.SalespersonID
	}

	if input.IsDeleted != nil {
		data["is_deleted"] = input.IsDeleted
	}

	if input.UpdatedBy != nil {
		data["updated_by"] = input.UpdatedBy
	}

	if input.LeadID != nil {
		query.Where("lead_id = ?", input.LeadID)
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
	if input.LeadID != nil {
		query.Where("lead_id = ?", input.LeadID)
	}

	err = query.Delete(&model.Table{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
