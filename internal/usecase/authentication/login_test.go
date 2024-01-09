package authentication

import (
	"app/internal/port/driven"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type fakeDeviceProvider struct{}

func (f *fakeDeviceProvider) NewQRCodeSession(ctx context.Context, id string) (*driven.AuthQRResponse, error) {
	if id == "iderror" {
		return nil, errors.New("error")
	}
	return &driven.AuthQRResponse{
		QRCode:  "123456789",
		Timeout: 60,
	}, nil
}

func TestLoginUsecase_LoginWithQR(t *testing.T) {
	type fields struct {
		validator      *validator.Validate
		deviceProvider driven.DeviceProvider
	}
	type args struct {
		ctx    context.Context
		params *LoginQRParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *LoginQRResponse
		wantErr bool
	}{
		{
			name: "when invalid struct validation should return error",
			fields: fields{
				validator:      validator.New(),
				deviceProvider: new(fakeDeviceProvider),
			},
			args: args{
				ctx:    context.Background(),
				params: &LoginQRParam{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "when call provider return error should return error",
			fields: fields{
				validator:      validator.New(),
				deviceProvider: new(fakeDeviceProvider),
			},
			args: args{
				ctx:    context.Background(),
				params: &LoginQRParam{ID: "iderror"},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				validator:      validator.New(),
				deviceProvider: new(fakeDeviceProvider),
			},
			args: args{
				ctx:    context.Background(),
				params: &LoginQRParam{ID: faker.Username()},
			},
			want: &LoginQRResponse{
				QrCode:    "123456789",
				ExpiredAt: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lu := &LoginUsecase{
				validator:      tt.fields.validator,
				deviceProvider: tt.fields.deviceProvider,
			}
			got, err := lu.LoginWithQR(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginUsecase.LoginWithQR() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			assert.Equal(t, tt.want.QrCode, got.QrCode)
			assert.Greater(t, got.ExpiredAt, time.Now().Unix())
		})
	}
}
