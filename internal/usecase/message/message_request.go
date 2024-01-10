package message

type SendMessageRequest struct {
	ReceiverPhone []string `validate:"required,gt=0"`
	Message       string   `validate:"required"`
	Sender        *Sender  `validate:"required"`
}

type Sender struct {
	Phone string `validate:"required"`
	ID    string `validate:"required"`
}
