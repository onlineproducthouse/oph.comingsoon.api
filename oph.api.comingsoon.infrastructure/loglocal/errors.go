package loglocal

import (
	"strings"

	"github.com/rs/zerolog/log"
	"oph.api.comingsoon/oph.api.comingsoon.util/errorlocal"
)

func (logger Logger) AppError(err errorlocal.IAppError) {
	if strings.ToLower(logger.config.EnvName()) == "local" ||
		strings.ToLower(logger.config.EnvName()) == "test" {
		log.Err(err)
	}

	log.Error().
		Str("kind", err.Kind()).
		Int("httpStatusCode", err.StatusCode()).
		// Str("innerMessage", err.Trace().InnerMessage).
		Str("trace", strings.Join(err.Trace().Ops, ", ")).
		Msg(err.Error())
}
