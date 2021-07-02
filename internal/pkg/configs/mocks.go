package configs

import "os"

func loadMock() Config {
	port := os.Getenv("TEST_API_PORT")

	return Config{
		Port: port,
	}
}