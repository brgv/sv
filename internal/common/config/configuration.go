package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func init() {
	viper.AddConfigPath(".")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("")

	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	viper.SetDefault("cwd", cwd)

	viper.SetDefault("migrate.dir", fmt.Sprintf("%s/data/migrations", cwd)) // "file:///migrations"

	viper.SetDefault("logger.level", "debug")

	viper.SetDefault("metric.port", "9090")
	viper.SetDefault("metric.readtimeout", "30s")
	viper.SetDefault("metric.writetimeout", "30s")

	viper.SetDefault("http.port", "8080")
	viper.SetDefault("http.readtimeout", "30s")
	viper.SetDefault("http.writetimeout", "30s")
	viper.SetDefault("http.gracetimeout", "30s")

	viper.SetDefault("grpc.port", "50051")
}

func NewConfiguration() (*Configuration, error) {
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var configuration Configuration

	if err := viper.Unmarshal(&configuration); err != nil {
		return nil, err
	}

	return &configuration, nil
}

type Configuration struct {
	Cwd      string
	Database DatabaseConfiguration
	Migrate  MigrateConfiguration
	Logger   LoggerConfiguration
	Metric   MetricConfiguration
	Http     HttpServerConfiguration
	Grpc     GrpcServerConfiguration
}
