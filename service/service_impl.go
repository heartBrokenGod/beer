package service

import "github.com/rs/zerolog"

type ServiceImpl struct {
	logger *zerolog.Logger
	config *Config
}

func NewServiceImpl(config *Config, logger *zerolog.Logger) *ServiceImpl {
	return &ServiceImpl{
		logger: logger,
		config: config,
	}
}
