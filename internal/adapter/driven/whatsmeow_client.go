package driven

import (
	"app/internal/infra"
	"app/internal/port/driven"
	"context"
	"errors"

	_ "github.com/lib/pq"
	"go.mau.fi/whatsmeow"
	waproto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

type WhatsmewClient struct {
	db        *infra.PostgresDB
	container *sqlstore.Container
	dbLog     waLog.Logger
	session   map[string]*whatsmeow.Client
}

type qrCodeReceiverParams struct {
	qrChan <-chan whatsmeow.QRChannelItem
	userID string
	client *whatsmeow.Client
}

func NewWhatsmewClient(db *infra.PostgresDB) *WhatsmewClient {
	store.DeviceProps.Os = proto.String("go-api")
	store.DeviceProps.PlatformType = waproto.DeviceProps_CHROME.Enum()
	store.DeviceProps.RequireFullSync = proto.Bool(false)
	store.DeviceProps.HistorySyncConfig = &waproto.DeviceProps_HistorySyncConfig{
		FullSyncDaysLimit:   proto.Uint32(1),
		FullSyncSizeMbLimit: proto.Uint32(10),
		StorageQuotaMb:      proto.Uint32(10),
	}

	dbLog := waLog.Stdout("Database", "DEBUG", true)
	container, err := sqlstore.New("postgres", db.Dsn(), dbLog)
	if err != nil {
		panic("cannot initialize connection on whatsapp postgres: " + err.Error())
	}

	return &WhatsmewClient{
		db:        db,
		container: container,
		dbLog:     dbLog,
		session:   make(map[string]*whatsmeow.Client),
	}
}

func (ws *WhatsmewClient) NewQRCodeSession(userID string) (*driven.AuthQRResponse, error) {
	if _, ok := ws.session[userID]; ok {
		return nil, errors.New("failed generate session: please login in")
	}

	device := ws.container.NewDevice()
	client := whatsmeow.NewClient(device, nil)
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

func (ws *WhatsmewClient) whatsAppGenerateQR(params *qrCodeReceiverParams) (string, int64) {
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
					// need to save to db with mapping id <-> client.Store.ID
					err := ws.creatMappingUser(params.userID, params.client.Store.ID)
					if err != nil {
						break
					}
					ws.session[params.userID] = params.client
				}
			}
		}
	}()

	return <-qrChanCode, <-qrChanTimeout
}

func (ws *WhatsmewClient) creatMappingUser(userID string, jid *types.JID) error {
	_, err := ws.db.Conn().Exec(context.Background(), `
		INSERT INTO
			user_whatsmeow_map (user_id, jid)
		VALUES ($1, $2)
	`, userID, jid.String())
	return err
}
