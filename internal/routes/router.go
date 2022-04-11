package routes

import (
	"fmt"
	"github.com/alganbr/kedai-itemsvc/configs"
	_ "github.com/alganbr/kedai-itemsvc/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
)

type Router struct {
	Gin  *gin.Engine
	Path *gin.RouterGroup
}

func NewRouter(cfg *configs.Config) Router {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	group := router.Group(fmt.Sprintf("/%s", strings.ToLower(cfg.Server.Name)))

	return Router{
		Gin:  router,
		Path: group,
	}
}

func (r Router) Run(cfg *configs.Config) {
	go func() {
		err := r.Gin.Run(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
		if err != nil {
			panic(err)
		}
	}()
}
