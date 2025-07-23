package store

type Config struct{
	DatabaseURL string
}

func NewConfig() *Config{
	return &Config{
		DatabaseURL: "web:3758@/money_test?parseTime=true",
	}
}
