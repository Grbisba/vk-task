package main

import (
	"go.uber.org/fx"

	"github.com/vk-task/client/internal/config"
)

func main() {
	fx.New(buildOptions()).Run()
}

func buildOptions() fx.Option {
	return fx.Options(
		fx.Provide(
			config.New,
		),
		fx.Invoke(),
	)
}
