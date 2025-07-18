package jwx

import (
	"time"

	"crm/config"
	model "crm/internal/interactor/models/jwx"
	"crm/internal/interactor/pkg/jwx"
	"crm/internal/interactor/pkg/util"
	"crm/internal/interactor/pkg/util/log"
)

type Service interface {
	CreateAccessToken(input *model.JWX) (output *model.Token, err error)
	CreateRefreshToken(input *model.JWX) (output *model.Token, err error)
}

type service struct {
}

func Init() Service {
	return &service{}
}

func (s service) CreateAccessToken(input *model.JWX) (output *model.Token, err error) {
	other := map[string]any{
		"company_id": input.CompanyID,
		"user_id":    input.UserID,
		"name":       input.Name,
		"role_id":    input.RoleID,
	}

	accessExpiration := util.NowToUTC().Add(time.Minute * 5).Unix()
	j := &jwx.JWE{
		PublicKey:     config.AccessPublicKey,
		Other:         other,
		ExpirationKey: accessExpiration,
	}

	j, err = j.Create()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	accessToken := j.Token

	output = &model.Token{
		AccessToken: accessToken,
	}

	return output, nil
}

func (s service) CreateRefreshToken(input *model.JWX) (output *model.Token, err error) {
	other := map[string]any{
		"user_id": input.UserID,
	}

	refreshExpiration := util.NowToUTC().Add(time.Hour * 8).Unix()
	j := &jwx.JWT{
		PrivateKey:    config.RefreshPrivateKey,
		Other:         other,
		ExpirationKey: refreshExpiration,
	}

	j, err = j.Create()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	refreshToken := j.Token
	output = &model.Token{
		RefreshToken: refreshToken,
	}

	return output, nil
}
