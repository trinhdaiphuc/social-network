package config

type Config struct {
	Port     string
	JwtKey   string
	Env      string
	LogLevel string
	LogPath  string
	MongoURI string
}

var (
	configValue Config
)

func Load() Config {
	configValue = Config{
		Port:     env("PORT", "8080"),
		JwtKey:   env("JWT_KEY", "secret"),
		Env:      env("ENV", "local"),
		LogLevel: env("LOG_LEVEL", "INFO"),
		LogPath:  env("LOG_PATH", ""),
		MongoURI: env("DB_URI", "mongodb://localhost/social-network"),
	}
	return configValue
}

func GetConfig() Config {
	return configValue
}
