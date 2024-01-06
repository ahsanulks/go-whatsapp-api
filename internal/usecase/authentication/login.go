package authentication

import (
	"app/internal/port/driven"
	"context"
	"time"

	"github.com/go-playground/validator/v10"
)

// var globalChannel []

type LoginUsecase struct {
	validator *validator.Validate

	deviceProvider driven.DeviceProvider
}

func NewLoginUsecase(validation *validator.Validate, deviceProvider driven.DeviceProvider) *LoginUsecase {
	return &LoginUsecase{
		deviceProvider: deviceProvider,
		validator:      validation,
	}
}

func (lu *LoginUsecase) LoginWithQR(ctx context.Context, params *LoginQRParam) (*LoginQRResponse, error) {
	err := lu.validator.Struct(params)
	if err != nil {
		return nil, err
	}

	resp, err := lu.deviceProvider.NewQRCodeSession(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	return &LoginQRResponse{
		QrCode:    resp.QRCode,
		ExpiredAt: time.Now().Add(time.Second * time.Duration(resp.Timeout)).Unix(),
	}, nil
}
