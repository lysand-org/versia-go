package config

import (
	"net/url"
	"os"

	"git.devminer.xyz/devminer/unitel"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	PublicAddress       *url.URL
	InstanceName        string
	InstanceDescription *string

	NATSURI     string
	DatabaseURI string

	Telemetry unitel.Opts
}

var C Config

func Load() {
	if err := godotenv.Load(".env.local", ".env"); err != nil {
		log.Warn().Err(err).Msg("Failed to load .env file")
	}

	publicAddress, err := url.Parse(os.Getenv("PUBLIC_ADDRESS"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse PUBLIC_ADDRESS")
	}

	C = Config{
		PublicAddress:       publicAddress,
		InstanceName:        os.Getenv("INSTANCE_NAME"),
		InstanceDescription: optionalEnvStr("INSTANCE_DESCRIPTION"),

		NATSURI:     os.Getenv("NATS_URI"),
		DatabaseURI: os.Getenv("DATABASE_URI"),

		Telemetry: unitel.ParseOpts("versia-go"),
	}

	return
}

func optionalEnvStr(key string) *string {
	value := os.Getenv(key)
	if value == "" {
		return nil
	}
	return &value
}
