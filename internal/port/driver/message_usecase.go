package driver

import (
	"app/internal/usecase/message"
	"context"
)

type MessageUsecase interface {
	SendMessage(ctx context.Context, request *message.SendMessageRequest) error
}
