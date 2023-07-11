package account

import (
	"encoding/json"
	"errors"
	"sort"
	"strings"

	industryModel "app.eirc/internal/interactor/models/industries"
	userModel "app.eirc/internal/interactor/models/users"

	historicalRecordModel "app.eirc/internal/interactor/models/historical_records"

	historicalRecordService "app.eirc/internal/interactor/service/historical_record"
	industryService "app.eirc/internal/interactor/service/industry"
	userService "app.eirc/internal/interactor/service/user"

	contactModel "app.eirc/internal/interactor/models/contacts"
	"app.eirc/internal/interactor/pkg/util"

	accountModel "app.eirc/internal/interactor/models/accounts"
	accountService "app.eirc/internal/interactor/service/account"
	contactService "app.eirc/internal/interactor/service/contact"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *accountModel.Create) (int, interface{})
	GetByList(input *accountModel.Fields) (int, interface{})
	GetBySingle(input *accountModel.Field) (int, interface{})
	GetBySingleContacts(input *accountModel.Field) (int, interface{})
	Delete(input *accountModel.Field) (int, interface{})
	Update(trx *gorm.DB, input *accountModel.Update) (int, interface{})
}

type manager struct {
	AccountService          accountService.Service
	ContactService          contactService.Service
	HistoricalRecordService historicalRecordService.Service
	IndustryService         industryService.Service
	UserService             userService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		AccountService:          accountService.Init(db),
		ContactService:          contactService.Init(db),
		HistoricalRecordService: historicalRecordService.Init(db),
		IndustryService:         industryService.Init(db),
		UserService:             userService.Init(db),
	}
}

const sourceType = "帳戶"

func (m *manager) Create(trx *gorm.DB, input *accountModel.Create) (int, interface{}) {
	defer trx.Rollback()

	// 陣列排序
	sort.Strings(input.Type)
	accountBase, err := m.AccountService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步新增帳戶歷程記錄
	_, err = m.HistoricalRecordService.WithTrx(trx).Create(&historicalRecordModel.Create{
		SourceID:   *accountBase.AccountID,
		Action:     "建立",
		Content:    sourceType,
		ModifiedBy: *accountBase.CreatedBy,
	})
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, accountBase.AccountID)
}

