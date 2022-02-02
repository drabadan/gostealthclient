package config

const (
	LOG_LEVEL_DEBUG   = 0
	LOG_LEVEL_INFO    = 1
	LOG_LEVEL_WARNING = 2
	LOG_LEVEL_ERROR   = 3
)

type Config struct {
	LogLevel byte
}

func NewConfig(logLevel byte) *Config {
	return &Config{
		LogLevel: logLevel,
	}
}
