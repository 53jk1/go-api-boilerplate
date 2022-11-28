package config

type Env struct {
	// App
	AppName string `env:"APP_NAME" envDefault:"go-api-boilerplate"`
	AppEnv  string `env:"APP_ENV" envDefault:"development"`
	AppPort string `env:"APP_PORT" envDefault:"8080"`

	// Database
	DatabaseHost     string `env:"DATABASE_HOST" envDefault:"localhost"`
	DatabasePort     string `env:"DATABASE_PORT" envDefault:"5432"`
	DatabaseUser     string `env:"DATABASE_USER" envDefault:"postgres"`
	DatabasePassword string `env:"DATABASE_PASSWORD" envDefault:"postgres"`
	DatabaseName     string `env:"DATABASE_NAME" envDefault:"postgres"`
	DatabaseSSLMode  string `env:"DATABASE_SSL_MODE" envDefault:"disable"`
}
