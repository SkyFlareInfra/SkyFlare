package repository

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewDomainRepository),
	fx.Provide(NewNameserverRepository),
)
