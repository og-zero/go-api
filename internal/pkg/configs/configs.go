package configs

import "os"

type (
	Config struct {
		Port string
	}
)

func Load() Config {
	return loadMock()
}

func loadMock() Config {
	port := os.Getenv("TEST_API_PORT")

	return Config{
		Port: port,
	}
}
