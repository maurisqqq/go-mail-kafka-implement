package config

type Conf struct {
	App struct {
		Name string `env:"APP_NAME"`
		Port string `env:"APP_PORT"`
	}

	Kafka struct {
		Endpoint string `env:"KAFKA_ENDPOINT"`
	}
}
