package contact

import (
	"encoding/json"
	"errors"

	accountContactModel "app.eirc/internal/interactor/models/account_contacts"
	accountModel "app.eirc/internal/interactor/models/accounts"
	historicalRecordModel "app.eirc/internal/interactor/models/historical_records"
	userModel "app.eirc/internal/interactor/models/users"
	accountService "app.eirc/internal/interactor/service/account"
	accountContactService "app.eirc/internal/interactor/service/account_contact"
	historicalRecordService "app.eirc/internal/interactor/service/historical_record"
	userService "app.eirc/internal/interactor/service/user"

	"app.eirc/internal/interactor/pkg/util"

	contactModel "app.eirc/internal/interactor/models/contacts"
	contactService "app.eirc/internal/interactor/service/contact"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *contactModel.Create) (int, interface{})
	GetByList(input *contactModel.Fields) (int, interface{})
	GetBySingle(input *contactModel.Field) (int, interface{})
	Delete(input *contactModel.Field) (int, interface{})
	Update(trx *gorm.DB, input *contactModel.Update) (int, interface{})
}

type manager struct {
	ContactService          contactService.Service
	AccountContactService   accountContactService.Service
	HistoricalRecordService historicalRecordService.Service
	UserService             userService.Service
	AccountService          accountService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		ContactService:          contactService.Init(db),
		AccountContactService:   accountContactService.Init(db),
		HistoricalRecordService: historicalRecordService.Init(db),
		UserService:             userService.Init(db),
		AccountService:          accountService.Init(db),
	}
}

const sourceType = "聯絡人"

func (m *manager) Create(trx *gorm.DB, input *contactModel.Create) (int, interface{}) {
	defer trx.Rollback()

	contactBase, err := m.ContactService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步新增帳戶聯絡人關聯
	_, err = m.AccountContactService.WithTrx(trx).Create(&accountContactModel.Create{
		AccountID: input.AccountID,
		ContactID: *contactBase.ContactID,
		CreatedBy: *contactBase.CreatedBy,
	})
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步新增聯絡人歷程記錄
	_, err = m.HistoricalRecordService.WithTrx(trx).Create(&historicalRecordModel.Create{
		SourceID:   *contactBase.ContactID,
		Action:     "建立",
		SourceType: sourceType,
		ModifiedBy: *contactBase.CreatedBy,
	})
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, contactBase.ContactID)
}

