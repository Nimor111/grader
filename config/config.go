package config

type dbConfig struct {
	DBUser     string
	DBPassword string
	DBName     string
}

type Config struct {
	Port     int
	Postgres dbConfig
}

func NewConfig(port int, dbUser string, dbPass string, dbName string) *Config {
	return &Config{
		Port: port,
		Postgres: dbConfig{
			DBUser:     dbUser,
			DBPassword: dbPass,
			DBName:     dbName,
		},
	}
}
