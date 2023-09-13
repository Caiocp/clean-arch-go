package configs

import "github.com/spf13/viper"

type config struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBName            string `mapstructure:"DB_NAME"`
	WebServerPort     string `mapstructure:"WEB_SERVER_PORT"`
	GRPCServerPort    string `mapstructure:"GRPC_SERVER_PORT"`
	GraphQLServerPort string `mapstructure:"GRAPHQL_SERVER_PORT"`
	RabbitMQHost      string `mapstructure:"RABBITMQ_URL"`
	RabbitMQUser      string `mapstructure:"RABBITMQ_USER"`
	RabbitMQPassword  string `mapstructure:"RABBITMQ_PASSWORD"`
	RabbitMQPort      string `mapstructure:"RABBITMQ_PORT"`
}

func LoadConfig(path string) (*config, error) {
	var cfg *config

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg, nil
}
