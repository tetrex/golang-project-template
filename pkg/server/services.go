package server

import healthService "github.com/tetrex/golang-project-template/pkg/server/services/health"

func initServices() *Services {
	health_service := healthService.NewHealthService()

	return &Services{
		Health: health_service,
	}
}
