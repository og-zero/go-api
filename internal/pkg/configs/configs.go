package configs

type (
	Config struct {
		Port string
	}
)

func Load() Config {
	return loadMock()
}

