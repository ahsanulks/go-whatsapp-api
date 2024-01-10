package message

import (
	"context"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestMessageSender_SendMessage(t *testing.T) {
	type fields struct {
		validator *validator.Validate
	}
	type args struct {
		ctx     context.Context
		request *SendMessageRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "when receiver empty should return error",
			fields: fields{
				validator: validator.New(),
			},
			args: args{
				ctx: context.Background(),
				request: &SendMessageRequest{
					ReceiverPhone: []string{},
					Message:       "test message",
					Sender:        &Sender{},
				},
			},
			wantErr: true,
		},
		{
			name: "when message empty should return error",
			fields: fields{
				validator: validator.New(),
			},
			args: args{
				ctx: context.Background(),
				request: &SendMessageRequest{
					ReceiverPhone: []string{"1231231451"},
					Message:       "",
					Sender:        &Sender{},
				},
			},
			wantErr: true,
		},
		{
			name: "when sender empty should return error",
			fields: fields{
				validator: validator.New(),
			},
			args: args{
				ctx: context.Background(),
				request: &SendMessageRequest{
					ReceiverPhone: []string{"1231231451"},
					Message:       "testing message",
					Sender:        &Sender{},
				},
			},
			wantErr: true,
		},
		{
			name: "when id sender empty should return error",
			fields: fields{
				validator: validator.New(),
			},
			args: args{
				ctx: context.Background(),
				request: &SendMessageRequest{
					ReceiverPhone: []string{"1231231451"},
					Message:       "testing message",
					Sender: &Sender{
						Phone: "1231313131",
						ID:    "",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "when phone sender empty should return error",
			fields: fields{
				validator: validator.New(),
			},
			args: args{
				ctx: context.Background(),
				request: &SendMessageRequest{
					ReceiverPhone: []string{"1231231451"},
					Message:       "testing message",
					Sender: &Sender{
						Phone: "",
						ID:    "testid",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "when id and phone not found should return error",
			fields: fields{
				validator: validator.New(),
			},
			args: args{
				ctx: context.Background(),
				request: &SendMessageRequest{
					ReceiverPhone: []string{"1231231451"},
					Message:       "testing message",
					Sender: &Sender{
						Phone: "123131",
						ID:    "testid",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &MessageSender{
				validator: tt.fields.validator,
			}
			if err := ms.SendMessage(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("MessageSender.SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
