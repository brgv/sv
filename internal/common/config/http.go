package config

import "time"

type HttpConfiguration struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	GraceTimeout time.Duration
}
