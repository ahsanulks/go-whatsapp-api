package usecase

import (
	"app/internal/usecase/authentication"

	"github.com/google/wire"
)

// ProviderSet is usecase providers.
var ProviderSet = wire.NewSet(
	NewGreeterUsecase,
	authentication.NewLoginUsecase,
)