func (m *manager) GetByList(input *accountModel.Fields) (int, interface{}) {
	output := &accountModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, accountBase, err := m.AccountService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	accountByte, err := json.Marshal(accountBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(accountByte, &output.Accounts)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, accounts := range output.Accounts {
		accounts.IndustryName = *accountBase[i].Industries.Name
		accounts.CreatedBy = *accountBase[i].CreatedByUsers.Name
		accounts.UpdatedBy = *accountBase[i].UpdatedByUsers.Name
		accounts.SalespersonName = *accountBase[i].Salespeople.Name
		parentAccountsBase, err := m.AccountService.GetBySingle(&accountModel.Field{
			AccountID: accounts.ParentAccountID,
		})
		if err != nil {
			accounts.ParentAccountName = ""
		} else {
			accounts.ParentAccountName = *parentAccountsBase.Name
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *accountModel.Field) (int, interface{}) {
	accountBase, err := m.AccountService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &accountModel.Single{}
	accountByte, _ := json.Marshal(accountBase)
	err = json.Unmarshal(accountByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.IndustryName = *accountBase.Industries.Name
	output.CreatedBy = *accountBase.CreatedByUsers.Name
	output.UpdatedBy = *accountBase.UpdatedByUsers.Name
	output.SalespersonName = *accountBase.Salespeople.Name
	parentAccountsBase, err := m.AccountService.GetBySingle(&accountModel.Field{
		AccountID: *accountBase.ParentAccountID,
	})
	if err != nil {
		output.ParentAccountName = ""
	} else {
		output.ParentAccountName = *parentAccountsBase.Name
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingleContacts(input *accountModel.Field) (int, interface{}) {
	accountBase, err := m.AccountService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &accountModel.SingleContacts{}
	accountByte, _ := json.Marshal(accountBase)
	err = json.Unmarshal(accountByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.IndustryName = *accountBase.Industries.Name
	output.CreatedBy = *accountBase.CreatedByUsers.Name
	output.UpdatedBy = *accountBase.UpdatedByUsers.Name
	output.SalespersonName = *accountBase.Salespeople.Name
	parentAccountsBase, err := m.AccountService.GetBySingle(&accountModel.Field{
		AccountID: *accountBase.ParentAccountID,
	})
	if err != nil {
		output.ParentAccountName = ""
	} else {
		output.ParentAccountName = *parentAccountsBase.Name
	}
	for i, contacts := range output.AccountContacts {
		contactBase, _ := m.ContactService.GetBySingle(&contactModel.Field{
			ContactID: contacts.ContactID,
		})
		output.AccountContacts[i].ContactName = *contactBase.Name
		output.AccountContacts[i].ContactTitle = *contactBase.Title
		output.AccountContacts[i].ContactPhoneNumber = *contactBase.PhoneNumber
		output.AccountContacts[i].ContactCellPhone = *contactBase.CellPhone
		output.AccountContacts[i].ContactEmail = *contactBase.Email
		output.AccountContacts[i].ContactSalutation = *contactBase.Salutation
		output.AccountContacts[i].ContactDepartment = *contactBase.Department
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *accountModel.Field) (int, interface{}) {
	_, err := m.AccountService.GetBySingle(&accountModel.Field{
		AccountID: input.AccountID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.AccountService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(trx *gorm.DB, input *accountModel.Update) (int, interface{}) {
	defer trx.Rollback()

	accountBase, err := m.AccountService.GetBySingle(&accountModel.Field{
		AccountID: input.AccountID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 陣列排序
	sort.Strings(*input.Type)
	err = m.AccountService.WithTrx(trx).Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步新增帳戶歷程記錄
	var records []historicalRecordModel.AddHistoricalRecord

	if input.Name != nil && *input.Name != *accountBase.Name {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "名稱",
			Values: "為" + *input.Name,
		})
	}

	// 比對帳戶類型是否變更
	var inputType string
	for i, value := range *input.Type {
		if value != (*accountBase.Type)[i] {
			inputType = strings.Join(*input.Type, "、")
			break
		}
	}

	if inputType != "" {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "類型",
			Values: "為" + inputType,
		})
	}

	if input.PhoneNumber != nil && *input.PhoneNumber != *accountBase.PhoneNumber {
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "電話號碼",
			Values: "為" + *input.PhoneNumber,
		})
	}

	if input.IndustryID != nil && *input.IndustryID != *accountBase.IndustryID {
		industryBase, _ := m.IndustryService.GetBySingle(&industryModel.Field{
			IndustryID: *input.IndustryID,
		})
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "行業",
			Values: "為" + *industryBase.Name,
		})
	}

	if input.ParentAccountID != nil && *input.ParentAccountID != *accountBase.ParentAccountID {
		parentAccountBase, _ := m.AccountService.GetBySingle(&accountModel.Field{
			AccountID: input.AccountID,
		})
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "父系帳戶",
			Values: "為" + *parentAccountBase.Name,
		})
	}

	if input.SalespersonID != nil && *input.SalespersonID != *accountBase.SalespersonID {
		salespersonBase, _ := m.UserService.GetBySingle(&userModel.Field{
			UserID: *input.SalespersonID,
		})
		records = append(records, historicalRecordModel.AddHistoricalRecord{
			Fields: "業務員",
			Values: "為" + *salespersonBase.Name,
		})
	}

	for _, record := range records {
		_, err = m.HistoricalRecordService.WithTrx(trx).Create(&historicalRecordModel.Create{
			SourceID:   *accountBase.AccountID,
			Action:     "修改",
			Content:    sourceType + record.Fields + record.Values,
			ModifiedBy: *input.UpdatedBy,
		})
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, accountBase.AccountID)
}
