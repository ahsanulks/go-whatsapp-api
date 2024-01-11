package whatsmeowclient

import (
	"app/internal/port/driven"
	"context"
	"time"

	"go.mau.fi/whatsmeow"
	waproto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

// TODO: send message need to call by background service
// because it's possible to send into multiple receiver
// we need to add delay on each message to prevent blocking mechanism from whatsapp
func (ws *WhatsmeowClient) SendMessage(ctx context.Context, params *driven.MessageParam) error {
	session := ws.session[params.Sender]
	responses, err := session.IsOnWhatsApp(params.ReceiverNumbers)
	if err != nil {
		return err
	}
	for _, resp := range responses {
		if resp.IsIn {
			// TODO: add delay on each message to prevent blocking
			// use new context since it will breaking using existing context
			// please refer to banned reason
			time.Sleep(2 * time.Second)
			_ = session.SendChatPresence(resp.JID, types.ChatPresenceComposing, types.ChatPresenceMediaText)
			_, err = session.SendMessage(context.Background(), resp.JID, &waproto.Message{
				Conversation: proto.String(params.Message),
			}, whatsmeow.SendRequestExtra{
				ID: session.GenerateMessageID(),
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}
