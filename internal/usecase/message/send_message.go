package message

import (
	"context"

	"github.com/go-playground/validator/v10"
)

type MessageSender struct {
	validator *validator.Validate
}

func NewMessageSender(validator *validator.Validate) *MessageSender {
	return &MessageSender{
		validator: validator,
	}
}

func (ms *MessageSender) SendMessage(ctx context.Context, request *SendMessageRequest) error {
	err := ms.validator.Struct(request)
	if err != nil {
		return err
	}
	return nil
}
