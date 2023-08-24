package configs

import (
	"github.com/spf13/viper"
	"log"
)

type envConfig struct {
	DBDsn         string `mapstructure:"DB_DSN"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	MigrationDir  string `mapstructure:"MIGRATION_DIR"`
}

func LoadEnvConfig(filePath string) *envConfig {
	var cfg *envConfig
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.SetConfigFile(filePath)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	log.Println("arquivo .env carregado")
	return cfg
}
