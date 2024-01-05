package driven

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewWhatsappStore)
