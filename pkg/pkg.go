package pkg

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(LoadConfig),
	fx.Provide(CreateLogService),
)
