package wlog

import "github.com/rs/zerolog/log"

func Error(msg string, err error) {
	log.Error().Err(err).Msg(msg)
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func Print(msg interface{}) {
	log.Print(msg)
}
