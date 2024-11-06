package config

type Conf struct {
	App struct {
		Name string `env:"APP_NAME"`
		Port string `env:"APP_PORT"`
	}

	Mail struct {
		Address  string `env:"MAIL_ADDRESS"`
		Password string `env:"MAIL_PASSWORD"`
		Host     string `env:"MAIL_HOST"`
		Port     string `env:"MAIL_PORT"`
	}

	Kafka struct {
		Endpoint string `env:"KAFKA_ENDPOINT"`
	}
}
