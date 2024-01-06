package whatsmeowclient

import (
	"context"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
	"golang.org/x/sync/errgroup"
)

// LoadAllSession will be run on startup
func (wc *WhatsmeowClient) loadAllSession() {
	devices, err := wc.container.GetAllDevices()
	if err != nil {
		panic("failed to get all devices")
	}
	errGroup, _ := errgroup.WithContext(context.Background())
	for _, device := range devices {
		if _, ok := wc.session[device.ID.User]; ok {
			continue
		}
		d := device
		errGroup.Go(func() error {
			client := whatsmeow.NewClient(d, nil)
			wc.session[d.ID.User] = client
			err := client.Connect()
			if err != nil {
				return err
			}
			return client.SendPresence(types.PresenceAvailable)
		})
	}
	err = errGroup.Wait()
	if err != nil {
		panic("error when connection session: " + err.Error())
	}
}
