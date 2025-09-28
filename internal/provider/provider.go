//go:generate wire .

package provider

import (
	"github.com/NoANameGroup/DAOld-Backend/internal/config"
	"github.com/NoANameGroup/DAOld-Backend/internal/repository"
	"github.com/NoANameGroup/DAOld-Backend/internal/service"
	"github.com/google/wire"
)

var provider *Provider

func Init() {
	var err error
	provider, err = NewProvider()
	if err != nil {
		panic(err)
	}
}

func Get() *Provider {
	return provider
}

// Provider 提供controller依赖的对象
type Provider struct {
	Config      *config.Config
	UserService service.UserService
}

var ApplicationSet = wire.NewSet(
	service.UserServiceSet,
)

var InfrastructureSet = wire.NewSet(
	config.NewConfig,
	repository.NewUserRepository,
)

var AllProvider = wire.NewSet(
	ApplicationSet,
	InfrastructureSet,
)
