package whatsmeowclient

import (
	"app/internal/infra"

	_ "github.com/lib/pq"
	"go.mau.fi/whatsmeow"
	waproto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

type WhatsmeowClient struct {
	db        *infra.PostgresDB
	container *sqlstore.Container
	dbLog     waLog.Logger
	session   map[string]*whatsmeow.Client
}

// NOTE: PLEASE REFER BANNED REASON WHEN USING THIS
// "too many people blocked you",
// "you sent too many messages to people who don't have you in their address books"
// "you created too many groups with people who don't have you in their address books"
// "you sent the same message to too many people"
// "you sent too many messages to a broadcast list"
func NewWhatsmeowClient(db *infra.PostgresDB) *WhatsmeowClient {
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

	wc := &WhatsmeowClient{
		db:        db,
		container: container,
		dbLog:     dbLog,
		session:   make(map[string]*whatsmeow.Client),
	}
	wc.loadAllSession()
	return wc
}

// func (ws *WhatsmeowClient) reconnect(ctx context.Context, userID string) (*whatsmeow.Client, error) {
// 	if client, ok := ws.session[userID]; ok {
// 		return client, nil
// 	}

// 	jid, err := ws.getJidFromUserID(ctx, userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	device, err := ws.container.GetDevice(*jid)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// when not found device, whatsmeow return nil instead of error
// 	if device == nil {
// 		return nil, errors.New("cannot found session id for that number")
// 	}

// 	client := whatsmeow.NewClient(device, nil)
// 	err = client.Connect()
// 	if err != nil {
// 		return nil, err
// 	}
// 	ws.session[device.ID.User] = client

// 	// make server know the name of the user
// 	// client.Store.PushName = "ahsan test api"
// 	_ = client.SendPresence(types.PresenceAvailable)
// 	return client, nil
// }

// func (ws *WhatsmeowClient) getJidFromUserID(ctx context.Context, userID string) (*types.JID, error) {
// 	var stringJid string
// 	err := ws.db.Conn().QueryRow(ctx, `
// 		SELECT
// 			jid
// 		FROM
// 			user_whatsmeow_map
// 		WHERE
// 			user_id = $1
// 		LIMIT
// 			1
// 	`, userID).Scan(&stringJid)
// 	if err != nil {
// 		return nil, err
// 	}
// 	jid, err := types.ParseJID(stringJid)
// 	return &jid, err
// }
