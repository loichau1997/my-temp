package conf

import "github.com/caarlos0/env/v6"

// AppConfig presents app conf
type AppConfig struct {
	Port string `env:"PORT" envDefault:"8002"`
	//DB CONFIG
	LogFormat string `env:"LOG_FORMAT" envDefault:"127.0.0.1"`

	EnableDB                 string `env:"ENABLE_DB" envDefault:"false"`
	CronJobThread            string `env:"CRONJOB_THREAD" envDefault:"4"`
	NumberOfProductPerThread string `env:"NUMBER_OF_PRODUCT_BY_THREAD" envDefault:"100"`
	MongoHost                string `env:"MONGO_HOST" envDefault:"127.0.0.1"`
	MongoPort                string `env:"MONGO_PORT" envDefault:"27017"`
	MongoUsername            string `env:"MONGO_USERNAME" envDefault:"evendo_mongodb_username"`
	MongoPassword            string `env:"MONGO_PASSWORD" envDefault:"evendo_mongodb_password"`
	// ENV
	EnvName string `env:"ENV_NAME" envDefault:"dev"`

	//API Integration
	ViatorEndpoint string `env:"VIATOR_ENDPOINT" envDefault:"https://api.sandbox.viator.com/partner/"`

	//Transfer
	TransferEndpoint string `env:"TRANSFER_ENDPOINT" envDefault:"18.133.219.163/api/v1/viator"`
}

var config AppConfig

func LoadConfig() {
	_ = env.Parse(&config)
}

func GetConfig() AppConfig {
	return config
}
