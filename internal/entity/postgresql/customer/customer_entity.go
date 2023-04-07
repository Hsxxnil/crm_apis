package customer

import (
	"encoding/json"

	model "app.eirc/internal/entity/postgresql/db/customers"
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
	query := s.db.Model(&model.Table{}).Preload(clause.Associations)
	if input.CID != nil {
		query.Where("c_id = ?", input.CID)
	}

	err = query.Count(&quantity).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Find(&output).Error
	if err != nil {
		log.Error(err)
		return 0, nil, err
	}

	return quantity, output, nil
}

func (s *storage) GetBySingle(input *model.Base) (output *model.Table, err error) {
	query := s.db.Model(&model.Table{}).Preload(clause.Associations)
	if input.CID != nil {
		query.Where("c_id = ?", input.CID)
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
	if input.CID != nil {
		query.Where("c_id = ?", input.CID)
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

	if input.ShortName != nil {
		data["short_name"] = input.ShortName
	}

	if input.EngName != nil {
		data["eng_name"] = input.EngName
	}

	if input.Name != nil {
		data["name"] = input.Name
	}

	if input.ZipCode != nil {
		data["zip_code"] = input.ZipCode
	}

	if input.Address != nil {
		data["address"] = input.Address
	}

	if input.Tel != nil {
		data["tel"] = input.Tel
	}
	if input.Fax != nil {
		data["fax"] = input.Fax
	}
	if input.Map != nil {
		data["map"] = input.Map
	}
	if input.Liaison != nil {
		data["liaison"] = input.Liaison
	}
	if input.Mail != nil {
		data["mail"] = input.Mail
	}
	if input.LiaisonPhone != nil {
		data["liaison_phone"] = input.LiaisonPhone
	}
	if input.TaxIdNumber != nil {
		data["tax_id_number"] = input.TaxIdNumber
	}
	if input.Remark != nil {
		data["remark"] = input.Remark
	}

	if input.CID != nil {
		query.Where("c_id = ?", input.CID)
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
	if input.CID != nil {
		query.Where("c_id = ?", input.CID)
	}

	err = query.Delete(&model.Table{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
