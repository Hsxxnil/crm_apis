package account

import (
	"encoding/json"

	"github.com/lib/pq"

	model "crm/internal/entity/postgresql/db/accounts"
	"crm/internal/interactor/pkg/util/log"

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
	query := s.db.Model(&model.Table{}).Count(&quantity).Joins("Salespeople").Preload(clause.Associations)

	if input.AccountID != nil {
		query.Where("account_id = ?", input.AccountID)
	}
	if input.Sort.Field != "" && input.Sort.Direction != "" {
		if input.Sort.Field == "salesperson_name" && input.Sort.Direction != "" {
			query.Order(`"Salespeople".name` + " " + input.Sort.Direction)
		} else {
			query.Order(input.Sort.Field + " " + input.Sort.Direction)
		}
	}

	// filter
	isFiltered := false
	filter := s.db.Model(&model.Table{})
	if input.FilterName != "" {
		filter.Where("accounts.name like ?", "%"+input.FilterName+"%")
		isFiltered = true
	}

	if input.FilterType != nil {
		if isFiltered {
			filter.Or("array[accounts.type] && ?", pq.Array(input.FilterType))
		} else {
			filter.Where("array[accounts.type] && ?", pq.Array(input.FilterType))
		}
	}

	if input.FilterPhoneNumber != "" {
		if isFiltered {
			filter.Or("accounts.phone_number like ?", "%"+input.FilterPhoneNumber+"%")
		} else {
			filter.Where("accounts.phone_number like ?", "%"+input.FilterPhoneNumber+"%")
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
	if input.AccountID != nil {
		query.Where("account_id = ?", input.AccountID)
	}

	// filter
	isFiltered := false
	filter := s.db.Model(&model.Table{})
	if input.FilterName != "" {
		filter.Where("name like ?", "%"+input.FilterName+"%")
		isFiltered = true
	}

	if input.FilterType != nil {
		if isFiltered {
			filter.Or("array[accounts.type] && ?", pq.Array(input.FilterType))
		} else {
			filter.Where("array[accounts.type] && ?", pq.Array(input.FilterType))
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
	query := s.db.Model(&model.Table{}).Preload(clause.Associations)
	if input.AccountID != nil {
		query.Where("account_id = ?", input.AccountID)
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
	if input.AccountID != nil {
		query.Where("account_id = ?", input.AccountID)
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

	if input.PhoneNumber != nil {
		data["phone_number"] = input.PhoneNumber
	}

	if input.Type != nil {
		data["type"] = input.Type
	}

	if input.IndustryID != nil {
		data["industry_id"] = input.IndustryID
	} else {
		data["industry_id"] = nil
	}

	if input.ParentAccountID != nil {
		data["parent_account_id"] = input.ParentAccountID
	} else {
		data["parent_account_id"] = nil
	}

	if input.SalespersonID != nil {
		data["salesperson_id"] = input.SalespersonID
	}

	if input.UpdatedBy != nil {
		data["updated_by"] = input.UpdatedBy
	}

	if input.AccountID != nil {
		query.Where("account_id = ?", input.AccountID)
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
	if input.AccountID != nil {
		query.Where("account_id = ?", input.AccountID)
	}

	err = query.Delete(&model.Table{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
