package contact

import (
	"encoding/json"

	model "app.eirc/internal/entity/postgresql/db/contacts"
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
	query := s.db.Model(&model.Table{}).Count(&quantity).
		Joins("Salespeople", s.db.Where(`"Salespeople".is_deleted= ?`, false)).
		Joins("Accounts", s.db.Where(`"Accounts".is_deleted= ?`, false)).
		Preload(clause.Associations)

	if input.ContactID != nil {
		query.Where("contact_id = ?", input.ContactID)
	}

	if input.IsDeleted != nil {
		query.Where("contacts.is_deleted = ?", input.IsDeleted)
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
	if input.FilterName != "" {
		filter.Where("contacts.name like ?", "%"+input.FilterName+"%")
		isFiltered = true
	}

	if input.FilterAccountName != "" {
		if isFiltered {
			filter.Or(`"Accounts".name like ?`, "%"+input.FilterAccountName+"%")
		} else {
			filter.Where(`"Accounts".name like ?`, "%"+input.FilterAccountName+"%")
		}
	}

	if input.FilterEmail != "" {
		if isFiltered {
			filter.Or("contacts.email like ?", "%"+input.FilterEmail+"%")
		} else {
			filter.Where("contacts.email like ?", "%"+input.FilterEmail+"%")
		}
	}

	if input.FilterCellPhone != "" {
		if isFiltered {
			filter.Or("contacts.cell_phone like ?", "%"+input.FilterCellPhone+"%")
		} else {
			filter.Where("contacts.cell_phone like ?", "%"+input.FilterCellPhone+"%")
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
	if input.ContactID != nil {
		query.Where("contact_id = ?", input.ContactID)
	}

	if input.AccountID != nil {
		query.Where("account_id = ?", input.AccountID)
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
	if input.ContactID != nil {
		query.Where("contact_id = ?", input.ContactID)
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
	if input.ContactID != nil {
		query.Where("contact_id = ?", input.ContactID)
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

	if input.Title != nil {
		data["title"] = input.Title
	}

	if input.PhoneNumber != nil {
		data["phone_number"] = input.PhoneNumber
	}

	if input.CellPhone != nil {
		data["cell_phone"] = input.CellPhone
	}

	if input.Email != nil {
		data["email"] = input.Email
	}

	if input.Salutation != nil {
		data["salutation"] = input.Salutation
	}

	if input.Department != nil {
		data["department"] = input.Department
	}

	if input.SupervisorID != nil {
		data["supervisor_id"] = input.SupervisorID
	} else {
		data["supervisor_id"] = nil
	}

	if input.AccountID != nil {
		data["account_id"] = input.AccountID
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

	if input.ContactID != nil {
		query.Where("contact_id = ?", input.ContactID)
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
	if input.ContactID != nil {
		query.Where("contact_id = ?", input.ContactID)
	}

	err = query.UpdateColumn("is_deleted", true).Delete(&model.Table{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
