package adapter

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type logAdapter struct{}

func GetLogAdapter() *logAdapter {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	return &logAdapter{}
}

func (a *logAdapter) Error(err error) {
	log.Error().Msg(err.Error())
}

func (a *logAdapter) Info(msg string) {
	log.Info().Msg(msg)
}
