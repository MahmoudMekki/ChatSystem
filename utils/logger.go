package utils

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	host, _ := os.Hostname()
	log.Logger = log.With().Str("host", host).Logger()
	log.Logger = log.With().Str("service", "chat").Logger()
	log.Logger = log.With().Caller().Logger()
}
