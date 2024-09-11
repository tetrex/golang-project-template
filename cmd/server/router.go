package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func initRoutes(router *echo.Echo, services *Services, l *zerolog.Logger) {
	router.GET("/", services.Health.HealthCheck)

	router.GET("docs/*", echoSwagger.WrapHandler)

	l.Info().Msgf("docs v1 at :: http://localhost:%d/docs/index.html \n", 8000)

}
