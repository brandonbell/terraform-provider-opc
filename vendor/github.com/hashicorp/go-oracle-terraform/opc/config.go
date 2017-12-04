package opc

import (
	"net/http"
	"net/url"
)

type Config struct {
	Username       *string
	Password       *string
        ContainerPath  *string
	IdentityDomain *string
	APIEndpoint    *url.URL
	MaxRetries     *int
	LogLevel       LogLevelType
	Logger         Logger
	HTTPClient     *http.Client
	UserAgent      *string
}

func NewConfig() *Config {
	return &Config{}
}
