package product

import (
	"encoding/json"

	model "crm/internal/entity/postgresql/db/products"
	"crm/internal/interactor/pkg/util/log"

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
	query := s.db.Model(&model.Table{}).Preload(clause.Associations)
	if input.ProductID != nil {
		query.Where("product_id = ?", input.ProductID)
	}

	if input.Sort.Field != "" && input.Sort.Direction != "" {
		query.Order(input.Sort.Field + " " + input.Sort.Direction)
	}

	// filter
	isFiltered := false
	filter := s.db.Model(&model.Table{})
	if input.FilterName != "" {
		filter.Where("name like ?", "%"+input.FilterName+"%")
		isFiltered = true
	}

	if input.FilterCode != "" {
		if isFiltered {
			filter.Or("code like ?", "%"+input.FilterCode+"%")
		} else {
			filter.Where("code like ?", "%"+input.FilterCode+"%")
		}
	}

	if input.FilterDescription != "" {
		if isFiltered {
			filter.Or("description like ?", "%"+input.FilterDescription+"%")
		} else {
			filter.Where("description like ?", "%"+input.FilterDescription+"%")
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
	query := s.db.Model(&model.Table{}).Preload(clause.Associations)
	if input.ProductID != nil {
		query.Where("product_id = ?", input.ProductID)
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
	if input.ProductID != nil {
		query.Where("product_id = ?", input.ProductID)
	}

	if input.Code != nil {
		query.Where("code = ?", input.Code)
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

	if input.Code != nil {
		data["code"] = input.Code
	}

	if input.IsEnable != nil {
		data["is_enable"] = input.IsEnable
	}

	if input.Description != nil {
		data["description"] = input.Description
	}

	if input.Price != nil {
		data["price"] = input.Price
	}

	if input.UpdatedBy != nil {
		data["updated_by"] = input.UpdatedBy
	}

	if input.ProductID != nil {
		query.Where("product_id = ?", input.ProductID)
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
	if input.ProductID != nil {
		query.Where("product_id = ?", input.ProductID)
	}

	err = query.Delete(&model.Table{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
