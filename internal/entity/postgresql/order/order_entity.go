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
	GetByListNoQuantity(input *model.Base) (output []*model.Table, err error)
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
	query := s.db.Model(&model.Table{}).Count(&quantity).Joins("Accounts").Joins("Contracts").Preload(clause.Associations)
	if input.OrderID != nil {
		query.Where("order_id = ?", input.OrderID)
	}

	if input.ContractID != nil {
		query.Where("orders.contract_id = ?", input.ContractID)
	}

	if input.Sort.Field != "" && input.Sort.Direction != "" {
		if input.Sort.Field == "account_name" && input.Sort.Direction != "" {
			query.Order(`"Accounts".name` + " " + input.Sort.Direction)
		} else if input.Sort.Field == "contract_code" && input.Sort.Direction != "" {
			query.Order(`"Contracts".code` + " " + input.Sort.Direction)
		} else {
			query.Order(input.Sort.Field + " " + input.Sort.Direction)
		}
	}

	// filter
	isFiltered := false
	filter := s.db.Model(&model.Table{})
	if input.FilterCode != "" {
		filter.Where("orders.code like ?", "%"+input.FilterCode+"%")
		isFiltered = true
	}

	if input.FilterAccountName != "" {
		if isFiltered {
			filter.Or(`"Accounts".name like ?`, "%"+input.FilterAccountName+"%")
		} else {
			filter.Where(`"Accounts".name like ?`, "%"+input.FilterAccountName+"%")
		}
	}

	if input.FilterContractCode != "" {
		if isFiltered {
			filter.Or(`"Contracts".code like ?`, "%"+input.FilterContractCode+"%")
		} else {
			filter.Where(`"Contracts".code like ?`, "%"+input.FilterContractCode+"%")
		}
	}

	if input.FilterStatus != "" {
		if isFiltered {
			filter.Or("orders.status = ?", input.FilterStatus)
		} else {
			filter.Where("orders.status = ?", input.FilterStatus)
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

func (s *storage) GetByListNoQuantity(input *model.Base) (output []*model.Table, err error) {
	query := s.db.Model(&model.Table{}).Preload(clause.Associations)
	if input.OrderID != nil {
		query.Where("order_id = ?", input.OrderID)
	}

	if input.ContractID != nil {
		query.Where("contract_id = ?", input.ContractID)
	}

	err = query.Order("created_at desc").Find(&output).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}

func (s *storage) GetBySingle(input *model.Base) (output *model.Table, err error) {
	query := s.db.Model(&model.Table{}).Preload("OrderProducts.Products.CreatedByUsers").Preload("OrderProducts.Products.UpdatedByUsers").Preload(clause.Associations)
	if input.OrderID != nil {
		query.Where("order_id = ?", input.OrderID)
	}

	if input.ContractID != nil {
		query.Where("orders.contract_id = ?", input.ContractID)
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

	if input.IsDeleted != nil {
		data["is_deleted"] = input.IsDeleted
	}

	if input.UpdatedBy != nil {
		data["updated_by"] = input.UpdatedBy
	}

	if input.ActivatedBy != nil {
		data["activated_by"] = input.ActivatedBy
	} else {
		data["activated_by"] = nil
	}

	if input.ActivatedAt != nil {
		data["activated_at"] = input.ActivatedAt
	} else {
		data["activated_at"] = nil
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
