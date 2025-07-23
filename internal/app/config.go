package app

import(
	"money/internal/store"
)

type Config struct {
	BindAddr string
	LogLevel string
	Store *store.Config
}

func NewConfig() *Config{
	return &Config{
		BindAddr: ":5050",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}