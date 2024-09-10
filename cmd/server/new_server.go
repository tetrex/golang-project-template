package main

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	db "github.com/tetrex/golang-project-template/db/sqlc"
	healthService "github.com/tetrex/golang-project-template/services/health"
	"github.com/tetrex/golang-project-template/utils/config"
	util_validator "github.com/tetrex/golang-project-template/utils/validator"
	"golang.org/x/time/rate"
)

type Services struct {
	Health *healthService.HealthService
}
type Server struct {
	config   config.Config
	router   *echo.Echo
	logger   *zerolog.Logger
	queries  *db.Queries
	services *Services
}

func (s *Server) GetServices() *Services {
	return s.services
}

func (s *Server) GetConfig() config.Config {
	return s.config
}

func (s *Server) GetRouter() *echo.Echo {
	return s.router
}

func (s *Server) GetLogger() *zerolog.Logger {
	return s.logger
}

func (s *Server) GetQueries() *db.Queries {
	return s.queries
}

type ServerParams struct {
	Config  config.Config
	Logger  *zerolog.Logger
	Queries *db.Queries
}

func NewServer(c ServerParams) (*Server, error) {
	router := echo.New()
	new_validator := validator.New()

	// NOTE: custom validator
	// if err := new_validator.RegisterValidation("env_validator", utils.EnvValidator); err != nil {
	// 	fmt.Println("error in validator registration")
	// 	log.Fatal(err)
	// }

	router.Validator = &util_validator.CustomValidator{Validator: new_validator}

	// logger
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, remote_ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	// stack trace for debugging
	router.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))

	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: rate.Limit(100), Burst: 30, ExpiresIn: 1 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}

	router.Use(middleware.RateLimiterWithConfig(config))

	// TODO: enble before deployment
	// cors
	// router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins:     []string{"https://*.domain.com", "https://v2.domain.com"},
	// 	AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodOptions},
	// 	AllowCredentials: true,
	// 	AllowHeaders:     []string{"application/json", "text/plain", "*/*"},
	// 	Skipper: func(c echo.Context) bool {
	// 		// if User-Agent is postman then skip
	// 		h := c.Request().Header.Get("User-Agent")
	// 		user_agent := strings.Split(h, "/")

	// 		return strings.EqualFold("PostmanRuntime", user_agent[0])

	// 	},
	// }))

	// setup routes here

	// some specific tasks
	switch c.Config.AppEnv {
	case "stage":
	case "prod":
	case "local":
	default:
	}
	health_service := healthService.NewHealthService()

	return &Server{
		config:  c.Config,
		router:  router,
		logger:  c.Logger,
		queries: c.Queries,
		services: &Services{
			Health: health_service,
		},
	}, nil
}
