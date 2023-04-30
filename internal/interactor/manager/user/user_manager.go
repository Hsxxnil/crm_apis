package user

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	userModel "app.eirc/internal/interactor/models/users"
	userService "app.eirc/internal/interactor/service/user"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *userModel.Create) interface{}
	GetByList(input *userModel.Fields) interface{}
	GetBySingle(input *userModel.Field) interface{}
	Delete(input *userModel.Update) interface{}
	Update(input *userModel.Update) interface{}
}

type manager struct {
	UserService userService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		UserService: userService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *userModel.Create) interface{} {
	defer trx.Rollback()
	quantity, _ := m.UserService.GetByQuantity(&userModel.Field{
		UserName:  util.PointerString(input.UserName),
		CompanyID: util.PointerString(input.CompanyID),
	})

	if quantity > 0 {
		log.Info("UserName already exists. UserName: ", input.UserName, ",CompanyID:", input.CompanyID)
		return code.GetCodeMessage(code.BadRequest, "user already exists")
	}

	userBase, err := m.UserService.WithTrx(trx).Create(input)
	if err != nil {
		if err.Error() == "user already exists" {
			return code.GetCodeMessage(code.BadRequest, err.Error())
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.GetCodeMessage(code.Successful, userBase.UserID)
}

func (m *manager) GetByList(input *userModel.Fields) interface{} {
	input.IsDeleted = util.PointerBool(false)
	output := &userModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, userBase, err := m.UserService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	userByte, err := json.Marshal(userBase)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(userByte, &output.Users)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *userModel.Field) interface{} {
	input.IsDeleted = util.PointerBool(false)
	userBase, err := m.UserService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	output := &userModel.Single{}
	userByte, _ := json.Marshal(userBase)
	err = json.Unmarshal(userByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *userModel.Update) interface{} {
	_, err := m.UserService.GetBySingle(&userModel.Field{
		UserID:    input.UserID,
		IsDeleted: util.PointerBool(false),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.UserService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *userModel.Update) interface{} {
	userBase, err := m.UserService.GetBySingle(&userModel.Field{
		UserID:    input.UserID,
		IsDeleted: util.PointerBool(false),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.DoesNotExist, err)
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	err = m.UserService.Update(input)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, userBase.UserID)
}
