package errors

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewErrorBuilder),
)
