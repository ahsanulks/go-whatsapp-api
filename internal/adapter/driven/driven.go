package driven

import (
	userrepository "app/internal/adapter/driven/user_repository"
	whatsmeowclient "app/internal/adapter/driven/whatsmeow"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(whatsmeowclient.NewWhatsmeowClient, userrepository.NewUserRepository)
