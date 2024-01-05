package authentication

import (
	"app/internal/port/driven"
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"go.mau.fi/whatsmeow"
)

type LoginUsecase struct {
	validator *validator.Validate
	session   map[string]*whatsmeow.Client

	deviceProvider driven.DeviceProvider
}

func NewLoginUsecase(validation *validator.Validate, deviceProvider driven.DeviceProvider) *LoginUsecase {
	return &LoginUsecase{
		deviceProvider: deviceProvider,
		validator:      validation,
		session:        make(map[string]*whatsmeow.Client),
	}
}

func (lu *LoginUsecase) Login(ctx context.Context, params *LoginParam) (*LoginResponse, error) {
	err := lu.validator.Struct(params)
	if err != nil {
		return nil, err
	}

	if _, exist := lu.session[params.ID]; exist {
		return nil, errors.New("cannot initialize session, becuase session already exists")
	}

	device := lu.deviceProvider.NewDeviceSession()
	client := whatsmeow.NewClient(device, nil)
	// TODO: change this line when user already logged in
	lu.session[params.ID] = client

	qrChannel, _ := client.GetQRChannel(ctx)

	// initialize websocket
	err = client.Connect()
	if err != nil {
		return nil, errors.New("failed to connect whatsapp")
	}

	qrCode, timeout := whatsAppGenerateQR(qrChannel)
	return &LoginResponse{
		QrCode:  qrCode,
		Timeout: timeout,
	}, nil
}

func whatsAppGenerateQR(qrChan <-chan whatsmeow.QRChannelItem) (string, int) {
	qrChanCode := make(chan string)
	qrChanTimeout := make(chan int)

	// Get QR Code Data and Timeout
	go func() {
		for evt := range qrChan {
			if evt.Event == whatsmeow.QRChannelEventCode {
				qrChanCode <- evt.Code
				qrChanTimeout <- int(evt.Timeout.Seconds())
			}
		}
	}()

	qrCode := <-qrChanCode
	// qrPNG, _ := qrCode.Encode(qrTemp, qrCode.Medium, 256)

	return qrCode, <-qrChanTimeout
}
