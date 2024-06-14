package config

func (c TConfig) ProjectName() string {
	return get("PROJECT_NAME")
}

func (c TConfig) ProjectShortName() string {
	return get("PROJECT_SHORT_NAME")
}

func (c TConfig) Host() string {
	return get("API_HOST")
}

func (c TConfig) Port() string {
	return get("API_PORT")
}

func (c TConfig) EnvName() string {
	return get("ENVIRONMENT_NAME")
}

func (c TConfig) ReqHeaderRequestID() string {
	return "x-request-id"
}

func (c TConfig) RunSwagger() bool {
	return get("RUN_SWAGGER") == "true"
}

func (c TConfig) RequestIDKey() string {
	return "x-request-id"
}
