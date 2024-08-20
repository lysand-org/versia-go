package config

import (
	"net/url"
	"os"
	"strconv"

	"git.devminer.xyz/devminer/unitel"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Port int

	PublicAddress  *url.URL
	Host           string
	SharedInboxURL *url.URL

	InstanceName        string
	InstanceDescription *string

	NATSURI        string
	NATSStreamName string

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
		Port: getEnvInt("PORT", 80),

		PublicAddress:  publicAddress,
		Host:           publicAddress.Host,
		SharedInboxURL: publicAddress.ResolveReference(&url.URL{Path: "/api/inbox"}),

		InstanceName:        os.Getenv("INSTANCE_NAME"),
		InstanceDescription: optionalEnvStr("INSTANCE_DESCRIPTION"),

		NATSURI:        os.Getenv("NATS_URI"),
		NATSStreamName: getEnvStr("NATS_STREAM_NAME", "versia-go"),
		DatabaseURI:    os.Getenv("DATABASE_URI"),

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

func getEnvStr(key, default_ string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return default_
}

func getEnvInt(key string, default_ int) int {
	if value, ok := os.LookupEnv(key); ok {
		parsed, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}

		return parsed
	}

	return default_
}
