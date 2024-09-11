package main

import healthService "github.com/tetrex/golang-project-template/services/health"

func initServices() *Services {
	health_service := healthService.NewHealthService()

	return &Services{
		Health: health_service,
	}
}
