package driven

import "context"

type AuthQRResponse struct {
	QRCode  string
	Timeout int64
}

type SendMessageRequest struct {
	ReceiverPhone []string
	Message       string
}

type DeviceProvider interface {
	NewQRCodeSession(ctx context.Context, id string) (*AuthQRResponse, error)
}
