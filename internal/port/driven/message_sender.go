package driven

import "context"

type MessageParam struct {
	ReceiverNumbers []string
	Message         string
	Sender          string
}

type MessageSender interface {
	SendMessage(ctx context.Context, param *MessageParam) error
}
