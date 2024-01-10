package message

import (
	"app/internal/port/driven"
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
)

type MessageUsecase struct {
	validator     *validator.Validate
	messageSender driven.MessageSender
	phoneChecker  driven.PhoneChecker
}

func NewMessageUsecase(
	validator *validator.Validate,
	messageSender driven.MessageSender,
	phoneCheker driven.PhoneChecker,
) *MessageUsecase {
	return &MessageUsecase{
		validator:     validator,
		messageSender: messageSender,
		phoneChecker:  phoneCheker,
	}
}

func (ms *MessageUsecase) SendMessage(ctx context.Context, request *SendMessageRequest) error {
	err := ms.validator.Struct(request)
	if err != nil {
		return err
	}

	isValid := ms.phoneChecker.IsPhoneValid(ctx, request.Sender.ID, request.Sender.Phone)
	if !isValid {
		return errors.New("invalid phone sender/id")
	}
	return ms.messageSender.SendMessage(ctx, &driven.MessageParam{
		ReceiverNumbers: request.ReceiverPhone,
		Message:         request.Message,
		Sender:          request.Sender.Phone,
	})
}
