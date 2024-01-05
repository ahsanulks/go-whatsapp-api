package driven

import (
	"app/internal/infra"

	_ "github.com/lib/pq"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

type WhatsappStore struct {
	db        *infra.PostgresDB
	container *sqlstore.Container
	dbLog     waLog.Logger
}

func NewWhatsappStore(db *infra.PostgresDB) *WhatsappStore {
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	container, err := sqlstore.New("postgres", db.Dsn(), dbLog)
	// store.DeviceProps.Os = proto.String("windows")
	// store.DeviceProps.PlatformType = waproto.DeviceProps_CHROME.Enum()
	// store.DeviceProps.RequireFullSync = proto.Bool(false)
	if err != nil {
		panic("cannot create connection on postgres: " + err.Error())
	}
	return &WhatsappStore{
		db:        db,
		container: container,
		dbLog:     dbLog,
	}
}

func (ws *WhatsappStore) NewDeviceSession() *store.Device {
	device := ws.container.NewDevice()
	return device
}
