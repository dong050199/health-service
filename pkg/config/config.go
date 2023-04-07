package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

var (
	dbConfig  DBConfig
	jwtConfig JWTConfig
)

type DBConfig struct {
	DBName     string `envconfig:"DBNAME"`
	DBURL      string `envconfig:"DBURL"`
	DBPort     string `envconfig:"DBPORT"`
	DBUserName string `envconfig:"DBUSERNAME"`
	DBPassword string `envconfig:"DBPASSWORD"`
}

type JWTConfig struct {
	APISecret         string `envconfig:"API_SECRET"`
	TokenHourLifeSpan string `envconfig:"TOKEN_HOUR_LIFESPAN"`
}

func SetConfig() {
	configs := []interface{}{
		&dbConfig,
		&jwtConfig,
	}
	for _, instance := range configs {
		err := envconfig.Process("", instance)
		if err != nil {
			log.Fatalf("unable to init config: %v, err: %v", instance, err)
		}
	}
}

func DbConfig() DBConfig {
	return dbConfig
}

func JwtConfig() JWTConfig {
	return jwtConfig
}
