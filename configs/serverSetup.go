package configs

import "time"

type Server struct {
	Mode         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func ServerSetup() bool {
	return true
}
