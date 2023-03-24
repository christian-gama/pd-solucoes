package env

// db is the database environment variables.
type db struct {
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	Name     string `env:"DB_NAME,required"`
	SslMode  string `env:"DB_SSL_MODE,required"`
}

// app is the application environment variables.
type app struct {
	Env  string `env:"APP_ENV,required"`
	Port int    `env:"APP_PORT,required"`
	Host string `env:"APP_HOST,required"`
}

// config is the application configuration.
type config struct {
	GlobalRateLimit uint `env:"CONFIG_GLOBAL_RATE_LIMIT,required"`
	Debug           bool `env:"CONFIG_DEBUG,required"`
}

var (
	// DB is the database environment variables.c.
	DB = &db{}

	// App is the application environment variables.
	App = &app{}

	// Config is the application configuration.
	Config = &config{}
)

const (
	// Development is the development mode environment.
	Development = "dev"

	// Production is the production mode environment.
	Production = "prod"

	// Test is the test mode environment.
	Test = "test"
)
