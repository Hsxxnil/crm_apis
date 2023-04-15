package login

import (
	"encoding/json"
	"errors"

	"app.eirc/config"
	jwxModel "app.eirc/internal/interactor/models/jwx"
	loginsModel "app.eirc/internal/interactor/models/logins"
	usersModel "app.eirc/internal/interactor/models/users"
	"app.eirc/internal/interactor/pkg/jwx"
	"app.eirc/internal/interactor/pkg/util"
	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
	jwxService "app.eirc/internal/interactor/service/jwx"
	userService "app.eirc/internal/interactor/service/user"
	"gorm.io/gorm"
)

type Manager interface {
	Login(input *loginsModel.Login) interface{}
	Refresh(input *jwxModel.Refresh) interface{}
}

type manager struct {
	UserService userService.Service
	JwxService  jwxService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		UserService: userService.Init(db),
		JwxService:  jwxService.Init(),
	}
}

func (r *manager) Login(input *loginsModel.Login) interface{} {
	acknowledge, fields, err := r.UserService.AcknowledgeUser(&usersModel.Field{
		UserName:  util.PointerString(input.UserName),
		Password:  util.PointerString(input.Password),
		CompanyID: util.PointerString(input.CompanyID),
	})
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	if acknowledge == false {
		return code.GetCodeMessage(code.PermissionDenied, "Incorrect username or password")
	}

	output := &jwxModel.Token{}
	accessToken, err := r.JwxService.CreateAccessToken(&jwxModel.JWT{
		UserID:    fields[0].UserID,
		CompanyID: fields[0].CompanyID,
		Name:      fields[0].Name,
	})

	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	accessTokenByte, _ := json.Marshal(accessToken)
	err = json.Unmarshal(accessTokenByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	refreshToken, err := r.JwxService.CreateRefreshToken(&jwxModel.JWT{
		UserID: fields[0].UserID,
	})

	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	refreshTokenByte, _ := json.Marshal(refreshToken)
	err = json.Unmarshal(refreshTokenByte, &output)
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	return code.GetCodeMessage(code.Successful, output)
}

func (r *manager) Refresh(input *jwxModel.Refresh) interface{} {

	publicKey := config.RefreshPublicKey

	j := &jwx.JWT{
		PublicKey: publicKey,
		Token:     input.RefreshToken,
	}

	if len(input.RefreshToken) == 0 {
		return code.GetCodeMessage(code.JWTRejected, "RefreshToken is error")
	}

	j, err := j.Verify()
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.JWTRejected, "RefreshToken is error")
	}

	field, err := r.UserService.GetBySingle(&usersModel.Field{
		UserID:    j.Other["user_id"].(string),
		IsDeleted: util.PointerBool(false),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.GetCodeMessage(code.JWTRejected, "RefreshToken is error")
		}

		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err)
	}

	token, err := r.JwxService.CreateAccessToken(&jwxModel.JWT{
		UserID:    field.UserID,
		CompanyID: field.CompanyID,
		Name:      field.Name,
	})
	if err != nil {
		log.Error(err)
		return code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	token.RefreshToken = input.RefreshToken
	return code.GetCodeMessage(code.Successful, token)
}
