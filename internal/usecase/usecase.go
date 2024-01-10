package usecase

import (
	"app/internal/usecase/authentication"
	"app/internal/usecase/message"

	"github.com/google/wire"
)

// ProviderSet is usecase providers.
var ProviderSet = wire.NewSet(
	authentication.NewLoginUsecase,
	message.NewMessageUsecase,
)
