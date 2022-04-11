package routes

import (
	"fmt"
	"github.com/alganbr/kedai-itemsvc/configs"
	"github.com/alganbr/kedai-itemsvc/docs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"strings"
)

type SwaggerRoutes struct {
	router Router
	cfg    *configs.Config
}

func NewSwaggerRoutes(router Router, cfg *configs.Config) SwaggerRoutes {
	return SwaggerRoutes{
		router: router,
		cfg:    cfg,
	}
}

func (r *SwaggerRoutes) Setup() {
	docs.SwaggerInfo.BasePath = fmt.Sprintf("/%s", strings.ToLower(r.cfg.Server.Name))
	r.router.Path.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
