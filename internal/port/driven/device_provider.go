package driven

type AuthQRResponse struct {
	QRCode  string
	Timeout int64
}

type DeviceProvider interface {
	NewQRCodeSession(id string) (*AuthQRResponse, error)
}
