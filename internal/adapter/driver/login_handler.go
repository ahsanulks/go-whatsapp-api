package driver

import (
	v1 "app/api/v1"
	"app/internal/port/driver"
	"app/internal/usecase/authentication"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type LoginHandler struct {
	v1.UnimplementedAuthenticationServer
	auth driver.AuthenticationUsecase
}

func NewLoginHandler(auth driver.AuthenticationUsecase) *LoginHandler {
	return &LoginHandler{
		auth: auth,
	}
}

func (l *LoginHandler) Login(ctx context.Context, params *v1.LoginAuthenticationRequest) (*v1.LoginAuthenticationResponse, error) {
	resp, err := l.auth.Login(ctx, &authentication.LoginParam{
		ID: params.Id,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &v1.LoginAuthenticationResponse{
		QrCode:  resp.QrCode,
		Timeout: int32(resp.Timeout),
	}, nil
}
