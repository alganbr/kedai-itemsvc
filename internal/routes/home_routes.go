package routes

import "github.com/alganbr/kedai-itemsvc/internal/controllers"

type HomeRoutes struct {
	router         Router
	homeController controllers.IHomeController
}

func NewHomeRoutes(router Router, homeController controllers.IHomeController) HomeRoutes {
	return HomeRoutes{
		router:         router,
		homeController: homeController,
	}
}

func (r *HomeRoutes) Setup() {
	r.router.Path.GET("/", r.homeController.HealthCheck)

	homeGroup := r.router.Path.Group("/home")
	homeGroup.GET("/health-check", r.homeController.HealthCheck)
}
