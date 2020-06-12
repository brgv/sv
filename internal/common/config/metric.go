package config

import "time"

type MetricConfiguration struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
