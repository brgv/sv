package config

import "time"

type HttpServerConfiguration struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	GraceTimeout time.Duration
}
