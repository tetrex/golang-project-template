package health

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type HealthService struct{}

func NewHealthService() *HealthService {

	return &HealthService{}
}

// @tags			Health
// @summary			Get Health check status
// @description		returns server time
// @accept			json
// @produce			json
// @success			200	{object}	int64
// @failure			500	{object}	utils.ErrorResponse
// @router			/ [get]
func (s *HealthService) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, time.Now().UnixNano())
}
