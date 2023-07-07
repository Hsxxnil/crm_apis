package quote_product

import (
	"encoding/json"

	model "app.eirc/internal/entity/postgresql/db/quote_products"
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
	GetByLastCode(input *model.Base) (output *model.Table, err error)
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
	if input.QuoteProductID != nil {
		query.Where("quote_product_id = ?", input.QuoteProductID)
	}

	if input.QuoteID != nil {
		query.Where("quote_id = ?", input.QuoteID)
	}

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
	if input.QuoteProductID != nil {
		query.Where("quote_product_id = ?", input.QuoteProductID)
	}

	if input.QuoteID != nil {
		query.Where("quote_id = ?", input.QuoteID)
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
	if input.QuoteProductID != nil {
		query.Where("quote_product_id = ?", input.QuoteProductID)
	}

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
	if input.QuoteProductID != nil {
		query.Where("quote_product_id = ?", input.QuoteProductID)
	}

	if input.QuoteID != nil {
		query.Where("quote_id = ?", input.QuoteID)
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

	if input.ProductID != nil {
		data["product_id"] = input.ProductID
	}

	if input.Quantity != nil {
		data["quantity"] = input.Quantity
	}

	if input.UnitPrice != nil {
		data["unit_price"] = input.UnitPrice
	}

	if input.SubTotal != nil {
		data["sub_total"] = input.SubTotal
	}

	if input.TotalPrice != nil {
		data["total_price"] = input.TotalPrice
	}

	if input.Discount != nil {
		data["discount"] = input.Discount
	}

	if input.Description != nil {
		data["description"] = input.Description
	}

	if input.UpdatedBy != nil {
		data["updated_by"] = input.UpdatedBy
	}

	if input.QuoteProductID != nil {
		query.Where("quote_product_id = ?", input.QuoteProductID)
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
	if input.QuoteProductID != nil {
		query.Where("quote_product_id = ?", input.QuoteProductID)
	}

	err = query.Delete(&model.Table{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *storage) GetByLastCode(input *model.Base) (output *model.Table, err error) {
	query := s.db.Model(&model.Table{}).Preload(clause.Associations)
	if input.QuoteID != nil {
		query.Where("quote_id = ?", input.QuoteID)
	}

	err = query.Order("code desc").First(&output).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}
