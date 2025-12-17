package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Logger = log.Output(os.Stdout)

	log.Info().
		Str("user", "quan").
		Msg("User logged in")

	log.Error().
		Str("file", "config.json").
		Msg("File not found")
}
