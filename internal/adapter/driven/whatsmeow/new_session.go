package whatsmeowclient

import (
	"app/internal/port/driven"
	"context"
	"errors"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
)

type qrCodeReceiverParams struct {
	qrChan <-chan whatsmeow.QRChannelItem
	userID string
	client *whatsmeow.Client
}

func (ws *WhatsmeowClient) NewQRCodeSession(ctx context.Context, userID string) (*driven.AuthQRResponse, error) {
	device := ws.container.NewDevice()
	client := whatsmeow.NewClient(device, nil)
	// need use fresh context. because it can make we cannot login into whatsapp
	qrChannel, err := client.GetQRChannel(context.Background())
	if err != nil {
		return nil, err
	}

	err = client.Connect()
	if err != nil {
		return nil, errors.New("failed connect whatsapp: " + err.Error())
	}

	qrCode, timeout := ws.whatsAppGenerateQR(&qrCodeReceiverParams{
		qrChan: qrChannel,
		userID: userID,
		client: client,
	})
	return &driven.AuthQRResponse{
		QRCode:  qrCode,
		Timeout: timeout,
	}, nil
}

func (ws *WhatsmeowClient) whatsAppGenerateQR(params *qrCodeReceiverParams) (string, int64) {
	qrChanCode := make(chan string)
	qrChanTimeout := make(chan int64)

	// Get QR Code Data and Timeout
	// This go routine still running even this method already return
	go func() {
		for evt := range params.qrChan {
			if evt.Event == whatsmeow.QRChannelEventCode {
				qrChanCode <- evt.Code
				qrChanTimeout <- int64(evt.Timeout.Seconds())
			} else {
				if evt.Event == whatsmeow.QRChannelSuccess.Event {
					waID := params.client.Store.ID.User
					if _, exists := ws.session[waID]; exists {
						// if connection already exists, we don't need to create a new one
						// use existing
						err := params.client.Logout()
						if err != nil {
							params.client.Store.Delete()
						}
					} else {
						ws.session[waID] = params.client
						err := ws.creatMappingUser(params.userID, params.client.Store.ID)
						if err != nil {
							break
						}
					}
				}
			}
		}
	}()

	return <-qrChanCode, <-qrChanTimeout
}

func (ws *WhatsmeowClient) creatMappingUser(userID string, jid *types.JID) error {
	_, err := ws.db.Conn().Exec(context.Background(), `
		INSERT INTO
			user_whatsmeow_map (user_id, jid)
		VALUES ($1, $2)
	`, userID, jid.String())
	return err
}
