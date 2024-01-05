package driver

import (
	"app/internal/usecase/authentication"
	"context"
)

type AuthenticationUsecase interface {
	LoginWithQR(ctx context.Context, params *authentication.LoginQRParam) (*authentication.LoginQRResponse, error)
}
