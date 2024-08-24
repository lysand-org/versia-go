package config

import (
	"git.devminer.xyz/devminer/unitel"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"net/url"
	"os"
	"regexp"
	"strconv"
)

type Config struct {
	Port    int
	TLSKey  *string
	TLSCert *string

	PublicAddress  *url.URL
	Host           string
	SharedInboxURL *url.URL

	InstanceName        string
	InstanceDescription *string

	NATSURI        string
	NATSStreamName string

	DatabaseURI string

	Telemetry       unitel.Opts
	ForwardTracesTo *regexp.Regexp
}

var C Config

func Load() {
	if err := godotenv.Load(".env.local", ".env"); err != nil {
		log.Warn().Err(err).Msg("Failed to load .env file")
	}

	publicAddress, err := url.Parse(os.Getenv("VERSIA_INSTANCE_ADDRESS"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse VERSIA_INSTANCE_ADDRESS")
	}

	var forwardTracesTo *regexp.Regexp
	{
		rawForwardTracesTo := optionalEnvStr("FORWARD_TRACES_TO")
		if rawForwardTracesTo == nil {
			s := "matchnothing^"
			rawForwardTracesTo = &s
		}
		if forwardTracesTo, err = regexp.Compile(*rawForwardTracesTo); err != nil {
			log.Fatal().Err(err).Str("raw", *rawForwardTracesTo).Msg("Failed to compile")
		}
	}

	tlsKey := optionalEnvStr("VERSIA_TLS_KEY")
	tlsCert := optionalEnvStr("VERSIA_TLS_CERT")
	if (tlsKey != nil && tlsCert == nil) || (tlsKey == nil && tlsCert != nil) {
		log.Fatal().
			Msg("Both VERSIA_TLS_KEY and VERSIA_TLS_CERT have to be set if you want to use in-process TLS termination.")
	}

	C = Config{
		Port:    getEnvInt("VERSIA_PORT", 80),
		TLSCert: tlsCert,
		TLSKey:  tlsKey,

		PublicAddress:  publicAddress,
		Host:           publicAddress.Host,
		SharedInboxURL: publicAddress.ResolveReference(&url.URL{Path: "/api/inbox"}),

		InstanceName:        os.Getenv("VERSIA_INSTANCE_NAME"),
		InstanceDescription: optionalEnvStr("VERSIA_INSTANCE_DESCRIPTION"),

		NATSURI:        os.Getenv("NATS_URI"),
		NATSStreamName: getEnvStr("NATS_STREAM_NAME", "versia-go"),
		DatabaseURI:    os.Getenv("DATABASE_URI"),

		ForwardTracesTo: forwardTracesTo,
		Telemetry:       unitel.ParseOpts("versia-go"),
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
