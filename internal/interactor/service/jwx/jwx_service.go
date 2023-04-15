package jwx

import (
	"time"

	"app.eirc/config"
	model "app.eirc/internal/interactor/models/jwx"
	"app.eirc/internal/interactor/pkg/jwx"
	"app.eirc/internal/interactor/pkg/util"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Service interface {
	CreateAccessToken(input *model.JWT) (output *model.Token, err error)
	CreateRefreshToken(input *model.JWT) (output *model.Token, err error)
}

type service struct {
}

func Init() Service {
	return &service{}
}

func (s service) CreateAccessToken(input *model.JWT) (output *model.Token, err error) {

	other := map[string]interface{}{
		"company_id": input.CompanyID,
		"user_id":    input.UserID,
		"name":       input.Name,
	}

	publicKey := config.AccessPublicKey

	accessExpiration := util.NowToUTC().Add(time.Minute * 5).Unix()
	j := &jwx.JWE{
		PublicKey:     publicKey,
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

func (s service) CreateRefreshToken(input *model.JWT) (output *model.Token, err error) {
	// TODO implement me
	other := map[string]interface{}{
		"user_id": input.UserID,
	}
	privateKey := config.RefreshPrivateKey

	refreshExpiration := util.NowToUTC().Add(time.Hour * 8).Unix()
	j := &jwx.JWT{
		PrivateKey:    privateKey,
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
