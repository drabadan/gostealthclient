package config

const (
	LOG_LEVEL_DEBUG   = 0
	LOG_LEVEL_INFO    = 1
	LOG_LEVEL_WARNING = 2
	LOG_LEVEL_ERROR   = 3

	// In seconds
	SOCKET_TIMEOUT     = 2
	SOCKET_MAX_RETRIES = 10
)

type Config struct {
	LogLevel byte
}

func NewConfig(logLevel byte) *Config {
	return &Config{
		LogLevel: logLevel,
	}
}
