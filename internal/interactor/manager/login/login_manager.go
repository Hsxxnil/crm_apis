package login

import (
	"encoding/json"
	"errors"

	"crm/config"

	jwxModel "crm/internal/interactor/models/jwx"
	loginsModel "crm/internal/interactor/models/logins"
	usersModel "crm/internal/interactor/models/users"
	"crm/internal/interactor/pkg/jwx"
	"crm/internal/interactor/pkg/util"
	"crm/internal/interactor/pkg/util/code"
	"crm/internal/interactor/pkg/util/log"
	jwxService "crm/internal/interactor/service/jwx"
	userService "crm/internal/interactor/service/user"

	"gorm.io/gorm"
)

type Manager interface {
	Login(input *loginsModel.Login) (int, any)
	Refresh(input *jwxModel.Refresh) (int, any)
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

func (r *manager) Login(input *loginsModel.Login) (int, any) {
	// 驗證帳密
	acknowledge, fields, err := r.UserService.AcknowledgeUser(&usersModel.Field{
		UserName:  util.PointerString(input.UserName),
		Password:  util.PointerString(input.Password),
		CompanyID: util.PointerString(input.CompanyID),
	})
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	if acknowledge == false {
		return code.PermissionDenied, code.GetCodeMessage(code.PermissionDenied, "Incorrect username or password.")
	}

	// 產生accessToken
	output := &jwxModel.Token{}
	accessToken, err := r.JwxService.CreateAccessToken(&jwxModel.JWX{
		UserID:    fields[0].UserID,
		CompanyID: util.PointerString(input.CompanyID),
		Name:      fields[0].Name,
		RoleID:    fields[0].RoleID,
	})

	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	accessTokenByte, _ := json.Marshal(accessToken)
	err = json.Unmarshal(accessTokenByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 產生refreshToken
	refreshToken, err := r.JwxService.CreateRefreshToken(&jwxModel.JWX{
		UserID: fields[0].UserID,
	})

	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	refreshTokenByte, _ := json.Marshal(refreshToken)
	err = json.Unmarshal(refreshTokenByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (r *manager) Refresh(input *jwxModel.Refresh) (int, any) {
	// 驗證refreshToken
	j := &jwx.JWT{
		PublicKey: config.RefreshPublicKey,
		Token:     input.RefreshToken,
	}

	if len(input.RefreshToken) == 0 {
		return code.JWTRejected, code.GetCodeMessage(code.JWTRejected, "RefreshToken is null.")
	}

	j, err := j.Verify()
	if err != nil {
		log.Error(err)
		return code.JWTRejected, code.GetCodeMessage(code.JWTRejected, "RefreshToken is error.")
	}

	field, err := r.UserService.GetBySingle(&usersModel.Field{
		UserID: j.Other["user_id"].(string),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.JWTRejected, code.GetCodeMessage(code.JWTRejected, "RefreshToken is error.")
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 產生accessToken
	token, err := r.JwxService.CreateAccessToken(&jwxModel.JWX{
		UserID:    field.UserID,
		CompanyID: field.CompanyID,
		Name:      field.Name,
		RoleID:    field.RoleID,
	})
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	token.RefreshToken = input.RefreshToken
	return code.Successful, code.GetCodeMessage(code.Successful, token)
}
