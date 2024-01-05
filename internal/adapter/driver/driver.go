package driver

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewLoginHandler)
