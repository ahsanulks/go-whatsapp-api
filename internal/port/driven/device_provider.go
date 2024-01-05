package driven

import "go.mau.fi/whatsmeow/store"

type DeviceProvider interface {
	NewDeviceSession() *store.Device
}