func (m *manager) GetByList(input *contactModel.Fields) (int, interface{}) {
	input.IsDeleted = util.PointerBool(false)
	output := &contactModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, contactBase, err := m.ContactService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	contactByte, err := json.Marshal(contactBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(contactByte, &output.Contacts)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, contacts := range output.Contacts {
		contacts.CreatedBy = *contactBase[i].CreatedByUsers.Name
		contacts.UpdatedBy = *contactBase[i].UpdatedByUsers.Name
		contacts.SalespersonName = *contactBase[i].Salespeople.Name
		contacts.AccountName = *contactBase[i].Accounts.Name
		if contacts.SupervisorID != "" {
			supervisorBase, err := m.ContactService.GetBySingle(&contactModel.Field{
				ContactID: contacts.SupervisorID,
				IsDeleted: util.PointerBool(false),
			})
			if err != nil {
				contacts.SupervisorName = ""
			} else {
				contacts.SupervisorName = *supervisorBase.Name
			}
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *contactModel.Field) (int, interface{}) {
	input.IsDeleted = util.PointerBool(false)
	contactBase, err := m.ContactService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &contactModel.Single{}
	contactByte, _ := json.Marshal(contactBase)
	err = json.Unmarshal(contactByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.CreatedBy = *contactBase.CreatedByUsers.Name
	output.UpdatedBy = *contactBase.UpdatedByUsers.Name
	output.SalespersonName = *contactBase.Salespeople.Name
	output.AccountName = *contactBase.Accounts.Name
	if contactBase.SupervisorID != nil {
		supervisorBase, err := m.ContactService.GetBySingle(&contactModel.Field{
			ContactID: *contactBase.SupervisorID,
			IsDeleted: util.PointerBool(false),
		})
		if err != nil {
			output.SupervisorName = ""
		} else {
			output.SupervisorName = *supervisorBase.Name
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *contactModel.Field) (int, interface{}) {
	_, err := m.ContactService.GetBySingle(&contactModel.Field{
		ContactID: input.ContactID,
		IsDeleted: util.PointerBool(false),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.ContactService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步刪除帳戶聯絡人關聯
	accountContactBase, _ := m.AccountContactService.GetBySingle(&accountContactModel.Field{
		ContactID: util.PointerString(input.ContactID),
		IsDeleted: util.PointerBool(false),
	})
	err = m.AccountContactService.Delete(&accountContactModel.Field{
		AccountContactID: *accountContactBase.AccountContactID,
	})
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(trx *gorm.DB, input *contactModel.Update) (int, interface{}) {
	defer trx.Rollback()

	contactBase, err := m.ContactService.GetBySingle(&contactModel.Field{
		ContactID: input.ContactID,
		IsDeleted: util.PointerBool(false),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.ContactService.WithTrx(trx).Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步修改帳戶聯絡人關聯
	if input.AccountID != nil && *input.AccountID != *contactBase.AccountID {
		accountContactBase, err := m.AccountContactService.GetBySingle(&accountContactModel.Field{
			ContactID: util.PointerString(input.ContactID),
			IsDeleted: util.PointerBool(false),
		})
		err = m.AccountContactService.WithTrx(trx).Update(&accountContactModel.Update{
			AccountContactID: *accountContactBase.AccountContactID,
			AccountID:        input.AccountID,
			UpdatedBy:        input.UpdatedBy,
		})
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	// 同步新增聯絡人歷程記錄
	var records []historicalRecordModel.AddHistoricalRecord

	if input.Name != nil && *input.Name != *contactBase.Name {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "名稱為",
			Values: *input.Name,
		})
	}

	if input.Title != nil && *input.Title != *contactBase.Title {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "職稱為",
			Values: *input.Title,
		})
	}

	if input.PhoneNumber != nil && *input.PhoneNumber != *contactBase.PhoneNumber {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "電話為",
			Values: *input.PhoneNumber,
		})
	}

	if input.CellPhone != nil && *input.CellPhone != *contactBase.CellPhone {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "行動電話為",
			Values: *input.CellPhone,
		})
	}

	if input.Email != nil && *input.Email != *contactBase.Email {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "電子郵件為",
			Values: *input.Email,
		})
	}

	if input.Salutation != nil && *input.Salutation != *contactBase.Salutation {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "稱謂為",
			Values: *input.Salutation,
		})
	}

	if input.Department != nil && *input.Department != *contactBase.Department {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "部門為",
			Values: *input.Department,
		})
	}

	if input.SupervisorID != nil && *input.SupervisorID != *contactBase.SupervisorID {
		supervisorBase, _ := m.ContactService.GetBySingle(&contactModel.Field{
			ContactID: *input.SupervisorID,
			IsDeleted: util.PointerBool(false),
		})
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "直屬上司為",
			Values: *supervisorBase.Name,
		})
	}

	if input.AccountID != nil && *input.AccountID != *contactBase.AccountID {
		accountBase, _ := m.AccountService.GetBySingle(&accountModel.Field{
			AccountID: *input.AccountID,
			IsDeleted: util.PointerBool(false),
		})
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "帳戶為",
			Values: *accountBase.Name,
		})
	}

	if input.SalespersonID != nil && *input.SalespersonID != *contactBase.SalespersonID {
		salespersonBase, _ := m.UserService.GetBySingle(&userModel.Field{
			UserID:    *input.SalespersonID,
			IsDeleted: util.PointerBool(false),
		})
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "業務員為",
			Values: *salespersonBase.Name,
		})
	}

	for _, record := range records {
		_, err = m.HistoricalRecordService.WithTrx(trx).Create(&historicalRecordModel.Create{
			SourceID:   *contactBase.ContactID,
			Action:     "修改",
			SourceType: sourceType,
			Field:      record.Fields,
			Value:      record.Values,
			ModifiedBy: *input.UpdatedBy,
		})
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, contactBase.ContactID)
}
