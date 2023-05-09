package order

import (
	"encoding/json"

	model "app.eirc/internal/entity/postgresql/db/orders"
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
	query := s.db.Model(&model.Table{}).Preload("OrderProducts.Products").Preload(clause.Associations)
	if input.OrderID != nil {
		query.Where("order_id = ?", input.OrderID)
	}

	// filter
	//isFiltered := false
	filterdb := s.db.Model(&model.Table{})
	if *input.FilterCode != "" {
		filterdb.Where("code like ?", "%"+*input.FilterCode+"%")
		//isFiltered = true
	}

	query.Where(filterdb)

	err = query.Count(&quantity).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("created_at desc").Find(&output).Error
	if err != nil {
		log.Error(err)
		return 0, nil, err
	}

	return quantity, output, nil
}

func (s *storage) GetBySingle(input *model.Base) (output *model.Table, err error) {
	query := s.db.Model(&model.Table{}).Preload("OrderProducts.Products").Preload(clause.Associations)
	if input.OrderID != nil {
		query.Where("order_id = ?", input.OrderID)
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
	if input.OrderID != nil {
		query.Where("order_id = ?", input.OrderID)
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

	if input.StartDate != nil {
		data["start_date"] = input.StartDate
	}

	if input.AccountID != nil {
		data["account_id"] = input.AccountID
	}

	if input.ContractID != nil {
		data["contract_id"] = input.ContractID
	}

	if input.Description != nil {
		data["description"] = input.Description
	}

	if input.UpdatedBy != nil {
		data["updated_by"] = input.UpdatedBy
	}

	if input.ActivatedBy != nil {
		data["activated_by"] = input.ActivatedBy
	}

	if input.ActivatedAt != nil {
		data["activated_at"] = input.ActivatedAt
	}

	if input.OrderID != nil {
		query.Where("order_id = ?", input.OrderID)
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
	if input.OrderID != nil {
		query.Where("order_id = ?", input.OrderID)
	}

	err = query.Delete(&model.Table{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
