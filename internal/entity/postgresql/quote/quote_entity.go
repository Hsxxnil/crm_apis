package quote

import (
	"encoding/json"

	model "app.eirc/internal/entity/postgresql/db/quotes"
	"app.eirc/internal/interactor/pkg/util/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Entity interface {
	WithTrx(trx *gorm.DB) Entity
	Create(input *model.Base) (err error)
	GetByList(input *model.Base) (quantity int64, output []*model.Table, err error)
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
	query := s.db.Model(&model.Table{}).Count(&quantity).
		Joins("Opportunities", s.db.Where(`"Opportunities".is_deleted= ?`, false)).
		Preload(clause.Associations)

	if input.QuoteID != nil {
		query.Where("quote_id = ?", input.QuoteID)
	}

	if input.IsDeleted != nil {
		query.Where("quotes.is_deleted = ?", input.IsDeleted)
	}

	if input.Sort.Field != "" && input.Sort.Direction != "" {
		if input.Sort.Field == "opportunity_name" && input.Sort.Direction != "" {
			query.Order(`"Opportunities".name` + " " + input.Sort.Direction)
		} else {
			query.Order(input.Sort.Field + " " + input.Sort.Direction)
		}
	}

	// filter
	isFiltered := false
	filter := s.db.Model(&model.Table{})
	if input.FilterName != "" {
		filter.Where("quotes.name like ?", "%"+input.FilterName+"%")
		isFiltered = true
	}

	if input.FilterOpportunityName != "" {
		if isFiltered {
			filter.Or(`"Opportunities".name like ?`, "%"+input.FilterOpportunityName+"%")
		} else {
			filter.Where(`"Opportunities".name like ?`, "%"+input.FilterOpportunityName+"%")
		}
	}

	if input.FilterStatus != "" {
		if isFiltered {
			filter.Or("quotes.status = ?", input.FilterStatus)
		} else {
			filter.Where("quotes.status = ?", input.FilterStatus)
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

func (s *storage) GetBySingle(input *model.Base) (output *model.Table, err error) {
	query := s.db.Model(&model.Table{}).
		Preload("QuoteProducts", "is_deleted = ?", false).
		Preload("QuoteProducts.Products.CreatedByUsers").
		Preload("QuoteProducts.Products.UpdatedByUsers").
		Preload(clause.Associations)

	if input.QuoteID != nil {
		query.Where("quote_id = ?", input.QuoteID)
	}

	if input.OpportunityID != nil {
		query.Where("opportunity_id = ?", input.OpportunityID)
	}

	if input.IsFinal != nil {
		query.Where("is_final = ?", input.IsFinal)
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
	if input.QuoteID != nil {
		query.Where("quote_id = ?", input.QuoteID)
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

	if input.Name != nil {
		data["name"] = input.Name
	}

	if input.Status != nil {
		data["status"] = input.Status
	}

	if input.IsSyncing != nil {
		data["is_syncing"] = input.IsSyncing
	}

	if input.IsFinal != nil {
		data["is_final"] = input.IsFinal
	}

	if input.OpportunityID != nil {
		data["opportunity_id"] = input.OpportunityID
	}

	if input.AccountID != nil {
		data["account_id"] = input.AccountID
	}

	if input.ExpirationDate != nil {
		data["expiration_date"] = input.ExpirationDate
	}

	if input.Description != nil {
		data["description"] = input.Description
	}

	if input.Tax != nil {
		data["tax"] = input.Tax
	}

	if input.ShippingAndHandling != nil {
		data["shipping_and_handling"] = input.ShippingAndHandling
	}

	if input.IsDeleted != nil {
		data["is_deleted"] = input.IsDeleted
	}

	if input.UpdatedBy != nil {
		data["updated_by"] = input.UpdatedBy
	}

	if input.QuoteID != nil {
		query.Where("quote_id = ?", input.QuoteID)
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
	if input.QuoteID != nil {
		query.Where("quote_id = ?", input.QuoteID)
	}

	err = query.UpdateColumn("is_deleted", true).Delete(&model.Table{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
