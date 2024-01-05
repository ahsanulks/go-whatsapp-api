package authentication

type LoginQRResponse struct {
	QrCode    string
	ExpiredAt int64
}
