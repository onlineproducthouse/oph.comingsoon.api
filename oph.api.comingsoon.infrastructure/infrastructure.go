package infrastructure

import (
	"oph.api.comingsoon/oph.api.comingsoon.infrastructure/config"
	"oph.api.comingsoon/oph.api.comingsoon.infrastructure/loglocal"
)

type (
	IInfrastructure interface {
		Config() config.TConfig
		Logger() loglocal.Logger
	}

	Infrastructure struct {
		config config.TConfig
		logger loglocal.Logger
	}
)

func NewInfrastructure() Infrastructure {
	cfg := config.Config()

	return Infrastructure{
		config: cfg,
		logger: loglocal.New(cfg),
	}
}

func (infra Infrastructure) Config() config.TConfig {
	return infra.config
}

func (infra Infrastructure) Logger() loglocal.Logger {
	return infra.logger
}
