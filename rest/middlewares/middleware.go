package middleware

import "kholabazar/config"

type Middlewares struct {
	conf *config.Config
}

func NewMiddlewares(conf *config.Config) *Middlewares {
	return &Middlewares{
		conf: conf,
	}
}
