package config

func GetLogLevel() string {
	return getStringVariable("LOG_LEVEL", "info")
}
