package driven

import (
	whatsmeowclient "app/internal/adapter/driven/whatsmeow"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(whatsmeowclient.NewWhatsmeowClient)
