package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IHomeController interface {
	HealthCheck(*gin.Context)
}

type HomeController struct {
}

func NewHomeController() IHomeController {
	return &HomeController{}
}

// HealthCheck godoc
// @Description  Get health check status
// @Tags         Home
// @Success      200  {object}  string
// @Router       /home/health-check [get]
func (ctrl *HomeController) HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "Ok")
}
