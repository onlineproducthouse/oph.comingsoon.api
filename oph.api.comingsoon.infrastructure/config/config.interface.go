package config

type IConfig interface {
	ProjectName() string
	ProjectShortName() string
	Host() string
	Port() string
	EnvName() string
	ReqHeaderRequestID() string
	RunSwagger() bool
	RequestIDKey() string
}
