package driver

import (
	"app/internal/usecase/authentication"
	"context"
)

type AuthenticationUsecase interface {
	Login(ctx context.Context, params *authentication.LoginParam) (*authentication.LoginResponse, error)
}
