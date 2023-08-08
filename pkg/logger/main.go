package logger

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func Init(LogLevel int) {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if LogLevel == 0 {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
