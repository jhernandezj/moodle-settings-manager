package config

// Config holds the application configuration
type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

// New creates a new Config instance with default values
func New() *Config {
	return &Config{
		DBHost: "localhost",
		DBPort: 3306,
		DBUser: "root",
		DBName: "moodle",
	}
}
