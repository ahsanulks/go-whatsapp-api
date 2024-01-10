package driver

import (
	v1 "app/api/v1"
	"app/internal/port/driver"
	"app/internal/usecase/message"
	"context"
)

type MessageHandler struct {
	v1.UnimplementedMessageServer
	usecase driver.MessageUsecase
}

func NewMessageHandler(msgUsecase driver.MessageUsecase) *MessageHandler {
	return &MessageHandler{
		usecase: msgUsecase,
	}
}

func (mh MessageHandler) SendMessage(ctx context.Context, params *v1.SendMessageRequest) (*v1.SendMessageResponse, error) {
	err := mh.usecase.SendMessage(ctx, &message.SendMessageRequest{
		ReceiverPhone: params.ReceiverPhones,
		Message:       params.Message,
		Sender: &message.Sender{
			Phone: params.Phone,
			ID:    params.Id,
		},
	})
	if err != nil {
		return nil, err
	}
	return &v1.SendMessageResponse{
		Message: "success send message",
	}, nil
}
